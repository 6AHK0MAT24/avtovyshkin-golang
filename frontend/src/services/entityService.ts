import { apiClient, vehiclesApiClient } from './api';

export interface EntityListResponse<T> {
  data: T[];
  total: number;
  page: number;
  perPage: number;
}

export interface SearchFilters {
  query?: string;
  status?: string;
  [key: string]: any;
}

export interface FileUploadResponse {
  url: string;
  filename: string;
}

export class EntityService<T, CreateRequest, UpdateRequest> {
  private client: typeof apiClient | typeof vehiclesApiClient;
  private basePath: string;

  constructor(
    client: typeof apiClient | typeof vehiclesApiClient,
    basePath: string
  ) {
    this.client = client;
    this.basePath = basePath;
  }

  // Get all entities with pagination
  async getAll(page: number = 1, perPage: number = 10): Promise<EntityListResponse<T>> {
    const response = await this.client.get<EntityListResponse<T>>(this.basePath, {
      page,
      perPage,
    });
    return response.data;
  }

  // Get entity by ID
  async getById(id: string): Promise<T> {
    const response = await this.client.get<T>(`${this.basePath}/${id}`);
    return response.data;
  }

  // Create new entity
  async create(data: CreateRequest): Promise<T> {
    const response = await this.client.post<T>(this.basePath, data);
    return response.data;
  }

  // Update entity
  async update(id: string, data: UpdateRequest): Promise<T> {
    const response = await this.client.put<T>(`${this.basePath}/${id}`, data);
    return response.data;
  }

  // Delete entity
  async delete(id: string): Promise<void> {
    await this.client.delete(`${this.basePath}/${id}`);
  }

  // Search entities
  async search(filters: SearchFilters): Promise<EntityListResponse<T>> {
    const response = await this.client.get<EntityListResponse<T>>(
      `${this.basePath}/search`,
      filters
    );
    return response.data;
  }

  // Upload file
  async uploadFile(
    id: string,
    file: File,
    endpoint: string = 'photo',
    additionalData?: Record<string, string>
  ): Promise<FileUploadResponse> {
    const formData = new FormData();
    formData.append('file', file);
    
    if (additionalData) {
      Object.entries(additionalData).forEach(([key, value]) => {
        formData.append(key, value);
      });
    }

    const response = await this.client.upload<FileUploadResponse>(
      `${this.basePath}/${id}/${endpoint}`,
      formData
    );
    return response.data;
  }

  // Upload multiple files
  async uploadFiles(
    id: string,
    files: File[],
    endpoint: string = 'images'
  ): Promise<FileUploadResponse[]> {
    const formData = new FormData();
    files.forEach((file) => {
      formData.append('files', file);
    });

    const response = await this.client.upload<FileUploadResponse[]>(
      `${this.basePath}/${id}/${endpoint}`,
      formData
    );
    return response.data;
  }

  // Delete file
  async deleteFile(id: string, endpoint: string): Promise<void> {
    await this.client.delete(`${this.basePath}/${id}/${endpoint}`);
  }

  // Delete file by index
  async deleteFileByIndex(id: string, index: number, endpoint: string = 'images'): Promise<void> {
    await this.client.delete(`${this.basePath}/${id}/${endpoint}/${index}`);
  }

  // Update entity field
  async updateField(id: string, field: string, value: any): Promise<T> {
    const response = await this.client.put<T>(`${this.basePath}/${id}/${field}`, { value });
    return response.data;
  }
}
