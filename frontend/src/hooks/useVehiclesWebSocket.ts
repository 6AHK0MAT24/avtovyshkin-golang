import { useEffect, useRef } from 'react';
import { vehiclesWsClient } from '../services/websocket';
import type { Vehicle } from '../types/vehicle';

// Track WebSocket initialization
let isVehiclesWsInitialized = false;

// Track registered event types to prevent duplicates
const registeredVehicleEventTypes = new Set<string>();

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
    // Skip if already registered for this event type
    if (isRegisteredRef.current) {
      return;
    }

    // Connect to WebSocket only once
    if (!isVehiclesWsInitialized) {
      console.log('Connecting to Vehicles WebSocket...');
      vehiclesWsClient.connect();
      isVehiclesWsInitialized = true;
    }

    // Check if this event type is already registered globally
    if (registeredVehicleEventTypes.has('vehicle.created')) {
      console.log('vehicle.created already registered');
      return;
    }

    console.log('Registering vehicle.created handler');
    const unsubscribe = vehiclesWsClient.on('vehicle.created', (data) => {
      console.log('vehicle.created event received:', data);
      handlerRef.current(data);
    });

    registeredVehicleEventTypes.add('vehicle.created');
    isRegisteredRef.current = true;

    return () => {
      unsubscribe();
      registeredVehicleEventTypes.delete('vehicle.created');
      isRegisteredRef.current = false;
    };
  }, []);};

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
    // Skip if already registered for this event type
    if (isRegisteredRef.current) {
      return;
    }

    // Connect to WebSocket only once
    if (!isVehiclesWsInitialized) {
      console.log('Connecting to Vehicles WebSocket...');
      vehiclesWsClient.connect();
      isVehiclesWsInitialized = true;
    }

    // Check if this event type is already registered globally
    if (registeredVehicleEventTypes.has('vehicle.updated')) {
      console.log('vehicle.updated already registered');
      return;
    }

    console.log('Registering vehicle.updated handler');
    const unsubscribe = vehiclesWsClient.on('vehicle.updated', (data) => {
      console.log('vehicle.updated event received:', data);
      handlerRef.current(data);
    });

    registeredVehicleEventTypes.add('vehicle.updated');
    isRegisteredRef.current = true;

    return () => {
      unsubscribe();
      registeredVehicleEventTypes.delete('vehicle.updated');
      isRegisteredRef.current = false;
    };
  }, []);};

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
    // Skip if already registered for this event type
    if (isRegisteredRef.current) {
      return;
    }

    // Connect to WebSocket only once
    if (!isVehiclesWsInitialized) {
      console.log('Connecting to Vehicles WebSocket...');
      vehiclesWsClient.connect();
      isVehiclesWsInitialized = true;
    }

    // Check if this event type is already registered globally
    if (registeredVehicleEventTypes.has('vehicle.deleted')) {
      console.log('vehicle.deleted already registered');
      return;
    }

    console.log('Registering vehicle.deleted handler');
    const unsubscribe = vehiclesWsClient.on('vehicle.deleted', (data) => {
      console.log('vehicle.deleted event received:', data);
      handlerRef.current(data);
    });

    registeredVehicleEventTypes.add('vehicle.deleted');
    isRegisteredRef.current = true;

    return () => {
      unsubscribe();
      registeredVehicleEventTypes.delete('vehicle.deleted');
      isRegisteredRef.current = false;
    };
  }, []);};

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