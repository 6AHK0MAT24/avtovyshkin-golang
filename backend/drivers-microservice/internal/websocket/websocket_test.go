package websocket

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"drivers-service/internal/models"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHub_NewHub(t *testing.T) {
	hub := NewHub()

	assert.NotNil(t, hub)
	assert.NotNil(t, hub.clients)
	assert.NotNil(t, hub.broadcast)
	assert.NotNil(t, hub.register)
	assert.NotNil(t, hub.unregister)
	assert.Equal(t, 0, hub.GetClientCount())
}

func TestHub_RegisterClient(t *testing.T) {
	hub := NewHub()
	go hub.Run()

	client := &Client{
		hub:  hub,
		send: make(chan []byte, 256),
	}

	hub.register <- client

	// Wait for registration
	time.Sleep(100 * time.Millisecond)

	assert.Equal(t, 1, hub.GetClientCount())
}

func TestHub_UnregisterClient(t *testing.T) {
	hub := NewHub()
	go hub.Run()

	client := &Client{
		hub:  hub,
		send: make(chan []byte, 256),
	}

	hub.register <- client
	time.Sleep(100 * time.Millisecond)

	assert.Equal(t, 1, hub.GetClientCount())

	hub.unregister <- client
	time.Sleep(100 * time.Millisecond)

	assert.Equal(t, 0, hub.GetClientCount())
}

func TestHub_BroadcastEvent(t *testing.T) {
	hub := NewHub()
	go hub.Run()

	// Create test clients
	client1 := &Client{
		hub:  hub,
		send: make(chan []byte, 256),
	}
	client2 := &Client{
		hub:  hub,
		send: make(chan []byte, 256),
	}

	hub.register <- client1
	hub.register <- client2
	time.Sleep(100 * time.Millisecond)

	assert.Equal(t, 2, hub.GetClientCount())

	// Broadcast event
	eventData := map[string]interface{}{
		"driverId": "test-id",
		"action":   "created",
	}

	err := hub.BroadcastEvent("driver.created", eventData)
	require.NoError(t, err)

	// Wait for broadcast
	time.Sleep(100 * time.Millisecond)

	// Check if clients received the message
	select {
	case msg := <-client1.send:
		var event map[string]interface{}
		err := json.Unmarshal(msg, &event)
		require.NoError(t, err)
		assert.Equal(t, "driver.created", event["type"])
		assert.NotNil(t, event["data"])
	case <-time.After(1 * time.Second):
		t.Fatal("Client 1 did not receive message")
	}

	select {
	case msg := <-client2.send:
		var event map[string]interface{}
		err := json.Unmarshal(msg, &event)
		require.NoError(t, err)
		assert.Equal(t, "driver.created", event["type"])
		assert.NotNil(t, event["data"])
	case <-time.After(1 * time.Second):
		t.Fatal("Client 2 did not receive message")
	}
}

func TestHub_BroadcastEvent_NoClients(t *testing.T) {
	hub := NewHub()
	go hub.Run()

	eventData := map[string]interface{}{
		"driverId": "test-id",
		"action":   "created",
	}

	err := hub.BroadcastEvent("driver.created", eventData)
	assert.NoError(t, err)
}

func TestHub_GetClientCount(t *testing.T) {
	hub := NewHub()
	go hub.Run()

	assert.Equal(t, 0, hub.GetClientCount())

	client := &Client{
		hub:  hub,
		send: make(chan []byte, 256),
	}

	hub.register <- client
	time.Sleep(100 * time.Millisecond)

	assert.Equal(t, 1, hub.GetClientCount())
}

func TestManager_NewManager(t *testing.T) {
	manager := NewManager()

	assert.NotNil(t, manager)
	assert.NotNil(t, manager.GetHub())
	assert.Equal(t, 0, manager.GetConnectedClientsCount())
}

func TestManager_Start(t *testing.T) {
	manager := NewManager()

	// Should not panic
	manager.Start()

	time.Sleep(100 * time.Millisecond)

	assert.NotNil(t, manager.GetHub())
}

func TestManager_BroadcastDriverCreated(t *testing.T) {
	manager := NewManager()
	manager.Start()

	driver := &models.Driver{
		ID:        uuid.New(),
		FirstName: "Ivan",
		LastName:  "Ivanov",
	}

	err := manager.BroadcastDriverCreated(driver)
	assert.NoError(t, err)
}

func TestManager_BroadcastDriverUpdated(t *testing.T) {
	manager := NewManager()
	manager.Start()

	driver := &models.Driver{
		ID:        uuid.New(),
		FirstName: "Petr",
		LastName:  "Petrov",
	}

	err := manager.BroadcastDriverUpdated(driver)
	assert.NoError(t, err)
}

func TestManager_BroadcastDriverDeleted(t *testing.T) {
	manager := NewManager()
	manager.Start()

	driverID := "550e8400-e29b-41d4-a716-446655440000"

	err := manager.BroadcastDriverDeleted(driverID)
	assert.NoError(t, err)
}

func TestManager_BroadcastDriverDeleted_InvalidID(t *testing.T) {
	manager := NewManager()
	manager.Start()

	driverID := "invalid-uuid"

	err := manager.BroadcastDriverDeleted(driverID)
	assert.Error(t, err)
}

func TestManager_GetConnectedClientsCount(t *testing.T) {
	manager := NewManager()
	manager.Start()

	assert.Equal(t, 0, manager.GetConnectedClientsCount())

	client := &Client{
		hub:  manager.GetHub(),
		send: make(chan []byte, 256),
	}

	manager.GetHub().register <- client
	time.Sleep(100 * time.Millisecond)

	assert.Equal(t, 1, manager.GetConnectedClientsCount())
}

func TestManager_HandleWebSocket_Success(t *testing.T) {
	manager := NewManager()
	manager.Start()

	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		manager.HandleWebSocket(w, r)
	}))
	defer server.Close()

	// Convert http:// to ws://
	wsURL := "ws" + strings.TrimPrefix(server.URL, "http")

	// Connect as WebSocket client
	dialer := websocket.Dialer{}
	conn, _, err := dialer.Dial(wsURL, nil)
	require.NoError(t, err)
	defer conn.Close()

	// Wait for connection to be registered
	time.Sleep(100 * time.Millisecond)

	// Check if client is registered
	assert.Equal(t, 1, manager.GetConnectedClientsCount())
}

func TestManager_HandleWebSocket_MultipleClients(t *testing.T) {
	manager := NewManager()
	manager.Start()

	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		manager.HandleWebSocket(w, r)
	}))
	defer server.Close()

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http")

	// Connect multiple clients
	dialer := websocket.Dialer{}
	conn1, _, err := dialer.Dial(wsURL, nil)
	require.NoError(t, err)
	defer conn1.Close()

	conn2, _, err := dialer.Dial(wsURL, nil)
	require.NoError(t, err)
	defer conn2.Close()

	conn3, _, err := dialer.Dial(wsURL, nil)
	require.NoError(t, err)
	defer conn3.Close()

	// Wait for connections to be registered
	time.Sleep(100 * time.Millisecond)

	// Check if all clients are registered
	assert.Equal(t, 3, manager.GetConnectedClientsCount())
}

func TestManager_HandleWebSocket_BroadcastToClients(t *testing.T) {
	manager := NewManager()
	manager.Start()

	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		manager.HandleWebSocket(w, r)
	}))
	defer server.Close()

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http")

	// Connect clients
	dialer := websocket.Dialer{}
	conn1, _, err := dialer.Dial(wsURL, nil)
	require.NoError(t, err)
	defer conn1.Close()

	conn2, _, err := dialer.Dial(wsURL, nil)
	require.NoError(t, err)
	defer conn2.Close()

	// Wait for connections
	time.Sleep(100 * time.Millisecond)

	// Broadcast event
	driver := &models.Driver{
		ID:        uuid.New(),
		FirstName: "Ivan",
		LastName:  "Ivanov",
	}

	err = manager.BroadcastDriverCreated(driver)
	require.NoError(t, err)

	// Check if clients received the message
	for i, conn := range []*websocket.Conn{conn1, conn2} {
		conn.SetReadDeadline(time.Now().Add(2 * time.Second))
		_, message, err := conn.ReadMessage()
		require.NoError(t, err, "Client %d did not receive message", i+1)

		var event map[string]interface{}
		err = json.Unmarshal(message, &event)
		require.NoError(t, err)
		assert.Equal(t, "driver.created", event["type"])
		assert.NotNil(t, event["data"])
	}
}

func TestManager_HandleWebSocket_ClientDisconnection(t *testing.T) {
	manager := NewManager()
	manager.Start()

	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		manager.HandleWebSocket(w, r)
	}))
	defer server.Close()

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http")

	// Connect client
	dialer := websocket.Dialer{}
	conn, _, err := dialer.Dial(wsURL, nil)
	require.NoError(t, err)

	// Wait for connection
	time.Sleep(100 * time.Millisecond)

	assert.Equal(t, 1, manager.GetConnectedClientsCount())

	// Close connection
	conn.Close()

	// Wait for unregistration
	time.Sleep(100 * time.Millisecond)

	assert.Equal(t, 0, manager.GetConnectedClientsCount())
}

func TestHub_BroadcastEvent_ClientBufferFull(t *testing.T) {
	hub := NewHub()
	go hub.Run()

	// Create client with small buffer
	client := &Client{
		hub:  hub,
		send: make(chan []byte, 1), // Small buffer
	}

	hub.register <- client
	time.Sleep(100 * time.Millisecond)

	// Fill the buffer
	client.send <- []byte("message1")

	// Broadcast another message (should not block)
	eventData := map[string]interface{}{
		"driverId": "test-id",
		"action":   "created",
	}

	err := hub.BroadcastEvent("driver.created", eventData)
	assert.NoError(t, err)

	// Wait for broadcast
	time.Sleep(100 * time.Millisecond)

	// Client should be removed due to full buffer
	assert.Equal(t, 0, hub.GetClientCount())
}

func TestManager_BroadcastDriverCreated_EventStructure(t *testing.T) {
	manager := NewManager()
	manager.Start()

	driver := &models.Driver{
		ID:        uuid.New(),
		FirstName: "Ivan",
		LastName:  "Ivanov",
		Phone:     "+79001234567",
	}

	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		manager.HandleWebSocket(w, r)
	}))
	defer server.Close()

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http")

	// Connect client
	dialer := websocket.Dialer{}
	conn, _, err := dialer.Dial(wsURL, nil)
	require.NoError(t, err)
	defer conn.Close()

	// Wait for connection
	time.Sleep(100 * time.Millisecond)

	// Broadcast event
	err = manager.BroadcastDriverCreated(driver)
	require.NoError(t, err)

	// Receive and validate event
	conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	_, message, err := conn.ReadMessage()
	require.NoError(t, err)

	var event map[string]interface{}
	err = json.Unmarshal(message, &event)
	require.NoError(t, err)

	assert.Equal(t, "driver.created", event["type"])
	assert.NotNil(t, event["data"])

	data := event["data"].(map[string]interface{})
	assert.Equal(t, driver.ID.String(), data["driverId"])
	assert.Equal(t, "created", data["action"])
}
