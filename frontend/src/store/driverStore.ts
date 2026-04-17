import { create } from 'zustand';
import type { Driver, DriverListResponse, SearchFilters } from '../types/driver';

interface DriverState {
  // Состояние
  drivers: Driver[];
  currentDriver: Driver | null;
  loading: boolean;
  error: string | null;
  pagination: {
    page: number;
    perPage: number;
    total: number;
  };
  filters: SearchFilters;

  // Actions
  setLoading: (loading: boolean) => void;
  setError: (error: string | null) => void;
  setDrivers: (response: DriverListResponse) => void;  setCurrentDriver: (driver: Driver | null) => void;
  addDriver: (driver: Driver) => void;
  updateDriver: (id: string, driver: Partial<Driver>) => void;
  removeDriver: (id: string) => void;
  setPagination: (page: number, perPage: number, total: number) => void;
  setFilters: (filters: Partial<SearchFilters>) => void;
  resetFilters: () => void;
  clearError: () => void;
}

const initialFilters: SearchFilters = {
  page: 1,
  perPage: 10,
};

export const useDriverStore = create<DriverState>((set) => ({
  // Начальное состояние
  drivers: [],
  currentDriver: null,
  loading: false,
  error: null,
  pagination: {
    page: 1,
    perPage: 10,
    total: 0,
  },
  filters: initialFilters,

  // Actions
  setLoading: (loading) => set({ loading }),

  setError: (error) => set({ error }),

  setDrivers: (response) =>
    set({
      drivers: response.drivers,
      pagination: {
        page: response.page,
        perPage: response.perPage,
        total: response.total,
      },
      loading: false,
      error: null,
    }),

  setCurrentDriver: (driver) => set({ currentDriver: driver }),

  addDriver: (driver) =>
    set((state) => ({
      drivers: [...state.drivers, driver].sort((a, b) => a.lastName.localeCompare(b.lastName)),
      pagination: {
        ...state.pagination,
        total: state.pagination.total + 1,
      },
    })),
  updateDriver: (id, updatedDriver) =>
    set((state) => ({
      drivers: state.drivers
        .map((driver) =>
          driver.id === id ? { ...driver, ...updatedDriver } : driver
        )
        .sort((a, b) => a.lastName.localeCompare(b.lastName)),
      currentDriver:
        state.currentDriver?.id === id
          ? { ...state.currentDriver, ...updatedDriver }
          : state.currentDriver,
    })),
  removeDriver: (id) =>
    set((state) => ({
      drivers: state.drivers.filter((driver) => driver.id !== id),
      currentDriver:
        state.currentDriver?.id === id ? null : state.currentDriver,
      pagination: {
        ...state.pagination,
        total: Math.max(0, state.pagination.total - 1),
      },
    })),

  setPagination: (page, perPage, total) =>
    set({
      pagination: { page, perPage, total },
    }),

  setFilters: (newFilters) =>
    set((state) => ({
      filters: { ...state.filters, ...newFilters },
    })),

  resetFilters: () =>
    set({
      filters: initialFilters,
      pagination: {
        page: 1,
        perPage: 10,
        total: 0,
      },
    }),

  clearError: () => set({ error: null }),
}));

// Селекторы для оптимизации рендеринга
export const selectDrivers = (state: DriverState) => state.drivers;
export const selectCurrentDriver = (state: DriverState) => state.currentDriver;
export const selectDriverLoading = (state: DriverState) => state.loading;
export const selectDriverError = (state: DriverState) => state.error;
export const selectDriverPagination = (state: DriverState) => state.pagination;
export const selectDriverFilters = (state: DriverState) => state.filters;
export const selectDriverById = (id: string) => (state: DriverState) =>
  state.drivers.find((driver) => driver.id === id);
export const selectDriversByStatus = (status: string) => (state: DriverState) =>
  state.drivers.filter((driver) => driver.status === status);
