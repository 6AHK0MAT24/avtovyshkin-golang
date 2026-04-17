package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"vehicles-service/internal/config"
	"vehicles-service/internal/handler"
	"vehicles-service/internal/middleware"
	"vehicles-service/internal/repository"
	"vehicles-service/internal/service"
	"vehicles-service/internal/websocket"

	"github.com/gorilla/mux"
)

func main() {
	// Load configuration
	cfg := config.Load()
	log.Printf("Starting Vehicles Service...")
	log.Printf("Environment: %s", cfg.Env)
	log.Printf("Database: %s:%s/%s", cfg.Database.Host, cfg.Database.Port, cfg.Database.DBName)

	// Initialize database connection
	db, err := repository.InitDB(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Initialize repository
	repo := repository.NewVehicleRepository(db)

	// Initialize WebSocket manager
	wsManager := websocket.NewManager()
	wsManager.Start()
	log.Println("WebSocket manager started")

	// Initialize service
	vehicleService := service.NewVehicleService(repo)

	// Initialize handlers
	uploadConfig := handler.UploadConfig{
		MaxUploadSize: cfg.Upload.MaxUploadSize,
		AllowedTypes:  cfg.Upload.AllowedTypes,
	}
	vehicleHandler := handler.NewVehicleHandler(vehicleService, wsManager, uploadConfig)
	wsHandler := handler.NewWebSocketHandler(wsManager)

	// Create router
	router := mux.NewRouter()

	// Apply middleware globally
	router.Use(middleware.RecoveryMiddleware)
	router.Use(middleware.LoggingMiddleware)

	// API routes with middleware
	api := router.PathPrefix("/api").Subrouter()
	api.Use(middleware.CORSMiddleware(cfg.CORS.AllowedOrigins))

	// Global OPTIONS handler for CORS preflight requests
	api.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Static files for uploads
	fs := http.FileServer(http.Dir("uploads"))
	router.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", fs))

	// Vehicle routes
	api.HandleFunc("/vehicles", vehicleHandler.CreateVehicle).Methods("POST", "OPTIONS")
	api.HandleFunc("/vehicles", vehicleHandler.GetVehicles).Methods("GET", "OPTIONS")
	api.HandleFunc("/vehicles/search", vehicleHandler.SearchVehicles).Methods("GET", "OPTIONS")
	api.HandleFunc("/vehicles/{id}", vehicleHandler.GetVehicle).Methods("GET", "OPTIONS")
	api.HandleFunc("/vehicles/{id}", vehicleHandler.UpdateVehicle).Methods("PUT", "OPTIONS")
	api.HandleFunc("/vehicles/{id}", vehicleHandler.DeleteVehicle).Methods("DELETE", "OPTIONS")

	// Vehicle images routes
	api.HandleFunc("/vehicles/{id}/images", vehicleHandler.UploadVehicleImages).Methods("POST", "OPTIONS")
	api.HandleFunc("/vehicles/{id}/images/{index}", vehicleHandler.DeleteVehicleImage).Methods("DELETE", "OPTIONS")
	api.HandleFunc("/vehicles/{id}/main-image", vehicleHandler.SetMainImage).Methods("PUT", "OPTIONS")

	// WebSocket route - without CORS middleware to allow proper upgrade
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
