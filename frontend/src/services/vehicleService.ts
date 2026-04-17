import { vehiclesApiClient } from './api';
import type {
  Vehicle,
  CreateVehicleRequest,
  UpdateVehicleRequest,
  VehicleListResponse,
  SearchFilters,
  FileUploadResponse,
} from '../types/vehicle';

export const vehicleService = {
  // Get all vehicles with pagination
  getVehicles: async (page: number = 1, perPage: number = 10): Promise<VehicleListResponse> => {
    const response = await vehiclesApiClient.get<VehicleListResponse>('/vehicles', {
      page,
      perPage,
    });
    return response.data;
  },

  // Get vehicle by ID
  getVehicle: async (id: string): Promise<Vehicle> => {
    const response = await vehiclesApiClient.get<Vehicle>(`/vehicles/${id}`);
    return response.data;
  },

  // Create new vehicle
  createVehicle: async (data: CreateVehicleRequest): Promise<Vehicle> => {
    const response = await vehiclesApiClient.post<Vehicle>('/vehicles', data);
    return response.data;
  },

  // Update vehicle
  updateVehicle: async (id: string, data: UpdateVehicleRequest): Promise<Vehicle> => {
    const response = await vehiclesApiClient.put<Vehicle>(`/vehicles/${id}`, data);
    return response.data;
  },

  // Delete vehicle
  deleteVehicle: async (id: string): Promise<void> => {
    await vehiclesApiClient.delete(`/vehicles/${id}`);
  },

  // Search vehicles
  searchVehicles: async (filters: SearchFilters): Promise<VehicleListResponse> => {
    const response = await vehiclesApiClient.get<VehicleListResponse>('/vehicles/search', filters);
    return response.data;
  },

  // Upload vehicle images (multiple files)
  uploadVehicleImages: async (id: string, files: File[]): Promise<FileUploadResponse[]> => {
    const formData = new FormData();
    files.forEach((file) => {
      formData.append('files', file);
    });
    const response = await vehiclesApiClient.upload<FileUploadResponse[]>(`/vehicles/${id}/images`, formData);
    return response.data;
  },

  // Delete vehicle image by index
  deleteVehicleImage: async (id: string, index: number): Promise<void> => {
    await vehiclesApiClient.delete(`/vehicles/${id}/images/${index}`);
  },

  // Set main image
  setMainImage: async (id: string, index: number): Promise<Vehicle> => {
    const response = await vehiclesApiClient.put<Vehicle>(`/vehicles/${id}/main-image`, { index });
    return response.data;
  },};
