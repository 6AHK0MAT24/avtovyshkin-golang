import { useEffect, useRef } from 'react';
import { wsClient } from '../services/websocket';

// Track registered event types to prevent duplicates
const registeredEventTypes = new Set<string>();
let isInitialized = false;

export const useWebSocket = (eventType: string, handler: (data: any) => void) => {
  const handlerRef = useRef(handler);
  const isRegisteredRef = useRef(false);

  // Update ref when handler changes
  useEffect(() => {
    handlerRef.current = handler;
  }, [handler]);

  useEffect(() => {
    // Skip if already registered for this event type
    if (isRegisteredRef.current) {
      return;
    }

    // Connect to WebSocket only once
    if (!isInitialized) {
      wsClient.connect();
      isInitialized = true;
    }

    // Check if this event type is already registered globally
    if (registeredEventTypes.has(eventType)) {
      return;
    }

    // Subscribe to event
    const unsubscribe = wsClient.on(eventType, (data) => {
      handlerRef.current(data);
    });

    registeredEventTypes.add(eventType);
    isRegisteredRef.current = true;

    // Cleanup on unmount
    return () => {
      unsubscribe();
      registeredEventTypes.delete(eventType);
      isRegisteredRef.current = false;
    };
  }, [eventType]);

  return {
    isConnected: wsClient.isConnected(),
    send: wsClient.send.bind(wsClient),
  };
};
