import React, { useState, useRef } from 'react';
import { Modal, message } from 'antd';
import { VehicleList } from '../components/vehicles/VehicleList';
import { VehicleSearch } from '../components/vehicles/VehicleSearch';
import { VehicleForm } from '../components/vehicles/VehicleForm';
import { VehicleCard } from '../components/vehicles/VehicleCard';
import { useVehicles } from '../hooks/useVehicles';
import { useVehiclesWebSocket } from '../hooks/useVehiclesWebSocket';
import { useVehicleStore } from '../store/vehicleStore';
import type { Vehicle, VehicleStatus } from '../types/vehicle';
type ViewMode = 'list' | 'create' | 'edit' | 'view';

export const VehiclesPage: React.FC = () => {
  const [viewMode, setViewMode] = useState<ViewMode>('list');
  const [selectedVehicle, setSelectedVehicle] = useState<Vehicle | null>(null);
  const [searchQuery, setSearchQuery] = useState('');
  const [searchStatus, setSearchStatus] = useState<VehicleStatus | undefined>();

  const { isLoading, refetch } = useVehicles(1, 10);
  const { addVehicle, updateVehicle: updateVehicleInStore, removeVehicle } = useVehicleStore();

  // Debounce для предотвращения дубликатов сообщений
  const lastMessageRef = useRef<{ type: string; id: string; timestamp: number } | null>(null);

  const shouldShowMessage = (type: string, id: string): boolean => {
    const now = Date.now();
    const lastMessage = lastMessageRef.current;

    if (!lastMessage) {
      lastMessageRef.current = { type, id, timestamp: now };
      return true;
    }

    // Если то же сообщение было показано менее 1 секунды назад, пропускаем
    if (lastMessage.type === type && lastMessage.id === id && now - lastMessage.timestamp < 1000) {
      return false;
    }

    lastMessageRef.current = { type, id, timestamp: now };
    return true;
  };

  // Обработка WebSocket событий
  const handleVehicleCreated = (data: Vehicle) => {
    if (shouldShowMessage('created', data.id)) {
      message.success(`Новая автовышка добавлена: ${data.garageNumber}`);
    }
    addVehicle(data);
    refetch();
  };

  const handleVehicleUpdated = (data: Vehicle) => {
    if (shouldShowMessage('updated', data.id)) {
      message.info(`Автовышка обновлена: ${data.garageNumber}`);
    }
    updateVehicleInStore(data.id, data);
    refetch();
  };

  const handleVehicleDeleted = (data: { id: string }) => {
    if (shouldShowMessage('deleted', data.id)) {
      message.warning('Автовышка удалена');
    }
    removeVehicle(data.id);
    refetch();
  };

  // Обработка WebSocket событий через типизированные хуки
  useVehiclesWebSocket({
    onCreated: handleVehicleCreated,
    onUpdated: handleVehicleUpdated,
    onDeleted: handleVehicleDeleted,
  });
  const handleCreate = () => {
    setSelectedVehicle(null);
    setViewMode('create');
  };

  const handleEdit = (vehicle: Vehicle) => {
    setSelectedVehicle(vehicle);
    setViewMode('edit');
  };

  const handleView = (vehicle: Vehicle) => {
    setSelectedVehicle(vehicle);
    setViewMode('view');
  };

  const handleDelete = async (id: string) => {
    try {
      removeVehicle(id);
      message.success('Автовышка успешно удалена');
      setViewMode('list');
      refetch();
    } catch (error) {
      message.error('Ошибка при удалении автовышки');
    }
  };

  const handleFormSuccess = () => {
    setViewMode('list');
    setSelectedVehicle(null);
  };

  const handleClearSearch = () => {
    setSearchQuery('');
    setSearchStatus(undefined);
  };
  const handleModalClose = () => {
    setViewMode('list');
    setSelectedVehicle(null);
  };

  return (
    <div>
      <VehicleSearch
        searchQuery={searchQuery}
        searchStatus={searchStatus}
        onSearchChange={setSearchQuery}
        onStatusChange={setSearchStatus}
        onClear={handleClearSearch}
        onAdd={handleCreate}
      />

      <VehicleList
        onEdit={handleEdit}
        onView={handleView}
        searchQuery={searchQuery}
        searchStatus={searchStatus}
      />

      {/* Модальное окно создания/редактирования */}
      <Modal
        title={viewMode === 'create' ? 'Создание автовышки' : 'Редактирование автовышки'}
        open={viewMode === 'create' || viewMode === 'edit'}
        onCancel={handleModalClose}
        footer={null}
        width={1000}
      >
        <VehicleForm
          vehicle={selectedVehicle || undefined}
          onSuccess={handleFormSuccess}
          onCancel={handleModalClose}
        />
      </Modal>

      {/* Модальное окно просмотра */}
      <Modal
        title="Детали автовышки"
        open={viewMode === 'view'}
        onCancel={handleModalClose}
        footer={null}
        width={1200}
      >
        {selectedVehicle && (
          <VehicleCard
            vehicle={selectedVehicle}
            onEdit={handleEdit}
            onDelete={handleDelete}
            loading={isLoading}
          />
        )}
      </Modal>
    </div>
  );
};
