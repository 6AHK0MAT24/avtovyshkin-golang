import React from 'react';
import { Input, Select, Button, Space, Row, Col } from 'antd';
import { SearchOutlined, PlusOutlined, ClearOutlined } from '@ant-design/icons';
import type { VehicleStatus } from '../../types/vehicle';

const { Option } = Select;

interface VehicleSearchProps {
  searchQuery: string;
  searchStatus?: VehicleStatus;
  onSearchChange: (query: string) => void;
  onStatusChange: (status?: VehicleStatus) => void;
  onClear: () => void;
  onAdd: () => void;
}

export const VehicleSearch: React.FC<VehicleSearchProps> = ({
  searchQuery,
  searchStatus,
  onSearchChange,
  onStatusChange,
  onClear,
  onAdd,
}) => {
  return (
    <div style={{ marginBottom: 16 }}>
      <Row gutter={[16, 16]}>
        <Col xs={24} sm={12} md={8} lg={6}>
          <Input
            placeholder="Поиск по гаражному номеру или VIN"
            prefix={<SearchOutlined />}
            value={searchQuery}
            onChange={(e) => onSearchChange(e.target.value)}
            allowClear
          />
        </Col>
        <Col xs={24} sm={12} md={6} lg={4}>
          <Select
            placeholder="Статус"
            value={searchStatus}
            onChange={onStatusChange}
            allowClear
            style={{ width: '100%' }}
          >
            <Option value="active">Активен</Option>
            <Option value="inactive">Неактивен</Option>
            <Option value="blocked">Заблокирован</Option>
          </Select>
        </Col>
        <Col xs={24} sm={24} md={10} lg={14}>
          <Space>
            <Button
              type="primary"
              icon={<PlusOutlined />}
              onClick={onAdd}
            >
              Добавить автовышку
            </Button>
            {(searchQuery || searchStatus) && (
              <Button
                icon={<ClearOutlined />}
                onClick={onClear}
              >
                Очистить
              </Button>
            )}
          </Space>
        </Col>
      </Row>
    </div>
  );
};
