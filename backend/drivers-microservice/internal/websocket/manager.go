package websocket

import (
	"log"
	"net/http"

	"github.com/google/uuid"

	"drivers-service/internal/models"
)
// Manager manages WebSocket connections and broadcasts events
type Manager struct {
	hub *Hub
}

// NewManager creates a new WebSocket manager
func NewManager() *Manager {
	return &Manager{
		hub: NewHub(),
	}
}

// GetHub returns the hub
func (m *Manager) GetHub() *Hub {
	return m.hub
}

// Start starts the hub
func (m *Manager) Start() {
	go m.hub.Run()
}

// HandleWebSocket handles WebSocket connection upgrades
func (m *Manager) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		return
	}

	client := NewClient(m.hub, conn)
	m.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump()
}

// BroadcastDriverCreated broadcasts driver created event
func (m *Manager) BroadcastDriverCreated(driver *models.Driver) error {
	return m.hub.BroadcastEvent("driver.created", driver.ToResponse())
}

// BroadcastDriverUpdated broadcasts driver updated event
func (m *Manager) BroadcastDriverUpdated(driver *models.Driver) error {
	return m.hub.BroadcastEvent("driver.updated", driver.ToResponse())
}

// BroadcastDriverDeleted broadcasts driver deleted event
func (m *Manager) BroadcastDriverDeleted(driverID string) error {
	id, err := uuid.Parse(driverID)
	if err != nil {
		return err
	}
	event := map[string]interface{}{
		"id": id.String(),
	}
	return m.hub.BroadcastEvent("driver.deleted", event)
}
// GetConnectedClientsCount returns the number of connected clients
func (m *Manager) GetConnectedClientsCount() int {
	return m.hub.GetClientCount()
}
