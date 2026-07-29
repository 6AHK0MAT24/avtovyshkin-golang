import React from 'react';
import { Descriptions, Space } from 'antd';

export interface DescriptionItem {
  label: string;
  value: React.ReactNode;
  icon?: React.ReactNode;
  span?: number;
}

export interface EntityDescriptionProps {
  items: DescriptionItem[];
  column?: number | { xxl?: number; xl?: number; lg?: number; md?: number; sm?: number; xs?: number };
  size?: 'small' | 'middle';
  bordered?: boolean;
  layout?: 'horizontal' | 'vertical';
}

export const EntityDescription: React.FC<EntityDescriptionProps> = ({
  items,
  column = 1,
  size = 'small',
  bordered = false,
  layout = 'horizontal',
}) => {
  const renderValue = (item: DescriptionItem) => {
    if (item.icon) {
      return (
        <Space>
          {item.icon}
          {item.value}
        </Space>
      );
    }
    return item.value;
  };

  return (
    <Descriptions
      column={column}
      size={size}
      bordered={bordered}
      layout={layout}
      items={items.map((item) => ({
        label: item.label,
        children: renderValue(item),
        span: item.span,
      }))}
    />
  );
};

export default EntityDescription;