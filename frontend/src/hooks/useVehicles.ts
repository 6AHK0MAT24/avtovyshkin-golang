import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { vehicleService } from '../services/vehicleService';
import type { CreateVehicleRequest, UpdateVehicleRequest, SearchFilters } from '../types/vehicle';

export const useVehicles = (page: number = 1, perPage: number = 10) => {
  return useQuery({
    queryKey: ['vehicles', page, perPage],
    queryFn: () => vehicleService.getVehicles(page, perPage),
  });
};

export const useVehicle = (id: string) => {
  return useQuery({
    queryKey: ['vehicle', id],
    queryFn: () => vehicleService.getVehicle(id),
    enabled: !!id,
  });
};

export const useCreateVehicle = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: CreateVehicleRequest) => vehicleService.createVehicle(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['vehicles'] });
    },
  });
};

export const useUpdateVehicle = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: UpdateVehicleRequest }) =>
      vehicleService.updateVehicle(id, data),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: ['vehicles'] });
      queryClient.invalidateQueries({ queryKey: ['vehicle', variables.id] });
    },
  });
};

export const useDeleteVehicle = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => vehicleService.deleteVehicle(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['vehicles'] });
    },
  });
};

export const useSearchVehicles = (filters: SearchFilters) => {
  return useQuery({
    queryKey: ['vehicles', 'search', filters],
    queryFn: () => vehicleService.searchVehicles(filters),
    enabled: !!filters.query || !!filters.status,
  });
};

export const useUploadVehicleImages = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, files }: { id: string; files: File[] }) =>
      vehicleService.uploadVehicleImages(id, files),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: ['vehicle', variables.id] });
      queryClient.invalidateQueries({ queryKey: ['vehicles'] });
    },
  });
};

export const useDeleteVehicleImage = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, index }: { id: string; index: number }) =>
      vehicleService.deleteVehicleImage(id, index),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: ['vehicle', variables.id] });
      queryClient.invalidateQueries({ queryKey: ['vehicles'] });
    },
  });
};

export const useSetMainImage = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, index }: { id: string; index: number }) =>
      vehicleService.setMainImage(id, index),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: ['vehicle', variables.id] });
      queryClient.invalidateQueries({ queryKey: ['vehicles'] });
    },
  });
};
