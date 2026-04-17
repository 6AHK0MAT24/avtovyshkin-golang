import { apiClient } from '../services/api';
import { EntityService } from '../services/entityService';
import { useEntityCRUD } from './useEntityCRUD';
import type {
  Driver,
  CreateDriverRequest,
  UpdateDriverRequest,
} from '../types/driver';
// Create entity service for drivers
const driverService = new EntityService<Driver, CreateDriverRequest, UpdateDriverRequest>(
  apiClient,
  '/drivers'
);

// Create CRUD hooks for drivers
const driverCRUD = useEntityCRUD<Driver, CreateDriverRequest, UpdateDriverRequest>({
  service: driverService,
  queryKey: ['drivers'],
});

// Export individual hooks for backward compatibility
export const useDrivers = driverCRUD.useGetAll;
export const useDriver = driverCRUD.useGetById;
export const useCreateDriver = driverCRUD.useCreate;
export const useUpdateDriver = driverCRUD.useUpdate;
export const useDeleteDriver = driverCRUD.useDelete;
export const useSearchDrivers = driverCRUD.useSearch;

// File upload hooks for drivers
export const useUploadDriverPhoto = () => driverCRUD.useUploadFile('photo');
export const useUploadDriverLicense = () =>
  driverCRUD.useUploadFile('license', undefined);
export const useUploadDriverPassport = () =>
  driverCRUD.useUploadFile('passport', undefined);

// File delete hooks for drivers
export const useDeleteDriverPhoto = () => driverCRUD.useDeleteFile('photo');
export const useDeleteDriverLicense = () => driverCRUD.useDeleteFile('license');
export const useDeleteDriverPassport = () => driverCRUD.useDeleteFile('passport');

// Export service for direct use if needed
export { driverService };
