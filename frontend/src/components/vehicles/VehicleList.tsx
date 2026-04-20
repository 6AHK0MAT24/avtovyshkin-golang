import React, { useState } from 'react';
import { Table, Button, Space, Tag, Popconfirm, message, Image } from 'antd';
import { EditOutlined, DeleteOutlined } from '@ant-design/icons';
import { useVehicles, useDeleteVehicle } from '../../hooks/useVehicles';
import type { Vehicle, VehicleStatus } from '../../types/vehicle';
import type { ColumnsType } from 'antd/es/table';
import { ImageSlider } from './ImageSlider';
interface VehicleListProps {
  onEdit: (vehicle: Vehicle) => void;
  onView?: (vehicle: Vehicle) => void;
  searchQuery?: string;
  searchStatus?: VehicleStatus;
}

const statusColors: Record<VehicleStatus, string> = {
  active: 'green',
  inactive: 'orange',
  blocked: 'red',
};

const statusLabels: Record<VehicleStatus, string> = {
  active: 'Активен',
  inactive: 'Неактивен',
  blocked: 'Заблокирован',
};

export const VehicleList: React.FC<VehicleListProps> = ({
  onEdit,
  onView,
  searchQuery = '',
  searchStatus,
}) => {
  const [page, setPage] = useState(1);
  const [pageSize, setPageSize] = useState(10);
  const [sliderVisible, setSliderVisible] = useState(false);
  const [sliderImages, setSliderImages] = useState<string[]>([]);
  const [sliderInitialIndex, setSliderInitialIndex] = useState(0);

  const { data, isLoading, error } = useVehicles(page, pageSize);
  const deleteVehicle = useDeleteVehicle();
  const handleDelete = async (id: string) => {
    try {
      await deleteVehicle.mutateAsync(id);
      message.success('Автовышка успешно удалена');
    } catch (error) {
      message.error('Ошибка при удалении автовышки');
    }
  };

  const handleImageClick = (imgArray: string[], mainImageIndex: number) => {
    if (!imgArray || imgArray.length === 0) return;
    
    // Конвертируем из 1-based (пользователь) в 0-based (ImageSlider)
    const initialIndex = (mainImageIndex || 1) - 1;
    setSliderImages(imgArray);
    setSliderInitialIndex(initialIndex);
    setSliderVisible(true);
  };
  // Фильтрация данных
  const filteredVehicles = (data?.vehicles || [])
    .filter((vehicle) => {
      const matchesQuery =
        !searchQuery ||
        vehicle.garageNumber.toLowerCase().includes(searchQuery.toLowerCase()) ||
        vehicle.vin.toLowerCase().includes(searchQuery.toLowerCase()) ||
        (vehicle.brand && vehicle.brand.toLowerCase().includes(searchQuery.toLowerCase()));

      const matchesStatus = !searchStatus || vehicle.status === searchStatus;

      return matchesQuery && matchesStatus;
    })
    .sort((a, b) => a.garageNumber.localeCompare(b.garageNumber));

  const columns: ColumnsType<Vehicle> = [
    {
      title: 'Фото',
      dataIndex: 'imgArray',
      key: 'photo',
      width: 80,
      render: (imgArray, record) => {
        // Конвертируем из 1-based (пользователь) в 0-based (массив)
        const mainImageIndex = (record.mainImageIndex || 1) - 1;
        const mainImage = imgArray && imgArray.length > 0
          ? imgArray[mainImageIndex]
          : null;

        return mainImage ? (
          <Image
            src={`http://localhost:8081${mainImage}`}
            alt="Фото автовышки"
            width={50}
            height={50}
            style={{
              objectFit: 'cover',
              borderRadius: 4,
              cursor: 'pointer',
            }}
            preview={false}
            onClick={() => handleImageClick(imgArray, record.mainImageIndex || 1)}
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
        );
      },
    },    {
      title: 'Гаражный №',
      dataIndex: 'garageNumber',
      key: 'garageNumber',
      render: (garageNumber, record) => (
        <a onClick={() => onView?.(record)}>
          {garageNumber}
        </a>
      ),
      sorter: (a, b) => a.garageNumber.localeCompare(b.garageNumber),
    },
    {
      title: 'Бренд',
      dataIndex: 'brand',
      key: 'brand',
      render: (brand) => brand || '-',
    },
    {
      title: 'Высота',
      dataIndex: 'height',
      key: 'height',
      render: (height) => `${height} м`,
      sorter: (a, b) => a.height - b.height,
    },
    {
      title: 'Тип',
      dataIndex: 'type',
      key: 'type',
      render: (type) => type || '-',
    },
    {
      title: 'VIN',
      dataIndex: 'vin',
      key: 'vin',
      render: (vin) => vin || '-',
    },
    {
      title: 'Статус',
      dataIndex: 'status',
      key: 'status',
      render: (status: VehicleStatus) => (
        <Tag color={statusColors[status]}>{statusLabels[status]}</Tag>
      ),
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
            title="Вы уверены, что хотите удалить эту автовышку?"
            onConfirm={() => handleDelete(record.id)}
            okText="Ок"
            cancelText="Отмена"
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
      <Table
        columns={columns}
        dataSource={filteredVehicles}
        rowKey="id"
        loading={isLoading}
        pagination={{
          current: page,
          pageSize: pageSize,
          total: filteredVehicles.length,
          showSizeChanger: true,
          showTotal: (total) => `Всего: ${total}`,
          onChange: (newPage, newPageSize) => {
            setPage(newPage);
            setPageSize(newPageSize || 10);
          },
        }}
      />
      
      <ImageSlider
        images={sliderImages}
        initialIndex={sliderInitialIndex}
        visible={sliderVisible}
        onClose={() => setSliderVisible(false)}
      />
    </div>
  );
};
