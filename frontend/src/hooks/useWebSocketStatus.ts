import { useState, useEffect } from 'react';
import { wsClient } from '../services/websocket';

/**
 * Хук для получения статуса WebSocket соединения
 * Не подписывается на события, только отслеживает состояние подключения
 */
export const useWebSocketStatus = () => {
  const [isConnected, setIsConnected] = useState(false);

  useEffect(() => {
    // Подключаемся к WebSocket
    wsClient.connect();

    // Проверяем начальный статус
    setIsConnected(wsClient.isConnected());

    // Подписываемся на изменения статуса соединения
    const handleConnectionChange = (connected: boolean) => {
      setIsConnected(connected);
    };

    // Подписываемся на события соединения
    const unsubscribeConnect = wsClient.on('connected', () => {
      handleConnectionChange(true);
    });

    const unsubscribeDisconnect = wsClient.on('disconnected', () => {
      handleConnectionChange(false);
    });

    // Периодическая проверка статуса (fallback)
    const intervalId = setInterval(() => {
      setIsConnected(wsClient.isConnected());
    }, 5000);

    return () => {
      unsubscribeConnect();
      unsubscribeDisconnect();
      clearInterval(intervalId);
      // Не отключаем WebSocket здесь, так как другие компоненты могут его использовать
    };
  }, []);

  return {
    isConnected,
    connect: () => wsClient.connect(),
    disconnect: () => wsClient.disconnect(),
  };
};
