import React, { useState } from 'react';
import { Table, Button, Space, Tag, Popconfirm, message, Image } from 'antd';
import { PlusOutlined, EditOutlined, DeleteOutlined } from '@ant-design/icons';
import { useDrivers, useDeleteDriver } from '../../hooks/useDrivers';
import type { Driver, DriverStatus } from '../../types/driver';
import type { ColumnsType } from 'antd/es/table';interface DriverListProps {
  onCreate: () => void;
  onEdit: (driver: Driver) => void;
  onView?: (driver: Driver) => void;
  searchQuery?: string;
  searchStatus?: DriverStatus;
}
const statusColors: Record<DriverStatus, string> = {
  active: 'green',
  inactive: 'orange',
  blocked: 'red',
};

const statusLabels: Record<DriverStatus, string> = {
  active: 'Активен',
  inactive: 'Неактивен',
  blocked: 'Заблокирован',
};

export const DriverList: React.FC<DriverListProps> = ({
  onCreate,
  onEdit,
  onView,
  searchQuery = '',
  searchStatus,
}) => {
  const [page, setPage] = useState(1);
  const [pageSize, setPageSize] = useState(10);

  const { data, isLoading, error } = useDrivers(page, pageSize);
  const deleteDriver = useDeleteDriver();
  const handleDelete = async (id: string) => {
    try {
      await deleteDriver.mutateAsync(id);
      message.success('Водитель успешно удален');
    } catch (error) {
      message.error('Ошибка при удалении водителя');
    }
  };

  // Фильтрация данных
  const filteredDrivers = (data?.drivers || [])
    .filter((driver) => {
      const matchesQuery =
        !searchQuery ||
        `${driver.lastName} ${driver.firstName} ${driver.middleName || ''}`
          .toLowerCase()
          .includes(searchQuery.toLowerCase()) ||
        driver.phone.includes(searchQuery) ||
        driver.driverLicenseNumber.toLowerCase().includes(searchQuery.toLowerCase());

      const matchesStatus = !searchStatus || driver.status === searchStatus;

      return matchesQuery && matchesStatus;
    })
    .sort((a, b) => a.lastName.localeCompare(b.lastName));
  const columns: ColumnsType<Driver> = [
    {
      title: 'Фото',
      dataIndex: 'photo',
      key: 'photo',
      width: 80,
      render: (photo) => (
        photo ? (
          <Image
            src={`http://localhost:8080${photo}`}
            alt="Фото водителя"
            width={50}
            height={50}
            style={{
              objectFit: 'cover',
              borderRadius: 4,
              cursor: 'pointer',
            }}
            preview={{
              mask: 'Увеличить',
            }}
          />
        ) : (
          <div
            style={{
              width: 50,
              height: 50,
              borderRadius: 4,
              backgroundColor: '#f0f0f0',
              display: 'flex',
              alignItems: 'center',
              justifyContent: 'center',
              fontSize: 12,
              color: '#999',
            }}
          >
            Нет фото
          </div>
        )
      ),    },    {
      title: 'ФИО',
      dataIndex: 'lastName',
      key: 'fullName',
      render: (_, record) => (
        <a onClick={() => onView?.(record)}>
          {`${record.lastName} ${record.firstName} ${record.middleName || ''}`}
        </a>
      ),
      sorter: (a, b) => a.lastName.localeCompare(b.lastName),
    },
    {
      title: 'Email',
      dataIndex: 'email',
      key: 'email',
      render: (email) => email || '-',
    },
    {
      title: '№ ВУ',
      dataIndex: 'driverLicenseNumber',
      key: 'driverLicenseNumber',
    },
    {
      title: 'Стаж',
      dataIndex: 'experienceYears',
      key: 'experienceYears',
      render: (years) => `${years} лет`,
    },
    {
      title: 'Статус',
      dataIndex: 'status',
      key: 'status',
      render: (status: DriverStatus) => (
        <Tag color={statusColors[status]}>{statusLabels[status]}</Tag>
      ),
      filters: [
        { text: 'Активен', value: 'active' },
        { text: 'Неактивен', value: 'inactive' },
        { text: 'Заблокирован', value: 'blocked' },
      ],
    },
    {
      title: 'Действия',
      key: 'actions',
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
            title="Вы уверены, что хотите удалить этого водителя?"
            onConfirm={() => handleDelete(record.id)}
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

  if (error) {
    return <div>Ошибка загрузки данных</div>;
  }

  return (
    <div>
      <div style={{ marginBottom: 16 }}>
        <Button type="primary" icon={<PlusOutlined />} onClick={onCreate}>
          Добавить водителя
        </Button>
      </div>
      <Table
        columns={columns}
        dataSource={filteredDrivers}
        rowKey="id"
        loading={isLoading}
        pagination={{
          current: page,
          pageSize: pageSize,
          total: filteredDrivers.length,
          showSizeChanger: true,
          showTotal: (total) => `Всего: ${total}`,
          onChange: (newPage, newPageSize) => {
            setPage(newPage);
            setPageSize(newPageSize || 10);
          },
        }}
      />    </div>
  );
};
