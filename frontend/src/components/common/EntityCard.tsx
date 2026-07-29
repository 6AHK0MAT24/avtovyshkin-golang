import React from 'react';
import { Card, Space, Button, Typography } from 'antd';
import { EditOutlined, DeleteOutlined } from '@ant-design/icons';
import { StatusBadge, type StatusType } from './StatusBadge';

const { Title } = Typography;

export interface EntityCardProps<T> {
  entity: T;
  title: string;
  status?: StatusType;
  onEdit?: (entity: T) => void;
  onDelete?: (id: string) => void;
  loading?: boolean;
  icon?: React.ReactNode;
  extra?: React.ReactNode;
  children?: React.ReactNode;
}

export function EntityCard<T extends { id: string }>({
  entity,
  title,
  status,
  onEdit,
  onDelete,
  loading = false,
  icon,
  extra,
  children,
}: EntityCardProps<T>) {
  return (
    <Card
      loading={loading}
      title={
        <Space>
          {icon}
          <Title level={4} style={{ margin: 0 }}>
            {title}
          </Title>
          {status && <StatusBadge status={status} />}
        </Space>
      }
      extra={
        <Space>
          {onEdit && (
            <Button type="primary" icon={<EditOutlined />} onClick={() => onEdit(entity)}>
              Редактировать
            </Button>
          )}
          {onDelete && (
            <Button danger icon={<DeleteOutlined />} onClick={() => onDelete(entity.id)}>
              Удалить
            </Button>
          )}
          {extra}
        </Space>
      }
    >
      {children}
    </Card>
  );
}

export default EntityCard;