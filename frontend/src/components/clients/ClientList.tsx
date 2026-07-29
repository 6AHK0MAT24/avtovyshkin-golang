import React from 'react';
import { Table, Button, Space, Tag, Popconfirm } from 'antd';
import { EditOutlined, DeleteOutlined } from '@ant-design/icons';
import type { Client, ClientStatus } from '../../types/client';
import type { ColumnsType } from 'antd/es/table';

interface ClientListProps {
  clients: Client[];
  loading?: boolean;
  onEdit: (client: Client) => void;
  onDelete: (id: string) => void;
}

const statusColors: Record<ClientStatus, string> = {
  active: 'green',
  inactive: 'orange',
  blocked: 'red',
};

const statusLabels: Record<ClientStatus, string> = {
  active: 'Активен',
  inactive: 'Неактивен',
  blocked: 'Заблокирован',
};

export const ClientList: React.FC<ClientListProps> = ({
  clients,
  loading = false,
  onEdit,
  onDelete,
}) => {
  const getClientName = (client: Client) => {
    if (client.clientType === 'individual') {
      return `${client.lastName || ''} ${client.firstName || ''} ${client.middleName || ''}`.trim();
    } else {
      return client.companyName || client.companyLegalName || 'Без названия';
    }
  };

  const getClientInfo = (client: Client) => {
    if (client.clientType === 'individual') {
      const info = [];
      if (client.passportSeries && client.passportNumber) {
        info.push(`Паспорт: ${client.passportSeries} ${client.passportNumber}`);
      }
      if (client.inn) {
        info.push(`ИНН: ${client.inn}`);
      }
      return info.join(' | ');
    } else {
      const info = [];
      if (client.innLegal) {
        info.push(`ИНН: ${client.innLegal}`);
      }
      if (client.ogrn) {
        info.push(`ОГРН: ${client.ogrn}`);
      }
      if (client.directorName) {
        info.push(`Директор: ${client.directorName}`);
      }
      return info.join(' | ');
    }
  };

  const columns: ColumnsType<Client> = [
    {
      title: 'Название',
      dataIndex: 'id',
      key: 'name',
      render: (_, record) => (
        <div>
          <div className="font-semibold">{getClientName(record)}</div>
          <div className="text-xs text-gray-500">{getClientInfo(record)}</div>
        </div>
      ),
    },
    {
      title: 'Телефон',
      dataIndex: 'phone',
      key: 'phone',
    },
    {
      title: 'Email',
      dataIndex: 'email',
      key: 'email',
      render: (email) => email || '-',
    },
    {
      title: 'Адрес',
      dataIndex: 'address',
      key: 'address',
      render: (address) => address || '-',
      ellipsis: true,
    },
    {
      title: 'Статус',
      dataIndex: 'status',
      key: 'status',
      render: (status: ClientStatus) => (
        <Tag color={statusColors[status]}>{statusLabels[status]}</Tag>
      ),
    },
    {
      title: 'Действия',
      key: 'actions',
      width: 150,
      render: (_, record) => (
        <Space size="small">
          <Button
            type="link"
            icon={<EditOutlined />}
            onClick={() => onEdit(record)}
          >
            Редактировать
          </Button>
          <Popconfirm
            title="Удалить клиента?"
            description="Это действие нельзя отменить"
            onConfirm={() => onDelete(record.id)}
            okText="Да"
            cancelText="Нет"
          >
            <Button type="link" danger icon={<DeleteOutlined />}>
              Удалить
            </Button>
          </Popconfirm>
        </Space>
      ),
    },
  ];

  return (
    <Table
      columns={columns}
      dataSource={clients}
      rowKey="id"
      loading={loading}
      pagination={{
        showSizeChanger: true,
        showTotal: (total) => `Всего: ${total}`,
      }}
    />
  );
};
