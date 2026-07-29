import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { clientService } from '../services/clientService';
import type { CreateClientRequest, UpdateClientRequest } from '../types/client';
import { message } from 'antd';

export const useClients = (type: 'individual' | 'legal_entity', page: number = 1, pageSize: number = 10) => {
  return useQuery({
    queryKey: ['clients', type, page, pageSize],
    queryFn: () => clientService.getClientsByType(type, page, pageSize),
  });
};

export const useClient = (id: string) => {
  return useQuery({
    queryKey: ['client', id],
    queryFn: () => clientService.getClient(id),
    enabled: !!id,
  });
};

export const useCreateClient = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (data: CreateClientRequest) => clientService.createClient(data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['clients'] });
      message.success('Клиент успешно создан');
    },
    onError: () => {
      message.error('Ошибка при создании клиента');
    },
  });
};

export const useUpdateClient = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, data }: { id: string; data: UpdateClientRequest }) =>
      clientService.updateClient(id, data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['clients'] });
      queryClient.invalidateQueries({ queryKey: ['client'] });
      message.success('Клиент успешно обновлен');
    },
    onError: () => {
      message.error('Ошибка при обновлении клиента');
    },
  });
};

export const useDeleteClient = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: string) => clientService.deleteClient(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['clients'] });
      message.success('Клиент успешно удален');
    },
    onError: () => {
      message.error('Ошибка при удалении клиента');
    },
  });
};
