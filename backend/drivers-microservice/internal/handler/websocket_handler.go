package handler

import (
	"log"
	"net/http"

	"drivers-service/internal/websocket"
)

// WebSocketHandler handles WebSocket connections
type WebSocketHandler struct {
	wsManager *websocket.Manager
}

// NewWebSocketHandler creates a new WebSocket handler
func NewWebSocketHandler(wsManager *websocket.Manager) *WebSocketHandler {
	return &WebSocketHandler{
		wsManager: wsManager,
	}
}

// HandleWebSocket handles WebSocket connection upgrades
func (h *WebSocketHandler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	log.Printf("WebSocket connection attempt from %s", r.RemoteAddr)
	h.wsManager.HandleWebSocket(w, r)
}
