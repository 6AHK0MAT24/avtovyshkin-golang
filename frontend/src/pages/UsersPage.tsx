import React from 'react';
import { Typography, Empty } from 'antd';

const { Title } = Typography;

export const UsersPage: React.FC = () => {
  return (
    <div>
      <Title level={2}>Пользователи</Title>
      <Empty description="Страница пользователей находится в разработке" />
    </div>
  );
};
