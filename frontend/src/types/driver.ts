export type DriverStatus = 'active' | 'inactive' | 'blocked';

export interface Driver {
  id: string;
  firstName: string;
  lastName: string;
  middleName?: string;
  phone: string;
  email?: string;
  birthDate?: string;
  photo?: string;
  driverLicenseNumber: string;
  driverLicenseIssueDate: string;
  driverLicenseExpiryDate: string;
  driverLicensePhoto?: string;
  driverLicenseScan?: string;
  passportSeries?: string;
  passportNumber?: string;
  passportIssueDate?: string;
  passportPhoto?: string;
  passportScan?: string;
  address?: string;
  experienceYears: number;
  status: DriverStatus;
  createdAt: string;
  updatedAt: string;
}
export interface CreateDriverRequest {
  firstName: string;
  lastName: string;
  middleName?: string;
  phone: string;
  email?: string;
  birthDate?: string;
  photo?: string;
  driverLicenseNumber: string;
  driverLicenseIssueDate: string;
  driverLicenseExpiryDate: string;
  passportSeries?: string;
  passportNumber?: string;
  passportIssueDate?: string;
  address?: string;
  experienceYears: number;
  status?: DriverStatus;
}
export interface UpdateDriverRequest {
  firstName?: string;
  lastName?: string;
  middleName?: string;
  phone?: string;
  email?: string;
  birthDate?: string;
  photo?: string;
  driverLicenseNumber?: string;
  driverLicenseIssueDate?: string;
  driverLicenseExpiryDate?: string;
  passportSeries?: string;
  passportNumber?: string;
  passportIssueDate?: string;
  address?: string;
  experienceYears?: number;
  status?: DriverStatus;
}
export interface DriverListResponse {
  drivers: Driver[];
  total: number;
  page: number;
  perPage: number;
}

export interface SearchFilters {
  query?: string;
  status?: DriverStatus;
  page: number;
  perPage: number;
}

export interface FileUploadResponse {
  filePath: string;
  fileName: string;
  fileSize: number;
}
