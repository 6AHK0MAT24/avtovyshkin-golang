import { useEffect, useRef } from 'react';
import { vehiclesWsClient } from '../services/websocket';
import type { Vehicle } from '../types/vehicle';

// Track registered event types to prevent duplicates
const registeredVehicleEventTypes = new Set<string>();
let isVehiclesWsInitialized = false;

/**
 * Хук для подписки на событие создания автовышки
 * @param handler - функция-обработчик, принимающая данные созданной автовышки
 */
export const useVehicleCreated = (handler: (data: Vehicle) => void) => {
  const handlerRef = useRef(handler);
  const isRegisteredRef = useRef(false);

  useEffect(() => {
    handlerRef.current = handler;
  }, [handler]);

  useEffect(() => {
    if (isRegisteredRef.current) {
      return;
    }

    if (!isVehiclesWsInitialized) {
      vehiclesWsClient.connect();
      isVehiclesWsInitialized = true;
    }

    if (registeredVehicleEventTypes.has('vehicle.created')) {
      return;
    }

    const unsubscribe = vehiclesWsClient.on('vehicle.created', (data) => {
      handlerRef.current(data);
    });

    registeredVehicleEventTypes.add('vehicle.created');
    isRegisteredRef.current = true;

    return () => {
      unsubscribe();
      registeredVehicleEventTypes.delete('vehicle.created');
      isRegisteredRef.current = false;
    };
  }, []);
};

/**
 * Хук для подписки на событие обновления автовышки
 * @param handler - функция-обработчик, принимающая данные обновленной автовышки
 */
export const useVehicleUpdated = (handler: (data: Vehicle) => void) => {
  const handlerRef = useRef(handler);
  const isRegisteredRef = useRef(false);

  useEffect(() => {
    handlerRef.current = handler;
  }, [handler]);

  useEffect(() => {
    if (isRegisteredRef.current) {
      return;
    }

    if (!isVehiclesWsInitialized) {
      vehiclesWsClient.connect();
      isVehiclesWsInitialized = true;
    }

    if (registeredVehicleEventTypes.has('vehicle.updated')) {
      return;
    }

    const unsubscribe = vehiclesWsClient.on('vehicle.updated', (data) => {
      handlerRef.current(data);
    });

    registeredVehicleEventTypes.add('vehicle.updated');
    isRegisteredRef.current = true;

    return () => {
      unsubscribe();
      registeredVehicleEventTypes.delete('vehicle.updated');
      isRegisteredRef.current = false;
    };
  }, []);
};

/**
 * Хук для подписки на событие удаления автовышки
 * @param handler - функция-обработчик, принимающая объект с id удаленной автовышки
 */
export const useVehicleDeleted = (handler: (data: { id: string }) => void) => {
  const handlerRef = useRef(handler);
  const isRegisteredRef = useRef(false);

  useEffect(() => {
    handlerRef.current = handler;
  }, [handler]);

  useEffect(() => {
    if (isRegisteredRef.current) {
      return;
    }

    if (!isVehiclesWsInitialized) {
      vehiclesWsClient.connect();
      isVehiclesWsInitialized = true;
    }

    if (registeredVehicleEventTypes.has('vehicle.deleted')) {
      return;
    }

    const unsubscribe = vehiclesWsClient.on('vehicle.deleted', (data) => {
      handlerRef.current(data);
    });

    registeredVehicleEventTypes.add('vehicle.deleted');
    isRegisteredRef.current = true;

    return () => {
      unsubscribe();
      registeredVehicleEventTypes.delete('vehicle.deleted');
      isRegisteredRef.current = false;
    };
  }, []);
};

/**
 * Комбинированный хук для подписки на все события автовышек
 * @param handlers - объект с обработчиками для каждого типа события
 */
export const useVehiclesWebSocket = (handlers: {
  onCreated?: (data: Vehicle) => void;
  onUpdated?: (data: Vehicle) => void;
  onDeleted?: (data: { id: string }) => void;
}) => {
  useVehicleCreated(handlers.onCreated || (() => {}));
  useVehicleUpdated(handlers.onUpdated || (() => {}));
  useVehicleDeleted(handlers.onDeleted || (() => {}));
};
