import type { ReactNode } from 'react';
import { useState } from 'react';
import { Layout as AntLayout, Menu } from 'antd';
import { CarOutlined, UserOutlined, TeamOutlined, TruckOutlined, ContactsOutlined } from '@ant-design/icons';
import { Header } from './Header';
import { useNavigate, useLocation } from 'react-router-dom';
const { Sider, Content } = AntLayout;

interface LayoutProps {
  children: ReactNode;
}

export const Layout: React.FC<LayoutProps> = ({ children }) => {
  const [collapsed, setCollapsed] = useState(false);
  const navigate = useNavigate();
  const location = useLocation();

  const menuItems = [
    {
      key: '/',
      icon: <TruckOutlined />,
      label: 'Водители',
    },
    {
      key: '/roles',
      icon: <TeamOutlined />,
      label: 'Роли',
    },
    {
      key: '/users',
      icon: <UserOutlined />,
      label: 'Пользователи',
    },
    {
      key: '/vehicles',
      icon: <CarOutlined />,
      label: 'Автовышки',
    },
    {
      key: '/clients',
      icon: <ContactsOutlined />,
      label: 'Клиенты',
    },
  ];
  const handleMenuClick = ({ key }: { key: string }) => {
    navigate(key);
  };

  return (
    <AntLayout style={{ minHeight: '100vh' }}>
      <Sider
        trigger={null}
        collapsible
        collapsed={collapsed}
        style={{
          overflow: 'auto',
          height: '100vh',
          position: 'fixed',
          left: 0,
          top: 0,
          bottom: 0,
        }}
      >
        <div
          style={{
            height: 64,
            display: 'flex',
            alignItems: 'center',
            justifyContent: collapsed ? 'center' : 'flex-start',
            padding: collapsed ? 0 : '0 24px',
            background: '#001529',
            borderBottom: '1px solid #303030',
          }}
        >
          <CarOutlined style={{ fontSize: 24, color: '#1890ff' }} />
          {!collapsed && (
            <span style={{ marginLeft: 12, fontSize: 18, fontWeight: 'bold', color: 'white' }}>
              Автовышкин
            </span>
          )}
        </div>
        <Menu
          theme="dark"
          mode="inline"
          selectedKeys={[location.pathname]}
          items={menuItems}
          onClick={handleMenuClick}
        />
      </Sider>
      <AntLayout style={{ marginLeft: collapsed ? 80 : 200, transition: 'margin-left 0.2s' }}>
        <Header collapsed={collapsed} onToggle={() => setCollapsed(!collapsed)} />
        <Content style={{ padding: '24px', background: '#f0f2f5', minHeight: 'calc(100vh - 64px)' }}>
          <div style={{ background: 'white', padding: '24px', borderRadius: '8px', minHeight: 'calc(100vh - 112px)' }}>
            {children}
          </div>
        </Content>
      </AntLayout>
    </AntLayout>
  );
};
