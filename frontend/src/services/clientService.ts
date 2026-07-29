import { clientsApiClient } from './api';
import type {
  Client,
  CreateClientRequest,
  UpdateClientRequest,
  ClientListResponse,
  SearchFilters,
} from '../types/client';

export const clientService = {
  // Get all clients with pagination
  getClients: async (page: number = 1, perPage: number = 10): Promise<ClientListResponse> => {
    const response = await clientsApiClient.get<ClientListResponse>('/clients', {
      page,
      perPage,
    });
    return response.data;
  },

  // Get clients by type
  getClientsByType: async (
    type: 'individual' | 'legal_entity',
    page: number = 1,
    perPage: number = 10
  ): Promise<ClientListResponse> => {
    const response = await clientsApiClient.get<ClientListResponse>(`/clients/type/${type}`, {
      page,
      perPage,
    });
    return response.data;
  },

  // Get client by ID
  getClient: async (id: string): Promise<Client> => {
    const response = await clientsApiClient.get<Client>(`/clients/${id}`);
    return response.data;
  },

  // Create new client
  createClient: async (data: CreateClientRequest): Promise<Client> => {
    const response = await clientsApiClient.post<Client>('/clients', data);
    return response.data;
  },

  // Update client
  updateClient: async (id: string, data: UpdateClientRequest): Promise<Client> => {
    const response = await clientsApiClient.put<Client>(`/clients/${id}`, data);
    return response.data;
  },

  // Delete client
  deleteClient: async (id: string): Promise<void> => {
    await clientsApiClient.delete(`/clients/${id}`);
  },

  // Search clients
  searchClients: async (filters: SearchFilters): Promise<ClientListResponse> => {
    const response = await clientsApiClient.get<ClientListResponse>('/clients/search', filters);
    return response.data;
  },
};
