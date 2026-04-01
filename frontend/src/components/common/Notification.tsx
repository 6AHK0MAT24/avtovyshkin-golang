import React, { useEffect } from 'react';
import { notification, Button } from 'antd';
import { CheckCircleOutlined, CloseCircleOutlined, InfoCircleOutlined, WarningOutlined } from '@ant-design/icons';
import type { NotificationPlacement } from 'antd/es/notification/interface';export type NotificationType = 'success' | 'error' | 'info' | 'warning';

export interface NotificationConfig {
  type: NotificationType;
  message: string;
  description?: string;
  duration?: number;
  placement?: NotificationPlacement;
  showProgress?: boolean;
}

const notificationIcons = {
  success: <CheckCircleOutlined style={{ color: '#52c41a' }} />,
  error: <CloseCircleOutlined style={{ color: '#ff4d4f' }} />,
  info: <InfoCircleOutlined style={{ color: '#1890ff' }} />,
  warning: <WarningOutlined style={{ color: '#faad14' }} />,
};

export const showNotification = ({
  type,
  message,
  description,
  duration = 4.5,
  placement = 'topRight',
  showProgress = false,
}: NotificationConfig) => {
  notification[type]({
    icon: notificationIcons[type],
    message,
    description,
    duration,
    placement,
    showProgress,
    style: {
      borderRadius: '8px',
    },
  });
};

export const showSuccessNotification = (message: string, description?: string) => {
  showNotification({ type: 'success', message, description });
};

export const showErrorNotification = (message: string, description?: string) => {
  showNotification({ type: 'error', message, description, duration: 6 });
};

export const showInfoNotification = (message: string, description?: string) => {
  showNotification({ type: 'info', message, description });
};

export const showWarningNotification = (message: string, description?: string) => {
  showNotification({ type: 'warning', message, description });
};

// Компонент для отображения уведомлений о действиях с водителями
export const DriverNotification: React.FC = () => {
  useEffect(() => {
    // Здесь можно добавить логику для подписки на события WebSocket
    // и автоматического отображения уведомлений
    const handleDriverCreated = (event: CustomEvent) => {
      showSuccessNotification(
        'Водитель создан',
        `Водитель ${event.detail.name} успешно добавлен в систему`
      );
    };

    const handleDriverUpdated = (event: CustomEvent) => {
      showInfoNotification(
        'Водитель обновлен',
        `Данные водителя ${event.detail.name} успешно обновлены`
      );
    };

    const handleDriverDeleted = (event: CustomEvent) => {
      showWarningNotification(
        'Водитель удален',
        `Водитель ${event.detail.name} был удален из системы`
      );
    };

    window.addEventListener('driver.created', handleDriverCreated as EventListener);
    window.addEventListener('driver.updated', handleDriverUpdated as EventListener);
    window.addEventListener('driver.deleted', handleDriverDeleted as EventListener);

    return () => {
      window.removeEventListener('driver.created', handleDriverCreated as EventListener);
      window.removeEventListener('driver.updated', handleDriverUpdated as EventListener);
      window.removeEventListener('driver.deleted', handleDriverDeleted as EventListener);
    };
  }, []);

  return null;
};

// Компонент для подтверждения действий
export const ConfirmNotification: React.FC<{
  onConfirm: () => void;
  onCancel: () => void;
  title: string;
  description: string;
}> = ({ onConfirm, onCancel, title, description }) => {
  useEffect(() => {
    notification.open({
      message: title,
      description,
      duration: 0,
      placement: 'topRight',
      icon: <WarningOutlined style={{ color: '#faad14' }} />,
      btn: (
        <div style={{ display: 'flex', gap: 8 }}>
          <Button size="small" onClick={onCancel}>
            Отмена
          </Button>
          <Button type="primary" size="small" danger onClick={onConfirm}>
            Подтвердить
          </Button>
        </div>
      ),
      onClose: onCancel,
    });
  }, [title, description, onConfirm, onCancel]);
  return null;
};

// Хук для работы с уведомлениями
export const useNotification = () => {
  return {
    success: showSuccessNotification,
    error: showErrorNotification,
    info: showInfoNotification,
    warning: showWarningNotification,
  };
};
