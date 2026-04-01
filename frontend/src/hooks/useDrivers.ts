import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { driverService } from '../services/driverService';
import type { CreateDriverRequest, UpdateDriverRequest, SearchFilters } from '../types/driver';
export const useDrivers = (page: number = 1, perPage: number = 10) => {
  return useQuery({
    queryKey: ['drivers', page, perPage],
    queryFn: () => driverService.getDrivers(page, perPage),
  });
};

export const useDriver = (id: string) => {
  return useQuery({
    queryKey: ['driver', id],
    queryFn: () => driverService.getDriver(id),
    enabled: !!id,
  });
};

export const useCreateDriver = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: CreateDriverRequest) => driverService.createDriver(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['drivers'] });
    },
  });
};

export const useUpdateDriver = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: UpdateDriverRequest }) =>
      driverService.updateDriver(id, data),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: ['drivers'] });
      queryClient.invalidateQueries({ queryKey: ['driver', variables.id] });
    },
  });
};

export const useDeleteDriver = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => driverService.deleteDriver(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['drivers'] });
    },
  });
};

export const useSearchDrivers = (filters: SearchFilters) => {
  return useQuery({
    queryKey: ['drivers', 'search', filters],
    queryFn: () => driverService.searchDrivers(filters),
    enabled: !!filters.query || !!filters.status,
  });
};

export const useUploadDriverPhoto = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, file }: { id: string; file: File }) =>
      driverService.uploadDriverPhoto(id, file),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: ['driver', variables.id] });
    },
  });
};

export const useUploadDriverLicense = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, file }: { id: string; file: File }) =>
      driverService.uploadDriverLicense(id, file),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: ['driver', variables.id] });
    },
  });
};

export const useUploadDriverPassport = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, file }: { id: string; file: File }) =>
      driverService.uploadDriverPassport(id, file),
    onSuccess: (_, variables) => {
      queryClient.invalidateQueries({ queryKey: ['driver', variables.id] });
    },
  });
};

export const useDeleteDriverPhoto = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => driverService.deleteDriverPhoto(id),
    onSuccess: (_, id) => {
      queryClient.invalidateQueries({ queryKey: ['driver', id] });
      queryClient.invalidateQueries({ queryKey: ['drivers'] });
    },
  });
};
