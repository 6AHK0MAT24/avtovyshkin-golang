import type { ReactNode } from 'react';
import { Layout as AntLayout } from 'antd';
import { Header } from './Header';

const { Content } = AntLayout;

interface LayoutProps {
  children: ReactNode;
}

export const Layout: React.FC<LayoutProps> = ({ children }) => {
  return (
    <AntLayout style={{ minHeight: '100vh' }}>
      <Header />
      <Content style={{ padding: '24px', background: '#f0f2f5' }}>
        <div style={{ background: 'white', padding: '24px', borderRadius: '8px', minHeight: 'calc(100vh - 112px)' }}>
          {children}
        </div>
      </Content>
    </AntLayout>
  );
};
