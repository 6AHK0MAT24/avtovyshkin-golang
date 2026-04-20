export type VehicleStatus = 'active' | 'inactive' | 'blocked';

export type VehicleType = 'Телескопическая' | 'Телескоп + колено' | 'Телескоп + стрела и рукоять';

export interface Vehicle {
  id: string;
  garageNumber: string;
  vin: string;
  height: number;
  type?: VehicleType;
  power?: number;
  price5?: number;
  price22?: number;
  description?: string;
  brand?: string;
  machine?: string;
  length?: number;
  width?: number;
  heightTs?: number;
  widthWithSupports?: number;
  mass?: number;
  cradleWidthFolded?: number;
  cradleWidthExtended?: number;
  cradleLengthFolded?: number;
  cradleLengthExtended?: number;
  imgArray?: string[];
  mainImageIndex?: number;
  special?: string;
  rostechReg?: boolean;
  status: VehicleStatus;  createdAt: string;
  updatedAt: string;
}

export interface CreateVehicleRequest {
  garageNumber: string;
  vin: string;
  height: number;
  type?: VehicleType;
  power?: number;
  price5?: number;
  price22?: number;
  description?: string;
  brand?: string;
  machine?: string;
  length?: number;
  width?: number;
  heightTs?: number;
  widthWithSupports?: number;
  mass?: number;
  cradleWidthFolded?: number;
  cradleWidthExtended?: number;
  cradleLengthFolded?: number;
  cradleLengthExtended?: number;
  imgArray?: string[];
  mainImageIndex?: number;
  special?: string;
  rostechReg?: boolean;
  status?: VehicleStatus;
}
export interface CreateVehicleRequest {  garageNumber: string;
  vin: string;
  height: number;
  type?: VehicleType;
  power?: number;
  price5?: number;
  price22?: number;
  description?: string;
  brand?: string;
  machine?: string;
  length?: number;
  width?: number;
  heightTs?: number;
  mass?: number;
  cradleWidthFolded?: number;
  cradleWidthExtended?: number;
  cradleLengthFolded?: number;
  cradleLengthExtended?: number;
  imgArray?: string[];
  mainImageIndex?: number;
  special?: string;
  rostechReg?: string;
  status?: VehicleStatus;


export interface UpdateVehicleRequest {
  garageNumber?: string;
  vin?: string;
  height?: number;
  type?: VehicleType;
  power?: number;
  price5?: number;
  price22?: number;
  description?: string;
  brand?: string;
  machine?: string;
  length?: number;
  width?: number;
  heightTs?: number;
  widthWithSupports?: number;
  mass?: number;
  cradleWidthFolded?: number;
  cradleWidthExtended?: number;
  cradleLengthFolded?: number;
  cradleLengthExtended?: number;
  imgArray?: string[];
  mainImageIndex?: number;
  special?: string;
  rostechReg?: boolean;
  status?: VehicleStatus;
}

export interface VehicleListResponse {  vehicles: Vehicle[];
  total: number;
  page: number;
  perPage: number;
}

export interface SearchFilters {
  query?: string;
  status?: VehicleStatus;
  page: number;
  perPage: number;
}

export interface FileUploadResponse {
  filePath: string;
  fileName: string;
  fileSize: number;
}
}
