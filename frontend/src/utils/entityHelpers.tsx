import React from 'react';
import { Typography } from 'antd';
import {
  PhoneOutlined,
  MailOutlined,
  EnvironmentOutlined,
  CalendarOutlined,
  IdcardOutlined,
} from '@ant-design/icons';
import dayjs from 'dayjs';
import type { DescriptionItem } from '../components/common/EntityDescription';

const { Text } = Typography;

// Helper functions for common description items
export const createPhoneItem = (phone: string): DescriptionItem => ({
  label: 'Телефон',
  value: <a href={`tel:${phone}`}>{phone}</a>,
  icon: <PhoneOutlined />,
});

export const createEmailItem = (email: string): DescriptionItem => ({
  label: 'Email',
  value: <a href={`mailto:${email}`}>{email}</a>,
  icon: <MailOutlined />,
});

export const createDateItem = (label: string, date?: string, format: string = 'DD.MM.YYYY'): DescriptionItem => ({
  label,
  value: date ? dayjs(date).format(format) : '-',
  icon: <CalendarOutlined />,
});

export const createAddressItem = (address?: string): DescriptionItem => ({
  label: 'Адрес',
  value: address || '-',
  icon: <EnvironmentOutlined />,
});

export const createIdItem = (id: string, label: string = 'ID'): DescriptionItem => ({
  label,
  value: <Text copyable>{id}</Text>,
  icon: <IdcardOutlined />,
});

export const formatDate = (date?: string, format: string = 'DD.MM.YYYY'): string => {
  return date ? dayjs(date).format(format) : '-';
};

export const formatDateTime = (date?: string, format: string = 'DD.MM.YYYY HH:mm'): string => {
  return date ? dayjs(date).format(format) : '-';
};

export const calculateAge = (birthDate?: string): number | null => {
  if (!birthDate) return null;
  return dayjs().diff(dayjs(birthDate), 'year');
};

export const getDaysUntil = (date: string): number => {
  return dayjs(date).diff(dayjs(), 'day');
};

export const isExpiringSoon = (date: string, daysThreshold: number = 30): boolean => {
  const days = getDaysUntil(date);
  return days > 0 && days <= daysThreshold;
};

export const isExpired = (date: string): boolean => {
  return getDaysUntil(date) <= 0;
};