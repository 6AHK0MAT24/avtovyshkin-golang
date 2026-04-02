import React from 'react';
import { Typography, Empty } from 'antd';

const { Title } = Typography;

export const RolesPage: React.FC = () => {
  return (
    <div>
      <Title level={2}>Роли</Title>
      <Empty description="Страница ролей находится в разработке" />
    </div>
  );
};
