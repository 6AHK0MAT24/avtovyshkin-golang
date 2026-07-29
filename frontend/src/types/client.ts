export type ClientType = 'individual' | 'legal_entity';
export type ClientStatus = 'active' | 'inactive' | 'blocked';

export interface Client {
  id: string;
  clientType: ClientType;
  phone: string;
  email?: string;
  address?: string;
  status: ClientStatus;
  notes?: string;

  // Поля для физических лиц
  firstName?: string;
  lastName?: string;
  middleName?: string;
  birthDate?: string;
  passportSeries?: string;
  passportNumber?: string;
  passportIssueDate?: string;
  passportIssuedBy?: string;
  inn?: string;

  // Поля для юридических лиц
  companyName?: string;
  companyLegalName?: string;
  ogrn?: string;
  innLegal?: string;
  kpp?: string;
  legalAddress?: string;
  actualAddress?: string;
  bankName?: string;
  bic?: string;
  accountNumber?: string;
  correspondentAccount?: string;
  directorName?: string;
  directorPosition?: string;
  contactPerson?: string;
  contactPersonPhone?: string;
  contactPersonEmail?: string;

  createdAt: string;
  updatedAt: string;
}

export interface CreateClientRequest {
  clientType: ClientType;
  phone: string;
  email?: string;
  address?: string;
  status?: ClientStatus;
  notes?: string;

  // Поля для физических лиц
  firstName?: string;
  lastName?: string;
  middleName?: string;
  birthDate?: string;
  passportSeries?: string;
  passportNumber?: string;
  passportIssueDate?: string;
  passportIssuedBy?: string;
  inn?: string;

  // Поля для юридических лиц
  companyName?: string;
  companyLegalName?: string;
  ogrn?: string;
  innLegal?: string;
  kpp?: string;
  legalAddress?: string;
  actualAddress?: string;
  bankName?: string;
  bic?: string;
  accountNumber?: string;
  correspondentAccount?: string;
  directorName?: string;
  directorPosition?: string;
  contactPerson?: string;
  contactPersonPhone?: string;
  contactPersonEmail?: string;
}

export interface UpdateClientRequest {
  phone?: string;
  email?: string;
  address?: string;
  status?: ClientStatus;
  notes?: string;

  // Поля для физических лиц
  firstName?: string;
  lastName?: string;
  middleName?: string;
  birthDate?: string;
  passportSeries?: string;
  passportNumber?: string;
  passportIssueDate?: string;
  passportIssuedBy?: string;
  inn?: string;

  // Поля для юридических лиц
  companyName?: string;
  companyLegalName?: string;
  ogrn?: string;
  innLegal?: string;
  kpp?: string;
  legalAddress?: string;
  actualAddress?: string;
  bankName?: string;
  bic?: string;
  accountNumber?: string;
  correspondentAccount?: string;
  directorName?: string;
  directorPosition?: string;
  contactPerson?: string;
  contactPersonPhone?: string;
  contactPersonEmail?: string;
}

export interface ClientListResponse {
  clients: Client[];
  total: number;
  page: number;
  perPage: number;
}

export interface SearchFilters {
  query?: string;
  clientType?: ClientType;
  status?: ClientStatus;
  page: number;
  perPage: number;
}
