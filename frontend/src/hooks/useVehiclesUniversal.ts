import { vehiclesApiClient } from '../services/api';
import { EntityService } from '../services/entityService';
import { useEntityCRUD } from './useEntityCRUD';
import type {
  Vehicle,
  CreateVehicleRequest,
  UpdateVehicleRequest,
} from '../types/vehicle';
// Create entity service for vehicles
const vehicleService = new EntityService<Vehicle, CreateVehicleRequest, UpdateVehicleRequest>(
  vehiclesApiClient,
  '/vehicles'
);

// Create CRUD hooks for vehicles
const vehicleCRUD = useEntityCRUD<Vehicle, CreateVehicleRequest, UpdateVehicleRequest>({
  service: vehicleService,
  queryKey: ['vehicles'],
});
// Export individual hooks for backward compatibility
export const useVehicles = vehicleCRUD.useGetAll;
export const useVehicle = vehicleCRUD.useGetById;
export const useCreateVehicle = vehicleCRUD.useCreate;
export const useUpdateVehicle = vehicleCRUD.useUpdate;
export const useDeleteVehicle = vehicleCRUD.useDelete;
export const useSearchVehicles = vehicleCRUD.useSearch;

// File upload hooks for vehicles
export const useUploadVehicleImages = () => vehicleCRUD.useUploadFiles('images');

// File delete hooks for vehicles
export const useDeleteVehicleImage = () => vehicleCRUD.useDeleteFileByIndex('images');

// Field update hooks for vehicles
export const useSetMainImage = () => vehicleCRUD.useUpdateField();

// Export service for direct use if needed
export { vehicleService };
