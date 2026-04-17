import React from 'react';
import { useDrivers, useCreateDriver, useUpdateDriver, useDeleteDriver } from '../hooks/useDriversUniversal';
import { DriverForm } from '../components/drivers/DriverForm';
import { DriverList } from '../components/drivers/DriverList';
import { DriverSearch } from '../components/drivers/DriverSearch';
import type { CreateDriverRequest, UpdateDriverRequest, SearchFilters } from '../types/driver';

export const DriversPageUniversal: React.FC = () => {
  const [page, setPage] = React.useState(1);
  const [searchFilters, setSearchFilters] = React.useState<SearchFilters>({});
  const [editingDriver, setEditingDriver] = React.useState<string | null>(null);

  const { data: driversData, isLoading, error } = useDrivers(page, 10);
  const createDriver = useCreateDriver();
  const updateDriver = useUpdateDriver();
  const deleteDriver = useDeleteDriver();

  const handleCreate = (data: CreateDriverRequest) => {
    createDriver.mutate(data, {
      onSuccess: () => {
        console.log('Водитель успешно создан');
      },
    });
  };

  const handleUpdate = (id: string, data: UpdateDriverRequest) => {
    updateDriver.mutate({ id, data }, {
      onSuccess: () => {
        console.log('Водитель успешно обновлен');
        setEditingDriver(null);
      },
    });
  };

  const handleDelete = (id: string) => {
    if (window.confirm('Вы уверены, что хотите удалить этого водителя?')) {
      deleteDriver.mutate(id, {
        onSuccess: () => {
          console.log('Водитель успешно удален');
        },
      });
    }
  };

  const handleSearch = (filters: SearchFilters) => {
    setSearchFilters(filters);
    setPage(1);
  };

  if (isLoading) {
    return <div className="text-center py-8">Загрузка водителей...</div>;
  }

  if (error) {
    return <div className="text-center py-8 text-red-500">Ошибка загрузки водителей</div>;
  }

  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="text-3xl font-bold mb-6">Управление водителями (Универсальный компонент)</h1>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <div className="lg:col-span-1">
          <DriverForm
            onSubmit={editingDriver ? (data) => handleUpdate(editingDriver, data) : handleCreate}
            onCancel={() => setEditingDriver(null)}
            isEditing={!!editingDriver}
          />
        </div>

        <div className="lg:col-span-2">
          <DriverSearch onSearch={handleSearch} />

          {driversData && (
            <>
              <DriverList
                drivers={driversData.data || []}
                onEdit={setEditingDriver}
                onDelete={handleDelete}
              />

              {driversData.total && driversData.total > 10 && (
                <div className="flex justify-center mt-6 gap-2">
                  <button
                    onClick={() => setPage(p => Math.max(1, p - 1))}
                    disabled={page === 1}
                    className="px-4 py-2 bg-blue-500 text-white rounded disabled:bg-gray-300"
                  >
                    Предыдущая
                  </button>
                  <span className="px-4 py-2">
                    Страница {page} из {Math.ceil((driversData.total || 0) / 10)}
                  </span>
                  <button
                    onClick={() => setPage(p => p + 1)}
                    disabled={page >= Math.ceil((driversData.total || 0) / 10)}
                    className="px-4 py-2 bg-blue-500 text-white rounded disabled:bg-gray-300"
                  >
                    Следующая
                  </button>
                </div>
              )}
            </>
          )}
        </div>
      </div>
    </div>
  );
};
