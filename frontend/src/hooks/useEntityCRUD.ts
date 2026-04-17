import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { EntityService, type SearchFilters } from '../services/entityService';
export interface UseEntityCRUDOptions<T, CreateRequest, UpdateRequest> {
  service: EntityService<T, CreateRequest, UpdateRequest>;
  queryKey: string[];
}
export function useEntityCRUD<T, CreateRequest, UpdateRequest>(
  options: UseEntityCRUDOptions<T, CreateRequest, UpdateRequest>
) {
  const { service, queryKey } = options;
  const queryClient = useQueryClient();
  // Get all entities with pagination
  const useGetAll = (page: number = 1, perPage: number = 10) => {
    return useQuery({
      queryKey: [...queryKey, page, perPage],
      queryFn: () => service.getAll(page, perPage),
    });
  };

  // Get entity by ID
  const useGetById = (id: string) => {
    return useQuery({
      queryKey: [...queryKey, id],
      queryFn: () => service.getById(id),
      enabled: !!id,
    });
  };

  // Create entity
  const useCreate = () => {
    return useMutation({
      mutationFn: (data: CreateRequest) => service.create(data),
      onSuccess: () => {
        queryClient.invalidateQueries({ queryKey });
      },
    });
  };

  // Update entity
  const useUpdate = () => {
    return useMutation({
      mutationFn: ({ id, data }: { id: string; data: UpdateRequest }) =>
        service.update(id, data),
      onSuccess: (_, variables) => {
        queryClient.invalidateQueries({ queryKey });
        queryClient.invalidateQueries({ queryKey: [...queryKey, variables.id] });
      },
    });
  };

  // Delete entity
  const useDelete = () => {
    return useMutation({
      mutationFn: (id: string) => service.delete(id),
      onSuccess: () => {
        queryClient.invalidateQueries({ queryKey });
      },
    });
  };

  // Search entities
  const useSearch = (filters: SearchFilters) => {
    return useQuery({
      queryKey: [...queryKey, 'search', filters],
      queryFn: () => service.search(filters),
      enabled: !!filters.query || !!filters.status,
    });
  };

  // Upload file
  const useUploadFile = (endpoint: string = 'photo', additionalData?: Record<string, string>) => {
    return useMutation({
      mutationFn: ({ id, file }: { id: string; file: File }) =>
        service.uploadFile(id, file, endpoint, additionalData),
      onSuccess: (_, variables) => {
        queryClient.invalidateQueries({ queryKey: [...queryKey, variables.id] });
        queryClient.invalidateQueries({ queryKey });
      },
    });
  };

  // Upload multiple files
  const useUploadFiles = (endpoint: string = 'images') => {
    return useMutation({
      mutationFn: ({ id, files }: { id: string; files: File[] }) =>
        service.uploadFiles(id, files, endpoint),
      onSuccess: (_, variables) => {
        queryClient.invalidateQueries({ queryKey: [...queryKey, variables.id] });
        queryClient.invalidateQueries({ queryKey });
      },
    });
  };

  // Delete file
  const useDeleteFile = (endpoint: string = 'photo') => {
    return useMutation({
      mutationFn: (id: string) => service.deleteFile(id, endpoint),
      onSuccess: (_, id) => {
        queryClient.invalidateQueries({ queryKey: [...queryKey, id] });
        queryClient.invalidateQueries({ queryKey });
      },
    });
  };

  // Delete file by index
  const useDeleteFileByIndex = (endpoint: string = 'images') => {
    return useMutation({
      mutationFn: ({ id, index }: { id: string; index: number }) =>
        service.deleteFileByIndex(id, index, endpoint),
      onSuccess: (_, variables) => {
        queryClient.invalidateQueries({ queryKey: [...queryKey, variables.id] });
        queryClient.invalidateQueries({ queryKey });
      },
    });
  };

  // Update entity field
  const useUpdateField = () => {
    return useMutation({
      mutationFn: ({ id, field, value }: { id: string; field: string; value: any }) =>
        service.updateField(id, field, value),
      onSuccess: (_, variables) => {
        queryClient.invalidateQueries({ queryKey: [...queryKey, variables.id] });
        queryClient.invalidateQueries({ queryKey });
      },
    });
  };

  return {
    useGetAll,
    useGetById,
    useCreate,
    useUpdate,
    useDelete,
    useSearch,
    useUploadFile,
    useUploadFiles,
    useDeleteFile,
    useDeleteFileByIndex,
    useUpdateField,
  };
}
