import { apiClient } from './api';
import type {
  Driver,
  CreateDriverRequest,
  UpdateDriverRequest,
  DriverListResponse,
  SearchFilters,
  FileUploadResponse,
} from '../types/driver';
export const driverService = {
  // Get all drivers with pagination
  getDrivers: async (page: number = 1, perPage: number = 10): Promise<DriverListResponse> => {
    const response = await apiClient.get<DriverListResponse>('/drivers', {
      page,
      perPage,
    });
    return response.data;
  },

  // Get driver by ID
  getDriver: async (id: string): Promise<Driver> => {
    const response = await apiClient.get<Driver>(`/drivers/${id}`);
    return response.data;
  },

  // Create new driver
  createDriver: async (data: CreateDriverRequest): Promise<Driver> => {
    const response = await apiClient.post<Driver>('/drivers', data);
    return response.data;
  },

  // Update driver
  updateDriver: async (id: string, data: UpdateDriverRequest): Promise<Driver> => {
    const response = await apiClient.put<Driver>(`/drivers/${id}`, data);
    return response.data;
  },

  // Delete driver
  deleteDriver: async (id: string): Promise<void> => {
    await apiClient.delete(`/drivers/${id}`);
  },

  // Search drivers
  searchDrivers: async (filters: SearchFilters): Promise<DriverListResponse> => {
    const response = await apiClient.get<DriverListResponse>('/drivers/search', filters);
    return response.data;
  },

  // Upload driver photo
  uploadDriverPhoto: async (id: string, file: File): Promise<FileUploadResponse> => {
    const formData = new FormData();
    formData.append('file', file);
    const response = await apiClient.upload<FileUploadResponse>(`/drivers/${id}/photo`, formData);
    return response.data;
  },

  // Upload driver license
  uploadDriverLicense: async (id: string, file: File): Promise<FileUploadResponse> => {
    const formData = new FormData();
    formData.append('file', file);
    const response = await apiClient.upload<FileUploadResponse>(`/drivers/${id}/license`, formData);
    return response.data;
  },

  // Upload driver passport
  uploadDriverPassport: async (id: string, file: File): Promise<FileUploadResponse> => {
    const formData = new FormData();
    formData.append('file', file);
    const response = await apiClient.upload<FileUploadResponse>(`/drivers/${id}/passport`, formData);
    return response.data;
  },

  // Delete driver photo
  deleteDriverPhoto: async (id: string): Promise<void> => {
    await apiClient.delete(`/drivers/${id}/photo`);
  },
};
