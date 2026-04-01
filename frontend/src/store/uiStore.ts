import { create } from 'zustand';

export type NotificationType = 'success' | 'error' | 'warning' | 'info';

export interface Notification {
  id: string;
  type: NotificationType;
  message: string;
  description?: string;
  duration?: number;
}

export type ModalType = 'createDriver' | 'editDriver' | 'viewDriver' | 'deleteDriver' | null;

export interface ModalState {
  type: ModalType;
  data?: any;
  isOpen: boolean;
}

export type Theme = 'light' | 'dark' | 'auto';

interface UIState {
  // Состояние уведомлений
  notifications: Notification[];

  // Состояние модальных окон
  modal: ModalState;

  // Тема
  theme: Theme;

  // Состояние загрузки (глобальное)
  globalLoading: boolean;

  // Actions для уведомлений
  showNotification: (notification: Omit<Notification, 'id'>) => void;
  removeNotification: (id: string) => void;
  clearNotifications: () => void;

  // Actions для модальных окон
  openModal: (type: ModalType, data?: any) => void;
  closeModal: () => void;

  // Actions для темы
  setTheme: (theme: Theme) => void;

  // Actions для глобальной загрузки
  setGlobalLoading: (loading: boolean) => void;
}

export const useUIStore = create<UIState>((set) => ({
  // Начальное состояние
  notifications: [],
  modal: {
    type: null,
    isOpen: false,
  },
  theme: 'light',
  globalLoading: false,

  // Actions для уведомлений
  showNotification: (notification) =>
    set((state) => {
      const id = Date.now().toString() + Math.random().toString(36).substr(2, 9);
      const newNotification: Notification = {
        id,
        ...notification,
      };

      // Автоматическое удаление уведомления через duration (по умолчанию 4.5 секунды)
      const duration = notification.duration ?? 4500;
      if (duration > 0) {
        setTimeout(() => {
          set((state) => ({
            notifications: state.notifications.filter((n) => n.id !== id),
          }));
        }, duration);
      }

      return {
        notifications: [...state.notifications, newNotification],
      };
    }),

  removeNotification: (id) =>
    set((state) => ({
      notifications: state.notifications.filter((n) => n.id !== id),
    })),

  clearNotifications: () => set({ notifications: [] }),

  // Actions для модальных окон
  openModal: (type, data) =>
    set({
      modal: {
        type,
        data,
        isOpen: true,
      },
    }),

  closeModal: () =>
    set({
      modal: {
        type: null,
        isOpen: false,
        data: undefined,
      },
    }),

  // Actions для темы
  setTheme: (theme) => {
    set({ theme });
    // Сохраняем тему в localStorage
    if (typeof window !== 'undefined') {
      localStorage.setItem('theme', theme);
      // Применяем тему к документу
      if (theme === 'dark') {
        document.documentElement.classList.add('dark');
      } else {
        document.documentElement.classList.remove('dark');
      }
    }
  },

  // Actions для глобальной загрузки
  setGlobalLoading: (loading) => set({ globalLoading: loading }),
}));

// Селекторы для оптимизации рендеринга
export const selectNotifications = (state: UIState) => state.notifications;
export const selectModal = (state: UIState) => state.modal;
export const selectTheme = (state: UIState) => state.theme;
export const selectGlobalLoading = (state: UIState) => state.globalLoading;
export const selectIsModalOpen = (state: UIState) => state.modal.isOpen;
export const selectModalType = (state: UIState) => state.modal.type;
export const selectModalData = (state: UIState) => state.modal.data;

// Хелперы для удобного использования
export const showSuccessNotification = (message: string, description?: string) => {
  useUIStore.getState().showNotification({
    type: 'success',
    message,
    description,
  });
};

export const showErrorNotification = (message: string, description?: string) => {
  useUIStore.getState().showNotification({
    type: 'error',
    message,
    description,
  });
};

export const showWarningNotification = (message: string, description?: string) => {
  useUIStore.getState().showNotification({
    type: 'warning',
    message,
    description,
  });
};

export const showInfoNotification = (message: string, description?: string) => {
  useUIStore.getState().showNotification({
    type: 'info',
    message,
    description,
  });
};
