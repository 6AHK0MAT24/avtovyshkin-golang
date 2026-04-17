import React from 'react';
import { Layout, Badge, Space, Typography, Button } from 'antd';
import { MenuFoldOutlined, MenuUnfoldOutlined, WifiOutlined, DisconnectOutlined } from '@ant-design/icons';
import { useWebSocketStatus } from '../../hooks/useWebSocketStatus';

const { Header: AntHeader } = Layout;
const { Text, Title } = Typography;

interface HeaderProps {
  collapsed: boolean;
  onToggle: () => void;
}

export const Header: React.FC<HeaderProps> = ({ collapsed, onToggle }) => {
  const { isConnected } = useWebSocketStatus();

  return (
    <AntHeader style={{
      display: 'flex',
      alignItems: 'center',
      justifyContent: 'space-between',
      background: '#fff',
      padding: '0 24px',
      borderBottom: '1px solid #f0f0f0',
    }}>
      <Space size="middle">
        <Button
          type="text"
          icon={collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
          onClick={onToggle}
          style={{
            fontSize: '16px',
            width: 64,
            height: 64,
          }}
        />
        <Title level={4} style={{ margin: 0, color: '#1890ff' }}>
          Автовышкин
        </Title>
      </Space>

      <Space size="large">
        <Badge
          status={isConnected ? 'success' : 'error'}
          text={
            <Text>
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
