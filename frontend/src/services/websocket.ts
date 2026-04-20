type WebSocketMessage = {
  type: string;
  data: any;
};

type WebSocketEventHandler = (data: any) => void;

class WebSocketClient {
  private ws: WebSocket | null = null;
  private url: string;
  private reconnectAttempts = 0;
  private maxReconnectAttempts = 5;
  private reconnectDelay = 3000;
  private eventHandlers: Map<string, Set<WebSocketEventHandler>> = new Map();

  constructor(url: string) {
    this.url = url;
  }

  connect(): void {
    if (this.ws?.readyState === WebSocket.OPEN) {
      return;
    }

    try {
      console.log(`Connecting to WebSocket: ${this.url}`);
      this.ws = new WebSocket(this.url);

      this.ws.onopen = () => {
        console.log(`WebSocket connected: ${this.url}`);
        this.reconnectAttempts = 0;
      };

      this.ws.onmessage = (event) => {
        try {
          const message: WebSocketMessage = JSON.parse(event.data);
          console.log(`WebSocket message received from ${this.url}:`, message);
          this.handleMessage(message);
        } catch (error) {
          console.error('Failed to parse WebSocket message:', error);
        }
      };

      this.ws.onclose = () => {
        console.log(`WebSocket closed: ${this.url}`);
        this.attemptReconnect();
      };

      this.ws.onerror = (error) => {
        console.error(`WebSocket error (${this.url}):`, error);
      };
    } catch (error) {
      console.error(`Failed to create WebSocket connection (${this.url}):`, error);
      this.attemptReconnect();
    }  }

  reconnect(): void {
    if (this.ws) {
      this.ws.close();
      this.ws = null;
    }
    this.connect();
  }  disconnect(): void {
    if (this.ws) {
      this.ws.close();
      this.ws = null;
    }
  }

  private handleMessage(message: WebSocketMessage): void {
    const handlers = this.eventHandlers.get(message.type);
    if (handlers) {
      handlers.forEach((handler) => {
        try {
          handler(message.data);
        } catch (error) {
          console.error('Error in event handler:', error);
        }
      });
    }
  }

  private attemptReconnect(): void {
    if (this.reconnectAttempts < this.maxReconnectAttempts) {
      this.reconnectAttempts++;
      console.log(`Attempting to reconnect (${this.reconnectAttempts}/${this.maxReconnectAttempts})...`);
      setTimeout(() => {
        this.connect();
      }, this.reconnectDelay);
    }
  }

  on(eventType: string, handler: WebSocketEventHandler): () => void {
    if (!this.eventHandlers.has(eventType)) {
      this.eventHandlers.set(eventType, new Set());
    }

    const handlers = this.eventHandlers.get(eventType)!;

    // Check if this exact handler is already registered
    if (handlers.has(handler)) {
      return () => {}; // Return empty unsubscribe function
    }

    handlers.add(handler);

    // Return unsubscribe function
    return () => {
      const handlers = this.eventHandlers.get(eventType);
      if (handlers) {
        handlers.delete(handler);
        if (handlers.size === 0) {
          this.eventHandlers.delete(eventType);
        }
      }
    };
  }

  send(data: any): void {    if (this.ws?.readyState === WebSocket.OPEN) {
      this.ws.send(JSON.stringify(data));
    } else {
      console.warn('WebSocket is not connected');
    }
  }

  isConnected(): boolean {
    return this.ws?.readyState === WebSocket.OPEN;
  }
}

const WS_BASE_URL = import.meta.env.VITE_WS_BASE_URL || 'ws://localhost:8082';
export const wsClient = new WebSocketClient(`${WS_BASE_URL}/ws`);

// Vehicles WebSocket client for vehicles microservice (port 8081)
const VEHICLES_WS_BASE_URL = import.meta.env.VITE_VEHICLES_WS_BASE_URL || 'ws://localhost:8081';
export const vehiclesWsClient = new WebSocketClient(`${VEHICLES_WS_BASE_URL}/ws`);
