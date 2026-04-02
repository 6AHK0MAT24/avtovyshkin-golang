import React, { useState } from 'react';
import { Input, Select, Button, Space, Card } from 'antd';
import { SearchOutlined, ClearOutlined, PlusOutlined } from '@ant-design/icons';
import type { DriverStatus } from '../../types/driver';

interface DriverSearchProps {
  onSearch: (query: string, status?: DriverStatus) => void;
  onClear: () => void;
  onCreate?: () => void;
}export const DriverSearch: React.FC<DriverSearchProps> = ({ onSearch, onClear, onCreate }) => {
  const [query, setQuery] = useState('');
  const [status, setStatus] = useState<DriverStatus | undefined>();

  const handleSearch = () => {
    onSearch(query, status);
  };

  const handleStatusChange = (value: DriverStatus | undefined) => {
    setStatus(value);
    // Автоматически запускаем поиск при изменении статуса
    onSearch(query, value);
  };

  const handleClear = () => {
    setQuery('');
    setStatus(undefined);
    onClear();
  };
  return (
    <Card style={{ marginBottom: 16 }}>
      <Space size="middle" style={{ width: '100%' }}>
        <Input
          placeholder="Поиск по ФИО, телефону или номеру ВУ"
          value={query}
          onChange={(e) => setQuery(e.target.value)}
          onPressEnter={handleSearch}
          style={{ width: 300 }}
        />
<Select
          placeholder="Статус"
          value={status}
          onChange={handleStatusChange}
          style={{ width: 150 }}
          allowClear
          options={[
            { value: 'active', label: 'Активен' },
            { value: 'inactive', label: 'Неактивен' },
            { value: 'blocked', label: 'Заблокирован' },
          ]}
        />        <Button type="primary" icon={<SearchOutlined />} onClick={handleSearch}>
          Поиск
        </Button>
        <Button icon={<ClearOutlined />} onClick={handleClear}>
          Сбросить
        </Button>
        {onCreate && (
          <Button type="primary" icon={<PlusOutlined />} onClick={onCreate} style={{ backgroundColor: '#52c41a', borderColor: '#52c41a' }}>
            Добавить водителя
          </Button>
        )}
      </Space>
    </Card>
  );
};
