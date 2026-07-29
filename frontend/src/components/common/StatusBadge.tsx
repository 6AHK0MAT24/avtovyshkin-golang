import React from 'react';
import { Tag } from 'antd';

export type StatusType = 'active' | 'inactive' | 'blocked';

export interface StatusBadgeProps {
  status: StatusType;
  customLabels?: Record<StatusType, string>;
  customColors?: Record<StatusType, string>;
}

const defaultLabels: Record<StatusType, string> = {
  active: 'Активен',
  inactive: 'Неактивен',
  blocked: 'Заблокирован',
};

const defaultColors: Record<StatusType, string> = {
  active: 'green',
  inactive: 'orange',
  blocked: 'red',
};

export const StatusBadge: React.FC<StatusBadgeProps> = ({
  status,
  customLabels = defaultLabels,
  customColors = defaultColors,
}) => {
  const label = customLabels[status] || defaultLabels[status];
  const color = customColors[status] || defaultColors[status];

  return <Tag color={color}>{label}</Tag>;
};

export default StatusBadge;