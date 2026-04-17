package websocket

import (
	"log"
	"net/http"

	"github.com/google/uuid"

	"vehicles-service/internal/models"
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

// BroadcastVehicleCreated broadcasts vehicle created event
func (m *Manager) BroadcastVehicleCreated(vehicle *models.Vehicle) error {
	log.Printf("Broadcasting vehicle.created event, clients: %d", m.hub.GetClientCount())
	return m.hub.BroadcastEvent("vehicle.created", vehicle.ToResponse())
}
// BroadcastVehicleUpdated broadcasts vehicle updated event
func (m *Manager) BroadcastVehicleUpdated(vehicle *models.Vehicle) error {
	log.Printf("Broadcasting vehicle.updated event, clients: %d", m.hub.GetClientCount())
	return m.hub.BroadcastEvent("vehicle.updated", vehicle.ToResponse())
}
// BroadcastVehicleDeleted broadcasts vehicle deleted event
func (m *Manager) BroadcastVehicleDeleted(vehicleID string) error {
	id, err := uuid.Parse(vehicleID)
	if err != nil {
		return err
	}
	event := map[string]interface{}{
		"id": id.String(),
	}
	return m.hub.BroadcastEvent("vehicle.deleted", event)
}

// GetConnectedClientsCount returns the number of connected clients
func (m *Manager) GetConnectedClientsCount() int {
	return m.hub.GetClientCount()
}
