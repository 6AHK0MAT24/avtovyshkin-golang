import React, { useState, useRef } from 'react';
import { Modal, message } from 'antd';
import { DriverList } from '../components/drivers/DriverList';
import { DriverSearch } from '../components/drivers/DriverSearch';
import { DriverForm } from '../components/drivers/DriverForm';
import { DriverCard } from '../components/drivers/DriverCard';
import { useDrivers } from '../hooks/useDrivers';
import { useWebSocket } from '../hooks/useWebSocket';
import { useDriverStore } from '../store/driverStore';
import type { Driver, DriverStatus } from '../types/driver';
type ViewMode = 'list' | 'create' | 'edit' | 'view';

export const DriversPage: React.FC = () => {
  const [viewMode, setViewMode] = useState<ViewMode>('list');
  const [selectedDriver, setSelectedDriver] = useState<Driver | null>(null);
  const [searchQuery, setSearchQuery] = useState('');
  const [searchStatus, setSearchStatus] = useState<DriverStatus | undefined>();

  const { isLoading, refetch } = useDrivers(1, 10);
  const { addDriver, updateDriver: updateDriverInStore, removeDriver } = useDriverStore();

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
  const handleDriverCreated = (data: Driver) => {
    if (shouldShowMessage('created', data.id)) {
      message.success(`Новый водитель добавлен: ${data.lastName} ${data.firstName}`);
    }
    addDriver(data);
    refetch();
  };

  const handleDriverUpdated = (data: Driver) => {
    if (shouldShowMessage('updated', data.id)) {
      message.info(`Водитель обновлен: ${data.lastName} ${data.firstName}`);
    }
    updateDriverInStore(data.id, data);
    refetch();
  };

  const handleDriverDeleted = (data: { id: string }) => {
    if (shouldShowMessage('deleted', data.id)) {
      message.warning('Водитель удален');
    }
    removeDriver(data.id);
    refetch();
  };

  useWebSocket('driver.created', handleDriverCreated);
  useWebSocket('driver.updated', handleDriverUpdated);
  useWebSocket('driver.deleted', handleDriverDeleted);
  useWebSocket('driver.updated', (data: Driver) => {
    console.log('driver.updated event received:', data);
    message.info(`Водитель обновлен: ${data.lastName} ${data.firstName}`);
    updateDriverInStore(data.id, data);
    refetch();
  });

  useWebSocket('driver.deleted', (data: { id: string }) => {
    console.log('driver.deleted event received:', data);
    message.warning('Водитель удален');
    removeDriver(data.id);
    refetch();
  });  const handleCreate = () => {
    setSelectedDriver(null);
    setViewMode('create');
  };

  const handleEdit = (driver: Driver) => {
    setSelectedDriver(driver);
    setViewMode('edit');
  };

  const handleView = (driver: Driver) => {
    setSelectedDriver(driver);
    setViewMode('view');
  };

  const handleDelete = async (id: string) => {
    try {
      removeDriver(id);
      message.success('Водитель успешно удален');
      setViewMode('list');
      refetch();
    } catch (error) {
      message.error('Ошибка при удалении водителя');
    }
  };
  const handleFormSuccess = () => {
    setViewMode('list');
    setSelectedDriver(null);
  };
  const handleSearch = (query: string, status?: DriverStatus) => {
    setSearchQuery(query);
    setSearchStatus(status);
  };

  const handleClearSearch = () => {
    setSearchQuery('');
    setSearchStatus(undefined);
  };

  const handleModalClose = () => {
    setViewMode('list');
    setSelectedDriver(null);
  };

  return (
    <div>
      <DriverSearch onSearch={handleSearch} onClear={handleClearSearch} />

      <DriverList
        onCreate={handleCreate}
        onEdit={handleEdit}
        onView={handleView}
        searchQuery={searchQuery}
        searchStatus={searchStatus}
      />

      {/* Модальное окно создания/редактирования */}
      <Modal
        title={viewMode === 'create' ? 'Создание водителя' : 'Редактирование водителя'}
        open={viewMode === 'create' || viewMode === 'edit'}
        onCancel={handleModalClose}
        footer={null}
        width={800}
      >
        <DriverForm
          driver={selectedDriver || undefined}
          onSuccess={handleFormSuccess}
          onCancel={handleModalClose}
        />
      </Modal>
      {/* Модальное окно просмотра */}
      <Modal
        title="Детали водителя"
        open={viewMode === 'view'}
        onCancel={handleModalClose}
        footer={null}
        width={1000}
      >        {selectedDriver && (
          <DriverCard
            driver={selectedDriver}
            onEdit={handleEdit}
            onDelete={handleDelete}
            loading={isLoading}
          />
        )}
      </Modal>
    </div>
  );
};
