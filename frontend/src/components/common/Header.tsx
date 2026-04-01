import React, { useState } from 'react';
import { Layout, Menu, Badge, Space, Typography } from 'antd';
import { CarOutlined, UserOutlined, SettingOutlined, WifiOutlined, DisconnectOutlined } from '@ant-design/icons';
import { useWebSocketStatus } from '../../hooks/useWebSocketStatus';
const { Header: AntHeader } = Layout;
const { Text } = Typography;

export const Header: React.FC = () => {
  const [current, setCurrent] = useState('drivers');
  const { isConnected } = useWebSocketStatus();
  const menuItems = [
    {
      key: 'drivers',
      icon: <CarOutlined />,
      label: 'Водители',
    },
    {
      key: 'profile',
      icon: <UserOutlined />,
      label: 'Профиль',
    },
    {
      key: 'settings',
      icon: <SettingOutlined />,
      label: 'Настройки',
    },
  ];

  return (
    <AntHeader style={{ 
      display: 'flex', 
      alignItems: 'center', 
      justifyContent: 'space-between',
      background: '#001529',
      padding: '0 24px'
    }}>
      <Space size="large">
        <div style={{ 
          color: 'white', 
          fontSize: 20, 
          fontWeight: 'bold', 
          display: 'flex', 
          alignItems: 'center', 
          gap: 10 
        }}>
          <CarOutlined style={{ fontSize: 24 }} />
          Автовышкин
        </div>
        <Menu
          theme="dark"
          mode="horizontal"
          selectedKeys={[current]}
          items={menuItems}
          onClick={({ key }) => setCurrent(key)}
          style={{ 
            background: 'transparent',
            minWidth: 300
          }}
        />
      </Space>

      <Space size="large">
        <Badge 
          status={isConnected ? 'success' : 'error'} 
          text={
            <Text style={{ color: 'white' }}>
              {isConnected ? (
                <>
                  <WifiOutlined /> Online
                </>
              ) : (
                <>
                  <DisconnectOutlined /> Offline
                </>
              )}
            </Text>
          }
        />
      </Space>
    </AntHeader>
  );
};
