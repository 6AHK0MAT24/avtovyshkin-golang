# План реализации фронтенда для управления автовышками

## Резюме
Создать полноценный CRUD интерфейс для управления автовышками на фронтенде, аналогичный существующему функционалу для водителей. Backend уже готов с необходимыми API endpoints.

## Требования
- **Формат папки изображений**: `uploads/cars/{brand}_{height}`
- **Максимальное количество изображений**: 20 на одну автовышку
- **Максимальный размер изображения**: 10MB
- **Real-time обновления**: через WebSocket (vehicle.created, vehicle.updated, vehicle.deleted)
- **Основное изображение**: отображается в карточке, при клике открывается слайдер со всеми изображениями

## Шаги реализации

### 1. Создать типы TypeScript для Vehicle
**Файл**: `frontend/src/types/vehicle.ts`

Создать интерфейсы:
- `Vehicle` - полная модель автовышки
- `CreateVehicleRequest` - запрос на создание
- `UpdateVehicleRequest` - запрос на обновление
- `VehicleListResponse` - ответ со списком и пагинацией
- `SearchFilters` - фильтры поиска
- `FileUploadResponse` - ответ загрузки файла
- `VehicleStatus` - тип статуса ('active' | 'inactive' | 'blocked')

Поля Vehicle:
- id, garageNumber, vin, height (обязательные)
- type, power, price5, price22, description, brand, machine, length, width, heightTs, widthWithSupports, mass, cradleWidthFolded, cradleWidthExtended, cradleLengthFolded, cradleLengthExtended (опциональные)
- imgArray (string[]), mainImageIndex (number)
- special, rostechReg, status, createdAt, updatedAt

### 2. Создать service для работы с API vehicles
**Файл**: `frontend/src/services/vehicleService.ts`

Методы:
- `getVehicles(page, perPage)` - получить список с пагинацией
- `getVehicle(id)` - получить автовышку по ID
- `createVehicle(data)` - создать автовышку
- `updateVehicle(id, data)` - обновить автовышку
- `deleteVehicle(id)` - удалить автовышку
- `searchVehicles(filters)` - поиск автовышек
- `uploadVehicleImages(id, files)` - загрузить несколько изображений
- `deleteVehicleImage(id, index)` - удалить изображение по индексу
- `setMainImage(id, index)` - установить основное изображение

### 3. Создать store для управления состоянием vehicles
**Файл**: `frontend/src/store/vehicleStore.ts`

Использовать Zustand:
- Состояние: vehicles (map), loading, error
- Методы: setVehicles, addVehicle, updateVehicle, removeVehicle, getVehicleById
- Экспортировать в `frontend/src/store/index.ts`

### 4. Создать hooks для работы с vehicles
**Файл**: `frontend/src/hooks/useVehicles.ts`

Хуки с React Query:
- `useVehicles(page, perPage)` - получить список
- `useVehicle(id)` - получить по ID
- `useCreateVehicle()` - создать
- `useUpdateVehicle()` - обновить
- `useDeleteVehicle()` - удалить
- `useSearchVehicles(filters)` - поиск
- `useUploadVehicleImages()` - загрузка изображений
- `useDeleteVehicleImage()` - удаление изображения
- `useSetMainImage()` - установка основного изображения

### 5. Создать компонент VehicleList
**Файл**: `frontend/src/components/vehicles/VehicleList.tsx`

Функционал:
- Таблица или карточки со списком автовышек
- Пагинация
- Фильтрация по поисковому запросу и статусу
- Кнопки: Просмотр, Редактировать, Удалить
- Отображение: гаражный номер, бренд, высота, статус, основное изображение

### 6. Создать компонент VehicleCard
**Файл**: `frontend/src/components/vehicles/VehicleCard.tsx`

Функционал:
- Детальная информация об автовышке
- Отображение основного изображения (по mainImageIndex)
- При клике на изображение - слайдер со всеми изображениями из imgArray
- Кнопки: Редактировать, Удалить
- Организация информации по секциям:
  - Основная информация (гаражный номер, VIN, бренд, высота, тип, статус)
  - Технические характеристики (мощность, длина, ширина, масса и т.д.)
  - Цены (price5, price22)
  - Описание и особые отметки

### 7. Создать компонент VehicleForm
**Файл**: `frontend/src/components/vehicles/VehicleForm.tsx`

Функционал:
- Форма для создания/редактирования автовышки
- Все поля из модели Vehicle с валидацией
- Организация полей по группам (Row/Col из Ant Design):
  - Строка 1: Гаражный номер, VIN, Высота, Статус
  - Строка 2: Бренд, Тип, Мощность, РостехРег
  - Строка 3: Длина, Ширина, Высота ТС, Ширина с опорами
  - Строка 4: Масса, Ширина люльки (сложена), Ширина люльки (разложена)
  - Строка 5: Длина люльки (сложена), Длина люльки (разложена)
  - Строка 6: Цена (5м), Цена (22м)
  - Строка 7: Описание, Особые отметки
- Компонент VehicleImageUpload для работы с изображениями
- Валидация:
  - Гаражный номер: обязательный
  - VIN: обязательный, 17 символов
  - Высота: обязательная, > 0
  - Тип: один из (Телескопическая, Телескоп + колено, Телескоп + стрела и рукоять)

### 8. Создать компонент VehicleSearch
**Файл**: `frontend/src/components/vehicles/VehicleSearch.tsx`

Функционал:
- Поле поиска по гаражному номеру и VIN
- Фильтр по статусу (active, inactive, blocked)
- Кнопка "Добавить автовышку"
- Кнопка очистки поиска

### 9. Создать компонент VehicleImageUpload
**Файл**: `frontend/src/components/vehicles/VehicleImageUpload.tsx`

Функционал:
- Drag & drop зона для загрузки изображений
- Поддержка множественной загрузки (multiple)
- Валидация:
  - Тип файла: image/* (jpg, jpeg, png, gif, webp)
  - Размер файла: макс 10MB
  - Количество файлов: макс 20
- Предпросмотр загруженных изображений (grid layout)
- Возможность удаления любого изображения (кнопка удаления)
- Выбор основного изображения:
  - Radio button или star icon
  - Визуальное выделение основного изображения
- Отображение текущего количества изображений (X/20)

### 10. Создать компонент ImageSlider
**Файл**: `frontend/src/components/vehicles/ImageSlider.tsx`

Функционал:
- Модальное окно со слайдером изображений
- Навигация: стрелки влево/вправо
- Индикатор текущего изображения (X/Y)
- Кнопка закрытия
- Поддержка клавиш навигации (стрелки, ESC)

### 11. Создать страницу VehiclesPage
**Файл**: `frontend/src/pages/VehiclesPage.tsx`

Функционал:
- Интеграция всех компонентов: VehicleSearch, VehicleList, VehicleForm, VehicleCard
- Управление режимами просмотра: 'list', 'create', 'edit', 'view'
- Обработчики событий:
  - handleCreate - переход к созданию
  - handleEdit - переход к редактированию
  - handleView - переход к просмотру
  - handleDelete - удаление с подтверждением
  - handleFormSuccess - успешное сохранение формы
  - handleSearch - поиск
  - handleClearSearch - очистка поиска
- WebSocket интеграция:
  - vehicle.created - добавить в список, показать уведомление
  - vehicle.updated - обновить в списке, показать уведомление
  - vehicle.deleted - удалить из списка, показать уведомление
- Debounce для предотвращения дубликатов сообщений WebSocket
- Добавить экспорт в `frontend/src/pages/index.ts`

### 12. Добавить маршрут в App.tsx
**Файл**: `frontend/src/App.tsx`

Добавить маршрут:
```tsx
<Route path="/vehicles" element={<VehiclesPage />} />
```

### 13. Добавить пункт меню в Layout.tsx
**Файл**: `frontend/src/components/common/Layout.tsx`

Добавить пункт в menuItems:
```tsx
{
  key: '/vehicles',
  icon: <CarOutlined />,
  label: 'Автовышки',
}
```

### 14. Настроить WebSocket для vehicles
**Файл**: `frontend/src/services/websocket.ts`

Добавить поддержку событий:
- vehicle.created
- vehicle.updated
- vehicle.deleted

**Файл**: `frontend/src/hooks/useVehiclesWebSocket.ts`

Создать хук для подписки на события vehicles:
- useVehicleCreated
- useVehicleUpdated
- useVehicleDeleted

## Технические детали

### Валидация изображений
```typescript
const MAX_IMAGES = 20;
const MAX_FILE_SIZE = 10 * 1024 * 1024; // 10MB
const ALLOWED_TYPES = ['image/jpeg', 'image/jpg', 'image/png', 'image/gif', 'image/webp'];
```

### Структура директорий
```
frontend/src/
├── components/
│   └── vehicles/
│       ├── VehicleList.tsx
│       ├── VehicleCard.tsx
│       ├── VehicleForm.tsx
│       ├── VehicleSearch.tsx
│       ├── VehicleImageUpload.tsx
│       └── ImageSlider.tsx
├── hooks/
│   ├── useVehicles.ts
│   └── useVehiclesWebSocket.ts
├── pages/
│   ├── VehiclesPage.tsx
│   └── index.ts
├── services/
│   └── vehicleService.ts
├── store/
│   ├── vehicleStore.ts
│   └── index.ts
└── types/
    └── vehicle.ts
```

### API Endpoints (уже готовы в backend)
- `GET /api/vehicles` - список с пагинацией
- `GET /api/vehicles/{id}` - получить по ID
- `POST /api/vehicles` - создать
- `PUT /api/vehicles/{id}` - обновить
- `DELETE /api/vehicles/{id}` - удалить
- `GET /api/vehicles/search` - поиск
- `POST /api/vehicles/{id}/images` - загрузить изображения
- `DELETE /api/vehicles/{id}/images/{index}` - удалить изображение
- `PUT /api/vehicles/{id}/main-image` - установить основное изображение
- `GET /uploads/*` - статические файлы изображений

### WebSocket Events (уже готовы в backend)
- `vehicle.created` - новая автовышка создана
- `vehicle.updated` - автовышка обновлена
- `vehicle.deleted` - автовышка удалена

## Риски и решения

### Риск 1: Сложная форма с большим количеством полей
**Решение**: Организовать поля в логические группы с использованием Row/Col, использовать collapse panels для дополнительных полей

### Риск 2: Работа с массивом изображений и индексом основного изображения
**Решение**: Аккуратная синхронизация состояния, валидация индекса при удалении изображений, автоматическая корректировка mainImageIndex

### Риск 3: Real-time обновления через WebSocket - дублирование событий
**Решение**: Использовать debounce механизм (как в DriversPage), отслеживать последние сообщения по типу и ID

### Риск 4: Загрузка больших изображений
**Решение**: Валидация размера на клиенте (10MB), показ прогресса загрузки, обработка ошибок

### Риск 5: Превышение лимита изображений (20)
**Решение**: Валидация количества на клиенте, блокировка загрузки при достижении лимита, показ счетчика X/20

## Порядок реализации

1. Типы и Service (шаги 1-2)
2. Store и Hooks (шаги 3-4)
3. Базовые компоненты (шаги 5-6, 8-10)
4. Форма с загрузкой изображений (шаги 7, 9)
5. Страница и интеграция (шаги 11-14)
6. Тестирование и отладка

## Зависимости

Необходимые пакеты (уже установлены):
- antd (UI компоненты)
- @tanstack/react-query (запросы к API)
- zustand (state management)
- react-router-dom (маршрутизация)
- dayjs (работа с датами)

## Тестирование

После реализации проверить:
1. Создание автовышки с полями
2. Загрузка изображений (множественная)
3. Выбор основного изображения
4. Удаление изображений
5. Редактирование автовышки
6. Удаление автовышки
7. Поиск и фильтрация
8. Real-time обновления через WebSocket
9. Просмотр слайдера изображений
10. Валидация формы и изображений
