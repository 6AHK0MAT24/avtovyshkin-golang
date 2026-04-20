import React from 'react';
import { Card, Descriptions, Tag, Image, Space, Button, Typography, Row, Col } from 'antd';
import {
  UserOutlined,
  PhoneOutlined,
  MailOutlined,
  IdcardOutlined,
  SafetyOutlined,
  EnvironmentOutlined,
  CalendarOutlined,
  EditOutlined,
  DeleteOutlined
} from '@ant-design/icons';
import type { Driver } from '../../types/driver';
import dayjs from 'dayjs';

const { Title, Text } = Typography;

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8082';
interface DriverCardProps {
  driver: Driver;
  onEdit?: (driver: Driver) => void;
  onDelete?: (id: string) => void;
  loading?: boolean;
}

const statusColors = {
  active: 'green',
  inactive: 'orange',
  blocked: 'red',
} as const;

const statusLabels = {
  active: 'Активен',
  inactive: 'Неактивен',
  blocked: 'Заблокирован',
};

export const DriverCard: React.FC<DriverCardProps> = ({ 
  driver, 
  onEdit, 
  onDelete,
  loading = false 
}) => {
  const fullName = [driver.lastName, driver.firstName, driver.middleName]
    .filter(Boolean)
    .join(' ');

  const calculateAge = (birthDate?: string): number | null => {
    if (!birthDate) return null;
    return dayjs().diff(dayjs(birthDate), 'year');
  };

  const age = calculateAge(driver.birthDate);
  const licenseExpiryDays = dayjs(driver.driverLicenseExpiryDate).diff(dayjs(), 'day');
  const isLicenseExpiringSoon = licenseExpiryDays > 0 && licenseExpiryDays <= 30;
  const isLicenseExpired = licenseExpiryDays <= 0;

  return (
    <Card
      loading={loading}
      title={
        <Space>
          <UserOutlined />
          <Title level={4} style={{ margin: 0 }}>
            {fullName}
          </Title>
          <Tag color={statusColors[driver.status]}>
            {statusLabels[driver.status]}
          </Tag>
        </Space>
      }
      extra={
        <Space>
          {onEdit && (
            <Button 
              type="primary" 
              icon={<EditOutlined />} 
              onClick={() => onEdit(driver)}
            >
              Редактировать
            </Button>
          )}
          {onDelete && (
            <Button 
              danger 
              icon={<DeleteOutlined />} 
              onClick={() => onDelete(driver.id)}
            >
              Удалить
            </Button>
          )}
        </Space>
      }
    >
      <Row gutter={[24, 24]}>
        {/* Основная информация */}
        <Col xs={24} lg={12}>
          <Card 
            size="small" 
            title={<Space><UserOutlined />Основная информация</Space>}
          >
            <Descriptions column={1} size="small">
              <Descriptions.Item label="ФИО">{fullName}</Descriptions.Item>
              <Descriptions.Item label="Телефон">
                <Space>
                  <PhoneOutlined />
                  <a href={`tel:${driver.phone}`}>{driver.phone}</a>
                </Space>
              </Descriptions.Item>
              {driver.email && (
                <Descriptions.Item label="Email">
                  <Space>
                    <MailOutlined />
                    <a href={`mailto:${driver.email}`}>{driver.email}</a>
                  </Space>
                </Descriptions.Item>
              )}
              {driver.birthDate && (
                <Descriptions.Item label="Дата рождения">
                  <Space>
                    <CalendarOutlined />
                    {dayjs(driver.birthDate).format('DD.MM.YYYY')}
                    {age !== null && <Text type="secondary">({age} лет)</Text>}
                  </Space>
                </Descriptions.Item>
              )}
              {driver.address && (
                <Descriptions.Item label="Адрес">
                  <Space>
                    <EnvironmentOutlined />
                    {driver.address}
                  </Space>
                </Descriptions.Item>
              )}
              <Descriptions.Item label="Стаж">
                {driver.experienceYears} {driver.experienceYears === 1 ? 'год' : 'лет'}
              </Descriptions.Item>
            </Descriptions>
          </Card>
        </Col>

        {/* Водительские права */}
        <Col xs={24} lg={12}>
          <Card 
            size="small" 
            title={
              <Space>
                <SafetyOutlined />
                Водительские права
                {isLicenseExpired && <Tag color="red">Истекли</Tag>}
                {isLicenseExpiringSoon && <Tag color="orange">Скоро истекут</Tag>}
              </Space>
            }
          >
            <Descriptions column={1} size="small">
              <Descriptions.Item label="Номер">{driver.driverLicenseNumber}</Descriptions.Item>
              <Descriptions.Item label="Дата выдачи">
                {dayjs(driver.driverLicenseIssueDate).format('DD.MM.YYYY')}
              </Descriptions.Item>
              <Descriptions.Item label="Срок действия">
                <Space>
                  {dayjs(driver.driverLicenseExpiryDate).format('DD.MM.YYYY')}
                  {licenseExpiryDays > 0 && (
                    <Text type="secondary">(осталось {licenseExpiryDays} дней)</Text>
                  )}
                </Space>
              </Descriptions.Item>
            </Descriptions>

            {driver.driverLicensePhoto && (
              <div style={{ marginTop: 16 }}>
                <Text strong>Фото прав:</Text>
                <div style={{ marginTop: 8 }}>
                  <Image
                    width={200}
                    src={`${API_BASE_URL}${driver.driverLicensePhoto}`}
                    alt="Фото водительских прав"
                    style={{ borderRadius: 8 }}
                  />
                </div>
              </div>
            )}

            {driver.driverLicenseScan && (
              <div style={{ marginTop: 16 }}>
                <Text strong>Скан прав:</Text>
                <div style={{ marginTop: 8 }}>
                  <Image
                    width={200}
                    src={`${API_BASE_URL}${driver.driverLicenseScan}`}
                    alt="Скан водительских прав"
                    style={{ borderRadius: 8 }}
                  />
                </div>
              </div>
            )}          </Card>
        </Col>

        {/* Паспорт */}
        {(driver.passportSeries || driver.passportNumber) && (
          <Col xs={24} lg={12}>
            <Card 
              size="small" 
              title={<Space><IdcardOutlined />Паспорт</Space>}
            >
              <Descriptions column={1} size="small">
                {driver.passportSeries && (
                  <Descriptions.Item label="Серия">{driver.passportSeries}</Descriptions.Item>
                )}
                {driver.passportNumber && (
                  <Descriptions.Item label="Номер">{driver.passportNumber}</Descriptions.Item>
                )}
                {driver.passportIssueDate && (
                  <Descriptions.Item label="Дата выдачи">
                    {dayjs(driver.passportIssueDate).format('DD.MM.YYYY')}
                  </Descriptions.Item>
                )}
              </Descriptions>

              {driver.passportPhoto && (
                <div style={{ marginTop: 16 }}>
                  <Text strong>Фото паспорта:</Text>
                  <div style={{ marginTop: 8 }}>
                    <Image
                      width={200}
                      src={`${API_BASE_URL}${driver.passportPhoto}`}
                      alt="Фото паспорта"
                      style={{ borderRadius: 8 }}
                    />
                  </div>
                </div>
              )}

              {driver.passportScan && (
                <div style={{ marginTop: 16 }}>
                  <Text strong>Скан паспорта:</Text>
                  <div style={{ marginTop: 8 }}>
                    <Image
                      width={200}
                      src={`${API_BASE_URL}${driver.passportScan}`}
                      alt="Скан паспорта"
                      style={{ borderRadius: 8 }}
                    />
                  </div>
                </div>
              )}            </Card>
          </Col>
        )}

        {/* Метаданные */}
        <Col xs={24} lg={12}>
          <Card 
            size="small" 
            title={<Space><CalendarOutlined />Информация о записи</Space>}
          >
            <Descriptions column={1} size="small">
              <Descriptions.Item label="Создан">
                {dayjs(driver.createdAt).format('DD.MM.YYYY HH:mm')}
              </Descriptions.Item>
              <Descriptions.Item label="Обновлен">
                {dayjs(driver.updatedAt).format('DD.MM.YYYY HH:mm')}
              </Descriptions.Item>
              <Descriptions.Item label="ID">
                <Text copyable>{driver.id}</Text>
              </Descriptions.Item>
            </Descriptions>
          </Card>
        </Col>
      </Row>
    </Card>
  );
};
export default DriverCard;
