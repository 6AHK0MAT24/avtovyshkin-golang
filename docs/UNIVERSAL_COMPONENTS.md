# Универсальные компоненты для CRUD операций

## Обзор

В проекте реализованы универсальные компоненты для работы с CRUD операциями, которые могут быть переиспользованы для различных сущностей (водители, автовышки и т.д.).

## Фронтенд

### 1. EntityService (`frontend/src/services/entityService.ts`)

Универсальный сервис для работы с API endpoints.

**Основные методы:**
- `getAll(page, perPage)` - получить все сущности с пагинацией
- `getById(id)` - получить сущность по ID
- `create(data)` - создать новую сущность
- `update(id, data)` - обновить сущность
- `delete(id)` - удалить сущность
- `search(filters)` - поиск сущностей
- `uploadFile(id, file, endpoint, additionalData)` - загрузка файла
- `uploadFiles(id, files, endpoint)` - загрузка нескольких файлов
- `deleteFile(id, endpoint)` - удаление файла
- `deleteFileByIndex(id, index, endpoint)` - удаление файла по индексу
- `updateField(id, field, value)` - обновление поля сущности

**Пример использования:**
```typescript
import { apiClient } from './api';
import { EntityService } from './entityService';

const driverService = new EntityService<Driver, CreateDriverRequest, UpdateDriverRequest>(
  apiClient,
  '/drivers'
);

// Получить всех водителей
const drivers = await driverService.getAll(1, 10);

// Создать водителя
const newDriver = await driverService.create({ name: 'Иван', license: '12345' });
```

### 2. useEntityCRUD (`frontend/src/hooks/useEntityCRUD.ts`)

Универсальный React Query хук для работы с сущностями.

**Основные хуки:**
- `useGetAll(page, perPage)` - получить все сущности
- `useGetById(id)` - получить сущность по ID
- `useCreate()` - создать сущность
- `useUpdate()` - обновить сущность
- `useDelete()` - удалить сущность
- `useSearch(filters)` - поиск сущностей
- `useUploadFile(endpoint, additionalData)` - загрузка файла
- `useUploadFiles(endpoint)` - загрузка нескольких файлов
- `useDeleteFile(endpoint)` - удаление файла
- `useDeleteFileByIndex(endpoint)` - удаление файла по индексу
- `useUpdateField()` - обновление поля

**Пример использования:**
```typescript
import { useEntityCRUD } from './useEntityCRUD';
import { EntityService } from './entityService';

const driverService = new EntityService<Driver, CreateDriverRequest, UpdateDriverRequest>(
  apiClient,
  '/drivers'
);

const driverCRUD = useEntityCRUD<Driver, CreateDriverRequest, UpdateDriverRequest>({
  entityName: 'driver',
  service: driverService,
  queryKey: ['drivers'],
});

// В компоненте
const { data: drivers, isLoading } = driverCRUD.useGetAll(1, 10);
const createDriver = driverCRUD.useCreate();

const handleCreate = () => {
  createDriver.mutate({ name: 'Иван', license: '12345' });
};
```

### 3. Универсальные хуки для водителей (`frontend/src/hooks/useDriversUniversal.ts`)

Готовые хуки для работы с водителями, основанные на универсальных компонентах.

**Экспортируемые хуки:**
- `useDrivers(page, perPage)` - получить всех водителей
- `useDriver(id)` - получить водителя по ID
- `useCreateDriver()` - создать водителя
- `useUpdateDriver()` - обновить водителя
- `useDeleteDriver()` - удалить водителя
- `useSearchDrivers(filters)` - поиск водителей
- `useUploadDriverPhoto()` - загрузить фото водителя
- `useUploadDriverLicense()` - загрузить лицензию водителя
- `useUploadDriverPassport()` - загрузить паспорт водителя
- `useDeleteDriverPhoto()` - удалить фото водителя
- `useDeleteDriverLicense()` - удалить лицензию водителя
- `useDeleteDriverPassport()` - удалить паспорт водителя

**Пример использования:**
```typescript
import { useDrivers, useCreateDriver } from './hooks/useDriversUniversal';

function DriversList() {
  const { data: drivers, isLoading } = useDrivers(1, 10);
  const createDriver = useCreateDriver();

  if (isLoading) return <div>Загрузка...</div>;

  return (
    <div>
      {drivers?.data.map(driver => (
        <div key={driver.id}>{driver.name}</div>
      ))}
    </div>
  );
}
```

### 4. Универсальные хуки для автовышек (`frontend/src/hooks/useVehiclesUniversal.ts`)

Готовые хуки для работы с автовышками, основанные на универсальных компонентах.

**Экспортируемые хуки:**
- `useVehicles(page, perPage)` - получить все автовышки
- `useVehicle(id)` - получить автовышку по ID
- `useCreateVehicle()` - создать автовышку
- `useUpdateVehicle()` - обновить автовышку
- `useDeleteVehicle()` - удалить автовышку
- `useSearchVehicles(filters)` - поиск автовышек
- `useUploadVehicleImages()` - загрузить изображения автовышки
- `useDeleteVehicleImage()` - удалить изображение автовышки
- `useSetMainImage()` - установить главное изображение

**Пример использования:**
```typescript
import { useVehicles, useCreateVehicle } from './hooks/useVehiclesUniversal';

function VehiclesList() {
  const { data: vehicles, isLoading } = useVehicles(1, 10);
  const createVehicle = useCreateVehicle();

  if (isLoading) return <div>Загрузка...</div>;

  return (
    <div>
      {vehicles?.data.map(vehicle => (
        <div key={vehicle.id}>{vehicle.garageNumber}</div>
      ))}
    </div>
  );
}
```

## Бекенд

### 1. BaseRepository (`backend/vehicles-microservice/internal/repository/base_repository.go`)

Универсальный репозиторий для работы с базой данных.

**Основные методы:**
- `Create(ctx, entity)` - создать сущность
- `GetByID(ctx, id, dest)` - получить сущность по ID
- `GetAll(ctx, limit, offset, dest)` - получить все сущности с пагинацией
- `Update(ctx, entity)` - обновить сущность
- `Delete(ctx, id)` - удалить сущность
- `Search(ctx, query, limit, offset, dest)` - поиск сущностей

**Пример использования:**
```go
import "vehicles-service/internal/repository"

type DriverRepository struct {
    *repository.BaseRepository
}

func NewDriverRepository(db *sqlx.DB) *DriverRepository {
    return &DriverRepository{
        BaseRepository: repository.NewBaseRepository(db, "drivers"),
    }
}
```

### 2. BaseService (`backend/vehicles-microservice/internal/service/base_service.go`)

Универсальный сервис для бизнес-логики.

**Основные методы:**
- `Create(ctx, req)` - создать сущность
- `GetByID(ctx, id)` - получить сущность по ID
- `GetAll(ctx, page, pageSize)` - получить все сущности с пагинацией
- `Update(ctx, id, req)` - обновить сущность
- `Delete(ctx, id)` - удалить сущность
- `Search(ctx, query, page, pageSize)` - поиск сущностей

**Пример использования:**
```go
import "vehicles-service/internal/service"

type DriverService struct {
    *service.BaseService
    repo repository.DriverRepository
}

func NewDriverService(repo repository.DriverRepository) *DriverService {
    return &DriverService{
        BaseService: service.NewBaseService(repo),
        repo:        repo,
    }
}
```

### 3. BaseHandler (`backend/vehicles-microservice/internal/handler/base_handler.go`)

Универсальный HTTP handler для обработки запросов.

**Основные методы:**
- `Create(w, r, entityName, createFunc)` - обработать POST запрос
- `GetByID(w, r, entityName, getFunc)` - обработать GET запрос по ID
- `GetAll(w, r, entityName, getAllFunc)` - обработать GET запрос для всех сущностей
- `Update(w, r, entityName, updateFunc)` - обработать PUT запрос
- `Delete(w, r, entityName, deleteFunc)` - обработать DELETE запрос
- `Search(w, r, entityName, searchFunc)` - обработать поиск

**Пример использования:**
```go
import "vehicles-service/internal/handler"

type DriverHandler struct {
    *handler.BaseHandler
    service service.DriverService
}

func NewDriverHandler(service service.DriverService) *DriverHandler {
    return &DriverHandler{
        BaseHandler: handler.NewBaseHandler(service),
        service:     service,
    }
}

func (h *DriverHandler) CreateDriver(w http.ResponseWriter, r *http.Request) {
    h.Create(w, r, "driver", h.service.CreateDriver)
}
```

## Преимущества универсальных компонентов

1. **Переиспользование**: Один и тот же код можно использовать для разных сущностей
2. **Поддерживаемость**: Изменения в логике CRUD операций нужно вносить только в одном месте
3. **Согласованность**: Все сущности работают одинаково, что упрощает понимание кода
4. **Тестирование**: Универсальные компоненты легче тестировать
5. **Расширяемость**: Легко добавить новые сущности, используя существующие компоненты

## Миграция существующего кода

### Фронтенд

Для миграции существующих хуков на универсальные компоненты:

1. Замените импорты:
```typescript
// Было
import { useDrivers } from './hooks/useDrivers';

// Стало
import { useDrivers } from './hooks/useDriversUniversal';
```

2. API остаётся тем же самым, поэтому изменения в компонентах не требуются

### Бекенд

Для миграции существующих handlers, services и repositories:

1. Создайте новые реализации на основе базовых классов
2. Перенесите специфичную логику в новые классы
3. Обновите маршрутизацию для использования новых handlers

## Следующие шаги

1. Протестировать универсальные компоненты с водителями
2. После успешного тестирования подключить к автовышкам
3. Удалить старые дублирующиеся файлы
4. Обновить документацию
