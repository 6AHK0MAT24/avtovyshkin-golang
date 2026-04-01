import React, { useState } from 'react';
import { Input, Select, Button, Space, Card } from 'antd';
import { SearchOutlined, ClearOutlined } from '@ant-design/icons';
import type { DriverStatus } from '../../types/driver';
const { Option } = Select;

interface DriverSearchProps {
  onSearch: (query: string, status?: DriverStatus) => void;
  onClear: () => void;
}

export const DriverSearch: React.FC<DriverSearchProps> = ({ onSearch, onClear }) => {
  const [query, setQuery] = useState('');
  const [status, setStatus] = useState<DriverStatus | undefined>();

  const handleSearch = () => {
    onSearch(query, status);
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
          onChange={setStatus}
          style={{ width: 150 }}
          allowClear
        >
          <Option value="active">Активен</Option>
          <Option value="inactive">Неактивен</Option>
          <Option value="blocked">Заблокирован</Option>
        </Select>
        <Button type="primary" icon={<SearchOutlined />} onClick={handleSearch}>
          Поиск
        </Button>
        <Button icon={<ClearOutlined />} onClick={handleClear}>
          Сбросить
        </Button>
      </Space>
    </Card>
  );
};
