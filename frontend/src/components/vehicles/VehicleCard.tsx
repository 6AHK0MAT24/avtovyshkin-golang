import React, { useState } from 'react';
import { Card, Descriptions, Tag, Image, Space, Button, Typography, Row, Col, Popconfirm } from 'antd';import {
  CarOutlined,
  EditOutlined,
  DeleteOutlined,
  EnvironmentOutlined,
  DollarOutlined,
  InfoCircleOutlined,
  ToolOutlined,
} from '@ant-design/icons';
import type { Vehicle, VehicleStatus } from '../../types/vehicle';
import { ImageSlider } from './ImageSlider';

const { Title, Text } = Typography;

interface VehicleCardProps {
  vehicle: Vehicle;
  onEdit?: (vehicle: Vehicle) => void;
  onDelete?: (id: string) => void;
  loading?: boolean;
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

export const VehicleCard: React.FC<VehicleCardProps> = ({
  vehicle,
  onEdit,
  onDelete,
  loading = false,
}) => {
  const [sliderVisible, setSliderVisible] = useState(false);
  const [sliderInitialIndex, setSliderInitialIndex] = useState(0);

  const handleImageClick = () => {
    // Конвертируем из 1-based (пользователь) в 0-based (ImageSlider)
    const userMainIndex = vehicle.mainImageIndex || 1;
    setSliderInitialIndex(userMainIndex - 1);
    setSliderVisible(true);
  };

  const handleDelete = async () => {
    if (onDelete) {
      try {
        await onDelete(vehicle.id);
      } catch (error) {
        console.error('Ошибка при удалении автовышки:', error);
      }
    }
  };
  const mainImage =
    vehicle.imgArray && vehicle.imgArray.length > 0
      ? vehicle.imgArray[(vehicle.mainImageIndex || 1) - 1] // Конвертируем из 1-based в 0-based
      : null;
  return (
    <>
      <Card
        loading={loading}
        title={
          <Space>
            <CarOutlined />
            <Title level={4} style={{ margin: 0 }}>
              Автовышка {vehicle.garageNumber}
            </Title>
            <Tag color={statusColors[vehicle.status]}>
              {statusLabels[vehicle.status]}
            </Tag>
          </Space>
        }
        extra={
          <Space>
            {onEdit && (
              <Button type="primary" icon={<EditOutlined />} onClick={() => onEdit(vehicle)}>
                Редактировать
              </Button>
            )}
            {onDelete && (
              <Popconfirm
                title="Вы уверены, что хотите удалить эту автовышку?"
                onConfirm={handleDelete}
                okText="Ок"
                cancelText="Отмена"
              >
                <Button danger icon={<DeleteOutlined />}>
                  Удалить
                </Button>
              </Popconfirm>
            )}
          </Space>
        }
      >
        <Row gutter={[24, 24]}>
          {/* Изображение */}
          <Col xs={24} lg={8}>
            <Card size="small" title={<Space><CarOutlined />Фото</Space>}>
              {mainImage ? (
                <div>
                  <Image
                    src={`http://localhost:8081${mainImage}`}
                    alt="Основное фото автовышки"
                    style={{
                      width: '100%',
                      borderRadius: 8,
                      cursor: 'pointer',
                    }}
                    onClick={() => handleImageClick()}
                    preview={false}
                  />                  {vehicle.imgArray && vehicle.imgArray.length > 1 && (
                    <Text type="secondary" style={{ display: 'block', marginTop: 8, textAlign: 'center' }}>
                      Нажмите для просмотра всех {vehicle.imgArray.length} изображений
                    </Text>
                  )}
                </div>
              ) : (
                <div
                  style={{
                    width: '100%',
                    height: 200,
                    borderRadius: 8,
                    backgroundColor: '#f0f0f0',
                    display: 'flex',
                    alignItems: 'center',
                    justifyContent: 'center',
                    color: '#999',
                    fontSize: 14,
                  }}
                >
                  Нет фото
                </div>
              )}
            </Card>
          </Col>

          {/* Основная информация */}
          <Col xs={24} lg={16}>
            <Card size="small" title={<Space><InfoCircleOutlined />Основная информация</Space>}>
              <Descriptions column={2} size="small">
                <Descriptions.Item label="Гаражный номер">{vehicle.garageNumber}</Descriptions.Item>
                <Descriptions.Item label="VIN">{vehicle.vin}</Descriptions.Item>
                <Descriptions.Item label="Бренд">{vehicle.brand || '-'}</Descriptions.Item>
                <Descriptions.Item label="Тип">{vehicle.type || '-'}</Descriptions.Item>
                <Descriptions.Item label="Высота">{vehicle.height} м</Descriptions.Item>
                <Descriptions.Item label="Статус">
                  <Tag color={statusColors[vehicle.status]}>{statusLabels[vehicle.status]}</Tag>
                </Descriptions.Item>
                {vehicle.machine && (
                  <Descriptions.Item label="Машина">{vehicle.machine}</Descriptions.Item>
                )}
                {vehicle.rostechReg && (
                  <Descriptions.Item label="РостехРег">{vehicle.rostechReg}</Descriptions.Item>
                )}
              </Descriptions>
            </Card>
          </Col>

          {/* Технические характеристики */}
          <Col xs={24} lg={12}>
            <Card size="small" title={<Space><ToolOutlined />Технические характеристики</Space>}>
              <Descriptions column={1} size="small">
                {vehicle.power !== undefined && (
                  <Descriptions.Item label="Мощность">{vehicle.power} л.с.</Descriptions.Item>
                )}
                {vehicle.length !== undefined && (
                  <Descriptions.Item label="Длина">{vehicle.length} м</Descriptions.Item>
                )}
                {vehicle.width !== undefined && (
                  <Descriptions.Item label="Ширина">{vehicle.width} м</Descriptions.Item>
                )}
                {vehicle.heightTs !== undefined && (
                  <Descriptions.Item label="Высота ТС">{vehicle.heightTs} м</Descriptions.Item>
                )}
                {vehicle.widthWithSupports !== undefined && (
                  <Descriptions.Item label="Ширина с опорами">{vehicle.widthWithSupports} м</Descriptions.Item>
                )}
                {vehicle.mass !== undefined && (
                  <Descriptions.Item label="Масса">{vehicle.mass} кг</Descriptions.Item>
                )}
              </Descriptions>
            </Card>
          </Col>

          {/* Характеристики люльки */}
          {(vehicle.cradleWidthFolded !== undefined ||
            vehicle.cradleWidthExtended !== undefined ||
            vehicle.cradleLengthFolded !== undefined ||
            vehicle.cradleLengthExtended !== undefined) && (
            <Col xs={24} lg={12}>
              <Card size="small" title={<Space><ToolOutlined />Характеристики люльки</Space>}>
                <Descriptions column={1} size="small">
                  {vehicle.cradleWidthFolded !== undefined && (
                    <Descriptions.Item label="Ширина (сложена)">
                      {vehicle.cradleWidthFolded} м
                    </Descriptions.Item>
                  )}
                  {vehicle.cradleWidthExtended !== undefined && (
                    <Descriptions.Item label="Ширина (разложена)">
                      {vehicle.cradleWidthExtended} м
                    </Descriptions.Item>
                  )}
                  {vehicle.cradleLengthFolded !== undefined && (
                    <Descriptions.Item label="Длина (сложена)">
                      {vehicle.cradleLengthFolded} м
                    </Descriptions.Item>
                  )}
                  {vehicle.cradleLengthExtended !== undefined && (
                    <Descriptions.Item label="Длина (разложена)">
                      {vehicle.cradleLengthExtended} м
                    </Descriptions.Item>
                  )}
                </Descriptions>
              </Card>
            </Col>
          )}

          {/* Цены */}
          {(vehicle.price5 !== undefined || vehicle.price22 !== undefined) && (
            <Col xs={24} lg={12}>
              <Card size="small" title={<Space><DollarOutlined />Цены</Space>}>
                <Descriptions column={1} size="small">
                  {vehicle.price5 !== undefined && (
                    <Descriptions.Item label="Цена (5м)">{vehicle.price5} ₽</Descriptions.Item>
                  )}
                  {vehicle.price22 !== undefined && (
                    <Descriptions.Item label="Цена (22м)">{vehicle.price22} ₽</Descriptions.Item>
                  )}
                </Descriptions>
              </Card>
            </Col>
          )}

          {/* Описание и особые отметки */}
          {(vehicle.description || vehicle.special) && (
            <Col xs={24} lg={12}>
              <Card size="small" title={<Space><EnvironmentOutlined />Дополнительно</Space>}>
                {vehicle.description && (
                  <div style={{ marginBottom: 16 }}>
                    <Text strong>Описание:</Text>
                    <div style={{ marginTop: 4 }}>
                      <Text>{vehicle.description}</Text>
                    </div>
                  </div>
                )}
                {vehicle.special && (
                  <div>
                    <Text strong>Особые отметки:</Text>
                    <div style={{ marginTop: 4 }}>
                      <Text>{vehicle.special}</Text>
                    </div>
                  </div>
                )}
              </Card>
            </Col>
          )}

          {/* Галерея всех изображений */}
          {vehicle.imgArray && vehicle.imgArray.length > 1 && (
            <Col xs={24}>
              <Card size="small" title={<Space><CarOutlined />Все фото ({vehicle.imgArray.length})</Space>}>
                <div
                  style={{
                    display: 'grid',
                    gridTemplateColumns: 'repeat(auto-fill, minmax(120px, 1fr))',
                    gap: '12px',
                  }}
                >
                  {vehicle.imgArray.map((img, index) => (
                    <div
                      key={index}
                      style={{
                        position: 'relative',
                        aspectRatio: '1',
                        borderRadius: '8px',
                        overflow: 'hidden',
                        border: (vehicle.mainImageIndex || 1) === index + 1 ? '3px solid #1890ff' : '2px solid #d9d9d9',
                        cursor: 'pointer',
                      }}
                      onClick={() => {
                        setSliderInitialIndex(index);
                        setSliderVisible(true);
                      }}
                    >
                      <img
                        src={`http://localhost:8081${img}`}
                        alt={`Фото ${index + 1}`}
                        style={{
                          width: '100%',
                          height: '100%',
                          objectFit: 'cover',
                        }}
                      />
                      {(vehicle.mainImageIndex || 1) === index + 1 && (
                        <div
                          style={{
                            position: 'absolute',
                            bottom: 0,
                            left: 0,
                            right: 0,
                            backgroundColor: 'rgba(24, 144, 255, 0.9)',
                            color: 'white',
                            textAlign: 'center',
                            padding: '4px 8px',
                            fontSize: '12px',
                            fontWeight: 'bold',
                          }}
                        >
                          Основное
                        </div>
                      )}
                    </div>
                  ))}
                </div>
              </Card>
            </Col>
          )}
        </Row>
      </Card>
      {/* Галерея изображений */}
      <ImageSlider
        images={vehicle.imgArray || []}
        initialIndex={sliderInitialIndex}
        visible={sliderVisible}
        onClose={() => setSliderVisible(false)}
      />
    </>
  );
};
