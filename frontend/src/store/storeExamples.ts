/**
 * Примеры использования Zustand stores
 *
 * Этот файл содержит примеры того, как использовать driverStore и uiStore в компонентах
 */

import { useEffect } from 'react';
import {
  useDriverStore,
  useUIStore,
  showSuccessNotification,
  showErrorNotification,
} from './index';
import { driverService } from '../services/driverService';

// ============================================
// Пример 1: Получение списка водителей
// ============================================
export function useFetchDrivers() {
  const { drivers, loading, error, pagination, filters, setDrivers, setLoading, setError } =
    useDriverStore();

  const fetchDrivers = async () => {
    setLoading(true);
    setError(null);
    try {
      const response = await driverService.getDrivers(filters.page, filters.perPage);
      setDrivers(response);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Ошибка загрузки водителей');
      showErrorNotification('Ошибка', 'Не удалось загрузить список водителей');
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchDrivers();
  }, [filters.page, filters.perPage, filters.status]);

  return { drivers, loading, error, pagination, refetch: fetchDrivers };
}

// ============================================
// Пример 2: Создание водителя
// ============================================
export function useCreateDriver() {
  const { addDriver } = useDriverStore();
  const { closeModal } = useUIStore();

  const createDriver = async (data: any) => {
    try {
      const newDriver = await driverService.createDriver(data);
      addDriver(newDriver);
      showSuccessNotification('Успешно', 'Водитель успешно создан');
      closeModal();
      return newDriver;
    } catch (err) {
      const message = err instanceof Error ? err.message : 'Ошибка создания водителя';
      showErrorNotification('Ошибка', message);
      throw err;
    }
  };

  return { createDriver };
}

// ============================================
// Пример 3: Обновление водителя
// ============================================
export function useUpdateDriver() {
  const { updateDriver } = useDriverStore();
  const { closeModal } = useUIStore();

  const updateDriverData = async (id: string, data: any) => {
    try {
      const updatedDriver = await driverService.updateDriver(id, data);
      updateDriver(id, updatedDriver);
      showSuccessNotification('Успешно', 'Данные водителя обновлены');
      closeModal();
      return updatedDriver;
    } catch (err) {
      const message = err instanceof Error ? err.message : 'Ошибка обновления водителя';
      showErrorNotification('Ошибка', message);
      throw err;
    }
  };

  return { updateDriver: updateDriverData };
}

// ============================================
// Пример 4: Удаление водителя
// ============================================
export function useDeleteDriver() {
  const { removeDriver } = useDriverStore();
  const { closeModal } = useUIStore();

  const deleteDriver = async (id: string) => {
    try {
      await driverService.deleteDriver(id);
      removeDriver(id);
      showSuccessNotification('Успешно', 'Водитель удален');
      closeModal();
    } catch (err) {
      const message = err instanceof Error ? err.message : 'Ошибка удаления водителя';
      showErrorNotification('Ошибка', message);
      throw err;
    }
  };

  return { deleteDriver };
}

// ============================================
// Пример 5: Поиск водителей
// ============================================
export function useSearchDrivers() {
  const { setDrivers, setLoading, setError, setFilters } = useDriverStore();

  const searchDrivers = async (query: string, page = 1, perPage = 10) => {
    setLoading(true);
    setError(null);
    setFilters({ query, page, perPage });

    try {
      const response = await driverService.searchDrivers({ query, page, perPage });
      setDrivers(response);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Ошибка поиска');
      showErrorNotification('Ошибка', 'Не удалось выполнить поиск');
    } finally {
      setLoading(false);
    }  };

  return { searchDrivers };
}

// ============================================
// Пример 6: Работа с модальными окнами
// ============================================
export function useModal() {
  const { modal, openModal, closeModal } = useUIStore();

  const openCreateModal = () => openModal('createDriver');
  const openEditModal = (driver: any) => openModal('editDriver', driver);
  const openViewModal = (driver: any) => openModal('viewDriver', driver);
  const openDeleteModal = (driver: any) => openModal('deleteDriver', driver);

  return {
    modal,
    openCreateModal,
    openEditModal,
    openViewModal,
    openDeleteModal,
    closeModal,
  };
}

// ============================================
// Пример 7: Работа с уведомлениями
// ============================================
export function useNotifications() {
  const { notifications, removeNotification, clearNotifications } = useUIStore();

  return {
    notifications,
    removeNotification,
    clearNotifications,
    showSuccess: showSuccessNotification,
    showError: showErrorNotification,
    showWarning: (message: string, description?: string) => {
      useUIStore.getState().showNotification({ type: 'warning', message, description });
    },
    showInfo: (message: string, description?: string) => {
      useUIStore.getState().showNotification({ type: 'info', message, description });
    },
  };
}

// ============================================
// Пример 8: Работа с темой
// ============================================
export function useTheme() {
  const { theme, setTheme } = useUIStore();

  const toggleTheme = () => {
    const newTheme = theme === 'light' ? 'dark' : 'light';
    setTheme(newTheme);
  };

  useEffect(() => {
    // Загружаем тему из localStorage при монтировании
    const savedTheme = localStorage.getItem('theme') as 'light' | 'dark' | 'auto';
    if (savedTheme) {
      setTheme(savedTheme);
    }
  }, []);

  return { theme, setTheme, toggleTheme };
}
