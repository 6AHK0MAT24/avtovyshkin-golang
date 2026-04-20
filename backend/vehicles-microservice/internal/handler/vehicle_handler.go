package handler

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"vehicles-service/internal/models"
	"vehicles-service/internal/service"
	"vehicles-service/internal/utils"
	"vehicles-service/internal/websocket"
	"github.com/google/uuid"
)

// VehicleHandler handles HTTP requests for vehicles
type VehicleHandler struct {
	service      service.VehicleService
	wsManager    *websocket.Manager
	uploadConfig UploadConfig
}

type UploadConfig struct {
	MaxUploadSize  int64
	AllowedTypes   []string
}

// NewVehicleHandler creates a new vehicle handler
func NewVehicleHandler(service service.VehicleService, wsManager *websocket.Manager, uploadConfig UploadConfig) *VehicleHandler {
	return &VehicleHandler{
		service:      service,
		wsManager:    wsManager,
		uploadConfig: uploadConfig,
	}
}

// CreateVehicle handles POST /api/vehicles
func (h *VehicleHandler) CreateVehicle(w http.ResponseWriter, r *http.Request) {
	var req models.CreateVehicleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	vehicle, err := h.service.CreateVehicle(context.Background(), &req)
	if err != nil {
		log.Printf("Error creating vehicle: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Broadcast event
	if err := h.wsManager.BroadcastVehicleCreated(vehicle); err != nil {
		log.Printf("Error broadcasting vehicle created event: %v", err)
	}

	respondJSON(w, http.StatusCreated, vehicle.ToResponse())
}

// GetVehicle handles GET /api/vehicles/{id}
func (h *VehicleHandler) GetVehicle(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/vehicles/")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid vehicle ID", http.StatusBadRequest)
		return
	}

	vehicle, err := h.service.GetVehicle(context.Background(), id)
	if err != nil {
		log.Printf("Error getting vehicle: %v", err)
		http.Error(w, "Vehicle not found", http.StatusNotFound)
		return
	}

	respondJSON(w, http.StatusOK, vehicle.ToResponse())
}

// GetVehicles handles GET /api/vehicles
func (h *VehicleHandler) GetVehicles(w http.ResponseWriter, r *http.Request) {
	// Parse pagination parameters
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	vehicles, total, err := h.service.GetVehicles(context.Background(), page, pageSize)
	if err != nil {
		log.Printf("Error getting vehicles: %v", err)
		http.Error(w, "Failed to get vehicles", http.StatusInternalServerError)
		return
	}

	// Convert to response format
	vehicleResponses := make([]models.VehicleResponse, len(vehicles))
	for i, vehicle := range vehicles {
		vehicleResponses[i] = vehicle.ToResponse()
	}

	response := models.VehicleListResponse{
		Vehicles: vehicleResponses,
		Total:    total,
		Page:     page,
		PerPage:  pageSize,
	}

	respondJSON(w, http.StatusOK, response)
}

// UpdateVehicle handles PUT /api/vehicles/{id}
func (h *VehicleHandler) UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/vehicles/")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid vehicle ID", http.StatusBadRequest)
		return
	}

	var req models.UpdateVehicleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	vehicle, err := h.service.UpdateVehicle(context.Background(), id, &req)
	if err != nil {
		log.Printf("Error updating vehicle: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Broadcast event
	if err := h.wsManager.BroadcastVehicleUpdated(vehicle); err != nil {
		log.Printf("Error broadcasting vehicle updated event: %v", err)
	}

	respondJSON(w, http.StatusOK, vehicle.ToResponse())
}

// DeleteVehicle handles DELETE /api/vehicles/{id}
func (h *VehicleHandler) DeleteVehicle(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/vehicles/")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid vehicle ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteVehicle(context.Background(), id); err != nil {
		log.Printf("Error deleting vehicle: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Broadcast event
	if err := h.wsManager.BroadcastVehicleDeleted(id.String()); err != nil {
		log.Printf("Error broadcasting vehicle deleted event: %v", err)
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "Vehicle deleted successfully"})
}

// SearchVehicles handles GET /api/vehicles/search
func (h *VehicleHandler) SearchVehicles(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	vehicles, total, err := h.service.SearchVehicles(context.Background(), query, page, pageSize)
	if err != nil {
		log.Printf("Error searching vehicles: %v", err)
		http.Error(w, "Failed to search vehicles", http.StatusInternalServerError)
		return
	}

	// Convert to response format
	vehicleResponses := make([]models.VehicleResponse, len(vehicles))
	for i, vehicle := range vehicles {
		vehicleResponses[i] = vehicle.ToResponse()
	}

	response := models.VehicleListResponse{
		Vehicles: vehicleResponses,
		Total:    total,
		Page:     page,
		PerPage:  pageSize,
	}

	respondJSON(w, http.StatusOK, response)
}

// UploadVehicleImages handles POST /api/vehicles/{id}/images
func (h *VehicleHandler) UploadVehicleImages(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/vehicles/")
	idStr = strings.TrimSuffix(idStr, "/images")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid vehicle ID", http.StatusBadRequest)
		return
	}

	// Parse multipart form
	if err := r.ParseMultipartForm(h.uploadConfig.MaxUploadSize); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	// Get all files from form
	files := r.MultipartForm.File["files"]
	if len(files) == 0 {
		http.Error(w, "No files provided", http.StatusBadRequest)
		return
	}

	// Read file data
	fileDataList := make([]service.FileData, 0, len(files))
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			log.Printf("Error opening file %s: %v", fileHeader.Filename, err)
			continue
		}

		// Validate file type
		if !utils.ValidateFileType(fileHeader.Filename, h.uploadConfig.AllowedTypes) {
			file.Close()
			http.Error(w, "Invalid file type: "+fileHeader.Filename, http.StatusBadRequest)
			return
		}

		data, err := io.ReadAll(file)
		file.Close()
		if err != nil {
			log.Printf("Error reading file %s: %v", fileHeader.Filename, err)
			continue
		}

		fileDataList = append(fileDataList, service.FileData{
			Data:     data,
			Filename: fileHeader.Filename,
		})
	}

	if len(fileDataList) == 0 {
		http.Error(w, "No valid files provided", http.StatusBadRequest)
		return
	}

	// Upload images
	imagePaths, err := h.service.UploadVehicleImages(context.Background(), id, fileDataList)
	if err != nil {
		log.Printf("Error uploading vehicle images: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	respondJSON(w, http.StatusOK, map[string]interface{}{
		"message": "Images uploaded successfully",
		"images":  imagePaths,
	})
}

// DeleteVehicleImage handles DELETE /api/vehicles/{id}/images/{index}
func (h *VehicleHandler) DeleteVehicleImage(w http.ResponseWriter, r *http.Request) {
	// Parse path: /api/vehicles/{id}/images/{index}
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 6 {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}

	idStr := pathParts[4]
	indexStr := pathParts[6]

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid vehicle ID", http.StatusBadRequest)
		return
	}

	index, err := strconv.Atoi(indexStr)
	if err != nil {
		http.Error(w, "Invalid image index", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteVehicleImage(context.Background(), id, index); err != nil {
		log.Printf("Error deleting vehicle image: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "Image deleted successfully"})
}

// SetMainImage handles PUT /api/vehicles/{id}/main-image
func (h *VehicleHandler) SetMainImage(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/vehicles/")
	idStr = strings.TrimSuffix(idStr, "/main-image")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid vehicle ID", http.StatusBadRequest)
		return
	}

	var req struct {
		Index int `json:"index"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Индекс уже 1-based, используем как есть
	if req.Index < 1 {
		http.Error(w, "Invalid image index", http.StatusBadRequest)
		return
	}

	if err := h.service.SetMainImageIndex(context.Background(), id, req.Index); err != nil {
		log.Printf("Error setting main image: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "Main image set successfully"})
}
