import React from 'react';
import { Card, Space, Button, Typography } from 'antd';
import { EditOutlined, DeleteOutlined } from '@ant-design/icons';
import type { Client } from '../../types/client';
import { StatusBadge } from '../common/StatusBadge';
import { EntityDescription } from '../common/EntityDescription';
import { formatDate, createPhoneItem, createEmailItem, createAddressItem } from '../../utils/entityHelpers';

const { Title } = Typography;

interface ClientCardProps {
  client: Client;
  onEdit: (client: Client) => void;
  onDelete: (id: string) => void;
  loading?: boolean;
}

export const ClientCard = ({ client, onEdit, onDelete, loading = false }: ClientCardProps) => {

  const getClientName = (): string => {
    if (client.clientType === 'individual') {
      return `${client.lastName || ''} ${client.firstName || ''} ${client.middleName || ''}`.trim();
    } else {
      return client.companyName || client.companyLegalName || 'Без названия';
    }
  };

  const getClientInfo = (): Array<{ label: string; value: React.ReactNode }> => {
    const items: Array<{ label: string; value: React.ReactNode }> = [];
    
    if (client.clientType === 'individual') {
      if (client.passportSeries && client.passportNumber) {
        items.push({
          label: 'Паспорт',
          value: `${client.passportSeries} ${client.passportNumber}`,
        });
      }
      if (client.inn) {
        items.push({
          label: 'ИНН',
          value: client.inn,
        });
      }
    } else {
      if (client.innLegal) {
        items.push({
          label: 'ИНН',
          value: client.innLegal,
        });
      }
      if (client.ogrn) {
        items.push({
          label: 'ОГРН',
          value: client.ogrn,
        });
      }
      if (client.directorName) {
        items.push({
          label: 'Директор',
          value: client.directorName,
        });
      }
    }
    
    return items;
  };

  return (
    <Card
      loading={loading}
      title={
        <Space>
          <Title level={4} style={{ margin: 0 }}>
            {getClientName()}
          </Title>
          <StatusBadge status={client.status} />
        </Space>
      }
      extra={
        <Space>
          <Button
            type="primary"
            icon={<EditOutlined />}
            onClick={() => onEdit(client)}
          >
            Редактировать
          </Button>
          <Button
            danger
            icon={<DeleteOutlined />}
            onClick={() => onDelete(client.id)}
          >
            Удалить
          </Button>
        </Space>
      }
    >
      <EntityDescription
        items={[
          createPhoneItem(client.phone),
          ...(client.email ? [createEmailItem(client.email)] : []),
          ...(client.address ? [createAddressItem(client.address)] : []),
          {
            label: 'Тип клиента',
            value: client.clientType === 'individual' ? 'Физическое лицо' : 'Юридическое лицо',
          },
          {
            label: 'Создан',
            value: formatDate(client.createdAt),
          },
          {
            label: 'Обновлен',
            value: formatDate(client.updatedAt),
          },
          ...getClientInfo(),
        ]}
      />
    </Card>
  );
};

export default ClientCard;
