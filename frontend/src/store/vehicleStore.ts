import { create } from 'zustand';
import type { Vehicle, VehicleListResponse, SearchFilters, VehicleStatus } from '../types/vehicle';
interface VehicleState {
  // Состояние
  vehicles: Vehicle[];
  currentVehicle: Vehicle | null;
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
  setVehicles: (response: VehicleListResponse) => void;
  setCurrentVehicle: (vehicle: Vehicle | null) => void;
  addVehicle: (vehicle: Vehicle) => void;
  updateVehicle: (id: string, vehicle: Partial<Vehicle>) => void;
  removeVehicle: (id: string) => void;
  setPagination: (page: number, perPage: number, total: number) => void;
  setFilters: (filters: Partial<SearchFilters>) => void;
  resetFilters: () => void;
  clearError: () => void;
}

const initialFilters: SearchFilters = {
  page: 1,
  perPage: 10,
};

export const useVehicleStore = create<VehicleState>((set) => ({
  // Начальное состояние
  vehicles: [],
  currentVehicle: null,
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

  setVehicles: (response) =>
    set({
      vehicles: response.vehicles,
      pagination: {
        page: response.page,
        perPage: response.perPage,
        total: response.total,
      },
      loading: false,
      error: null,
    }),

  setCurrentVehicle: (vehicle) => set({ currentVehicle: vehicle }),

  addVehicle: (vehicle) =>
    set((state) => ({
      vehicles: [...state.vehicles, vehicle].sort((a, b) =>
        a.garageNumber.localeCompare(b.garageNumber)
      ),
      pagination: {
        ...state.pagination,
        total: state.pagination.total + 1,
      },
    })),

  updateVehicle: (id, updatedVehicle) =>
    set((state) => ({
      vehicles: state.vehicles
        .map((vehicle) =>
          vehicle.id === id ? { ...vehicle, ...updatedVehicle } : vehicle
        )
        .sort((a, b) => a.garageNumber.localeCompare(b.garageNumber)),
      currentVehicle:
        state.currentVehicle?.id === id
          ? { ...state.currentVehicle, ...updatedVehicle }
          : state.currentVehicle,
    })),

  removeVehicle: (id) =>
    set((state) => ({
      vehicles: state.vehicles.filter((vehicle) => vehicle.id !== id),
      currentVehicle:
        state.currentVehicle?.id === id ? null : state.currentVehicle,
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
export const selectVehicles = (state: VehicleState) => state.vehicles;
export const selectCurrentVehicle = (state: VehicleState) => state.currentVehicle;
export const selectVehicleLoading = (state: VehicleState) => state.loading;
export const selectVehicleError = (state: VehicleState) => state.error;
export const selectVehiclePagination = (state: VehicleState) => state.pagination;
export const selectVehicleFilters = (state: VehicleState) => state.filters;
export const selectVehicleById = (id: string) => (state: VehicleState) =>
  state.vehicles.find((vehicle) => vehicle.id === id);
export const selectVehiclesByStatus = (status: VehicleStatus) => (state: VehicleState) =>
  state.vehicles.filter((vehicle) => vehicle.status === status);
