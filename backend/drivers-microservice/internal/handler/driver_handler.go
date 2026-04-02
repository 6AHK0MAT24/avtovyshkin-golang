package handler

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"drivers-service/internal/models"
	"drivers-service/internal/service"
	"drivers-service/internal/utils"
	"drivers-service/internal/websocket"
	"github.com/google/uuid"
)

// DriverHandler handles HTTP requests for drivers
type DriverHandler struct {
	service      service.DriverService
	wsManager    *websocket.Manager
	uploadConfig UploadConfig
}

type UploadConfig struct {
	MaxUploadSize  int64
	AllowedTypes   []string
}

// NewDriverHandler creates a new driver handler
func NewDriverHandler(service service.DriverService, wsManager *websocket.Manager, uploadConfig UploadConfig) *DriverHandler {
	return &DriverHandler{
		service:      service,
		wsManager:    wsManager,
		uploadConfig: uploadConfig,
	}
}

// CreateDriver handles POST /api/drivers
func (h *DriverHandler) CreateDriver(w http.ResponseWriter, r *http.Request) {
	var req models.CreateDriverRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	driver, err := h.service.CreateDriver(context.Background(), &req)
	if err != nil {
		log.Printf("Error creating driver: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Broadcast event
	if err := h.wsManager.BroadcastDriverCreated(driver); err != nil {
		log.Printf("Error broadcasting driver created event: %v", err)
	}

	respondJSON(w, http.StatusCreated, driver.ToResponse())
}

// GetDriver handles GET /api/drivers/{id}
func (h *DriverHandler) GetDriver(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/drivers/")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid driver ID", http.StatusBadRequest)
		return
	}

	driver, err := h.service.GetDriver(context.Background(), id)
	if err != nil {
		log.Printf("Error getting driver: %v", err)
		http.Error(w, "Driver not found", http.StatusNotFound)
		return
	}

	respondJSON(w, http.StatusOK, driver.ToResponse())
}

// GetDrivers handles GET /api/drivers
func (h *DriverHandler) GetDrivers(w http.ResponseWriter, r *http.Request) {
	// Parse pagination parameters
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	drivers, total, err := h.service.GetDrivers(context.Background(), page, pageSize)
	if err != nil {
		log.Printf("Error getting drivers: %v", err)
		http.Error(w, "Failed to get drivers", http.StatusInternalServerError)
		return
	}

	// Convert to response format
	driverResponses := make([]models.DriverResponse, len(drivers))
	for i, driver := range drivers {
		driverResponses[i] = driver.ToResponse()
	}

	response := models.DriverListResponse{
		Drivers: driverResponses,
		Total:   total,
		Page:    page,
		PerPage: pageSize,
	}

	respondJSON(w, http.StatusOK, response)
}

// UpdateDriver handles PUT /api/drivers/{id}
func (h *DriverHandler) UpdateDriver(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/drivers/")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid driver ID", http.StatusBadRequest)
		return
	}

	var req models.UpdateDriverRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	driver, err := h.service.UpdateDriver(context.Background(), id, &req)
	if err != nil {
		log.Printf("Error updating driver: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Broadcast event
	if err := h.wsManager.BroadcastDriverUpdated(driver); err != nil {
		log.Printf("Error broadcasting driver updated event: %v", err)
	}

	respondJSON(w, http.StatusOK, driver.ToResponse())
}

// DeleteDriver handles DELETE /api/drivers/{id}
func (h *DriverHandler) DeleteDriver(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/drivers/")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid driver ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteDriver(context.Background(), id); err != nil {
		log.Printf("Error deleting driver: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Broadcast event
	if err := h.wsManager.BroadcastDriverDeleted(id.String()); err != nil {
		log.Printf("Error broadcasting driver deleted event: %v", err)
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "Driver deleted successfully"})
}

// SearchDrivers handles GET /api/drivers/search
func (h *DriverHandler) SearchDrivers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	statusStr := r.URL.Query().Get("status")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	var status *models.DriverStatus
	if statusStr != "" {
		s := models.DriverStatus(statusStr)
		if s.IsValid() {
			status = &s
		}
	}

	drivers, total, err := h.service.SearchDrivers(context.Background(), query, status, page, pageSize)
	if err != nil {
		log.Printf("Error searching drivers: %v", err)
		http.Error(w, "Failed to search drivers", http.StatusInternalServerError)
		return
	}

	// Convert to response format
	driverResponses := make([]models.DriverResponse, len(drivers))
	for i, driver := range drivers {
		driverResponses[i] = driver.ToResponse()
	}

	response := models.DriverListResponse{
		Drivers: driverResponses,
		Total:   total,
		Page:    page,
		PerPage: pageSize,
	}

	respondJSON(w, http.StatusOK, response)
}

// UploadDriverPhoto handles POST /api/drivers/{id}/photo
func (h *DriverHandler) UploadDriverPhoto(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/drivers/")
	idStr = strings.TrimSuffix(idStr, "/photo")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid driver ID", http.StatusBadRequest)
		return
	}

	// Parse multipart form
	if err := r.ParseMultipartForm(h.uploadConfig.MaxUploadSize); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "No file provided", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Validate file type
	if !utils.ValidateFileType(header.Filename, h.uploadConfig.AllowedTypes) {
		http.Error(w, "Invalid file type", http.StatusBadRequest)
		return
	}

	// Read file content
	fileData, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	// Validate file size
	if int64(len(fileData)) > h.uploadConfig.MaxUploadSize {
		http.Error(w, "File too large", http.StatusBadRequest)
		return
	}

	// Upload file
	log.Printf("Uploading photo for driver %s, filename: %s, size: %d bytes", id, header.Filename, len(fileData))
	filePath, err := h.service.UploadDriverPhoto(context.Background(), id, fileData, header.Filename)
	if err != nil {
		log.Printf("Error uploading driver photo: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Photo uploaded successfully: %s", filePath)

	response := models.FileUploadResponse{
		FilePath: filePath,
		FileName: header.Filename,
		FileSize: int64(len(fileData)),
	}

	respondJSON(w, http.StatusOK, response)}

// UploadDriverLicense handles POST /api/drivers/{id}/license
func (h *DriverHandler) UploadDriverLicense(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/drivers/")
	idStr = strings.TrimSuffix(idStr, "/license")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid driver ID", http.StatusBadRequest)
		return
	}

	// Parse multipart form
	if err := r.ParseMultipartForm(h.uploadConfig.MaxUploadSize); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "No file provided", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Get file type (photo or scan)
	fileType := r.FormValue("fileType")
	if fileType != "photo" && fileType != "scan" {
		http.Error(w, "Invalid file type", http.StatusBadRequest)
		return
	}

	// Validate file type
	if !utils.ValidateFileType(header.Filename, h.uploadConfig.AllowedTypes) {
		http.Error(w, "Invalid file type", http.StatusBadRequest)
		return
	}

	// Read file content
	fileData, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	// Validate file size
	if int64(len(fileData)) > h.uploadConfig.MaxUploadSize {
		http.Error(w, "File too large", http.StatusBadRequest)
		return
	}

	// Upload file
	filePath, err := h.service.UploadDriverLicense(context.Background(), id, fileData, header.Filename, fileType)
	if err != nil {
		log.Printf("Error uploading driver license: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.FileUploadResponse{
		FilePath: filePath,
		FileName: header.Filename,
		FileSize: int64(len(fileData)),
	}

	respondJSON(w, http.StatusOK, response)
}
// UploadDriverPassport handles POST /api/drivers/{id}/passport
func (h *DriverHandler) UploadDriverPassport(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/drivers/")
	idStr = strings.TrimSuffix(idStr, "/passport")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid driver ID", http.StatusBadRequest)
		return
	}

	// Parse multipart form
	if err := r.ParseMultipartForm(h.uploadConfig.MaxUploadSize); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "No file provided", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Get file type (photo or scan)
	fileType := r.FormValue("fileType")
	if fileType != "photo" && fileType != "scan" {
		http.Error(w, "Invalid file type", http.StatusBadRequest)
		return
	}

	// Validate file type
	if !utils.ValidateFileType(header.Filename, h.uploadConfig.AllowedTypes) {
		http.Error(w, "Invalid file type", http.StatusBadRequest)
		return
	}

	// Read file content
	fileData, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	// Validate file size
	if int64(len(fileData)) > h.uploadConfig.MaxUploadSize {
		http.Error(w, "File too large", http.StatusBadRequest)
		return
	}

	// Upload file
	filePath, err := h.service.UploadDriverPassport(context.Background(), id, fileData, header.Filename, fileType)
	if err != nil {
		log.Printf("Error uploading driver passport: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := models.FileUploadResponse{
		FilePath: filePath,
		FileName: header.Filename,
		FileSize: int64(len(fileData)),
	}

	respondJSON(w, http.StatusOK, response)
}
// DeleteDriverPhoto handles DELETE /api/drivers/{id}/photo
func (h *DriverHandler) DeleteDriverPhoto(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/drivers/")
	idStr = strings.TrimSuffix(idStr, "/photo")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid driver ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteDriverPhoto(context.Background(), id); err != nil {
		log.Printf("Error deleting driver photo: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "Photo deleted successfully"})
}

// respondJSON sends a JSON response
func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
