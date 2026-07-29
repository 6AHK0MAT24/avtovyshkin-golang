import axios from 'axios';
import type { AxiosError, AxiosInstance, InternalAxiosRequestConfig, AxiosResponse } from 'axios';

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8082';

class ApiClient {
  protected client: AxiosInstance;

  constructor() {
    this.client = axios.create({
      baseURL: `${API_BASE_URL}/api`,
      headers: {
        'Content-Type': 'application/json',
      },
    });

    // Request interceptor
    this.client.interceptors.request.use(
      (config: InternalAxiosRequestConfig) => {
        // Add auth token if available
        const token = localStorage.getItem('auth_token');
        if (token) {
          config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
      },
      (error: AxiosError) => {
        return Promise.reject(error);
      }
    );

    // Response interceptor
    this.client.interceptors.response.use(
      (response: AxiosResponse) => response,
      (error: AxiosError) => {
        if (error.response?.status === 401) {
          // Handle unauthorized
          localStorage.removeItem('auth_token');
          window.location.href = '/login';
        }
        // Улучшенная обработка ошибок - добавляем текст ответа в error.response.data
        if (error.response?.data === undefined && error.response?.request?.responseText) {
          error.response.data = error.response.request.responseText;
        }
        return Promise.reject(error);
      }
    );
  }

  public get<T = unknown>(url: string, params?: Record<string, unknown>) {
    return this.client.get<T>(url, { params });
  }

  public post<T = unknown>(url: string, data?: unknown) {
    return this.client.post<T>(url, data);
  }

  public put<T = unknown>(url: string, data?: unknown) {
    return this.client.put<T>(url, data);
  }

  public delete<T = unknown>(url: string) {
    return this.client.delete<T>(url);
  }

  public upload<T = unknown>(url: string, formData: FormData) {
    return this.client.post<T>(url, formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });
  }
}

export const apiClient = new ApiClient();

// Vehicles API client for vehicles microservice (port 8081)
const VEHICLES_API_BASE_URL = import.meta.env.VITE_VEHICLES_API_BASE_URL || 'http://localhost:8081';

class VehiclesApiClient extends ApiClient {
  constructor() {
    super();
    // Override the client with vehicles-specific configuration
    this.client = axios.create({
      baseURL: `${VEHICLES_API_BASE_URL}/api`,
      headers: {
        'Content-Type': 'application/json',
      },
    });

    // Request interceptor
    this.client.interceptors.request.use(
      (config: InternalAxiosRequestConfig) => {
        const token = localStorage.getItem('auth_token');
        if (token) {
          config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
      },
      (error: AxiosError) => {
        return Promise.reject(error);
      }
    );

    // Response interceptor
    this.client.interceptors.response.use(
      (response: AxiosResponse) => response,
      (error: AxiosError) => {
        if (error.response?.status === 401) {
          localStorage.removeItem('auth_token');
          window.location.href = '/login';
        }
        // Улучшенная обработка ошибок - добавляем текст ответа в error.response.data
        if (error.response?.data === undefined && error.response?.request?.responseText) {
          error.response.data = error.response.request.responseText;
        }
        return Promise.reject(error);
      }
    );
  }
}

export const vehiclesApiClient = new VehiclesApiClient();

// Clients API client for clients microservice (port 8003)
const CLIENTS_API_BASE_URL = import.meta.env.VITE_CLIENTS_API_BASE_URL || 'http://localhost:8003';

class ClientsApiClient extends ApiClient {
  constructor() {
    super();
    // Override the client with clients-specific configuration
    this.client = axios.create({
      baseURL: `${CLIENTS_API_BASE_URL}/api`,
      headers: {
        'Content-Type': 'application/json',
      },
    });

    // Request interceptor
    this.client.interceptors.request.use(
      (config: InternalAxiosRequestConfig) => {
        const token = localStorage.getItem('auth_token');
        if (token) {
          config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
      },
      (error: AxiosError) => {
        return Promise.reject(error);
      }
    );

    // Response interceptor
    this.client.interceptors.response.use(
      (response: AxiosResponse) => response,
      (error: AxiosError) => {
        if (error.response?.status === 401) {
          localStorage.removeItem('auth_token');
          window.location.href = '/login';
        }
        // Улучшенная обработка ошибок - добавляем текст ответа в error.response.data
        if (error.response?.data === undefined && error.response?.request?.responseText) {
          error.response.data = error.response.request.responseText;
        }
        return Promise.reject(error);
      }
    );
  }
}

export const clientsApiClient = new ClientsApiClient();
