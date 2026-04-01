package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"drivers-service/internal/config"
	"drivers-service/internal/handler"
	"drivers-service/internal/middleware"
	"drivers-service/internal/repository"
	"drivers-service/internal/service"
	"drivers-service/internal/websocket"

	"github.com/gorilla/mux"
)

func main() {
	// Load configuration
	cfg := config.Load()
	log.Printf("Starting Drivers Service...")
	log.Printf("Environment: %s", cfg.Env)
	log.Printf("Database: %s:%s/%s", cfg.Database.Host, cfg.Database.Port, cfg.Database.DBName)

	// Initialize repository
	repo, err := repository.NewDriverRepository(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to initialize repository: %v", err)
	}
	defer repo.Close()

	// Initialize WebSocket manager
	wsManager := websocket.NewManager()
	wsManager.Start()
	log.Println("WebSocket manager started")

	// Initialize service
	driverService := service.NewDriverService(repo)

	// Initialize handlers
	uploadConfig := handler.UploadConfig{
		MaxUploadSize: cfg.Upload.MaxUploadSize,
		AllowedTypes:  cfg.Upload.AllowedTypes,
	}
	driverHandler := handler.NewDriverHandler(driverService, wsManager, uploadConfig)
	wsHandler := handler.NewWebSocketHandler(wsManager)

	// Create router
	router := mux.NewRouter()

	// Apply middleware
	router.Use(middleware.RecoveryMiddleware)
	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.CORSMiddleware(cfg.CORS.AllowedOrigins))

	// Global OPTIONS handler for CORS preflight requests
	router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Static files for uploads
	fs := http.FileServer(http.Dir("uploads"))
	router.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", fs))

	// API routes
	api := router.PathPrefix("/api").Subrouter()
	// Driver routes
	api.HandleFunc("/drivers", driverHandler.CreateDriver).Methods("POST", "OPTIONS")
	api.HandleFunc("/drivers", driverHandler.GetDrivers).Methods("GET", "OPTIONS")
	api.HandleFunc("/drivers/search", driverHandler.SearchDrivers).Methods("GET", "OPTIONS")
	api.HandleFunc("/drivers/{id}", driverHandler.GetDriver).Methods("GET", "OPTIONS")
	api.HandleFunc("/drivers/{id}", driverHandler.UpdateDriver).Methods("PUT", "OPTIONS")
	api.HandleFunc("/drivers/{id}", driverHandler.DeleteDriver).Methods("DELETE", "OPTIONS")

	// File upload routes
	api.HandleFunc("/drivers/{id}/photo", driverHandler.UploadDriverPhoto).Methods("POST", "OPTIONS")
	api.HandleFunc("/drivers/{id}/photo", driverHandler.DeleteDriverPhoto).Methods("DELETE", "OPTIONS")
	api.HandleFunc("/drivers/{id}/license", driverHandler.UploadDriverLicense).Methods("POST", "OPTIONS")
	api.HandleFunc("/drivers/{id}/passport", driverHandler.UploadDriverPassport).Methods("POST", "OPTIONS")

	// WebSocket route
	router.HandleFunc("/ws", wsHandler.HandleWebSocket)
	// Health check
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	// Create server
	server := &http.Server{
		Addr:         cfg.Server.Host + ":" + cfg.Server.Port,
		Handler:      router,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	// Start server in goroutine
	go func() {
		log.Printf("Server listening on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server stopped")
}
