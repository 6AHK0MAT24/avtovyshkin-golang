package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"clients-service/internal/models"
	"clients-service/internal/service"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// ClientHandler handles HTTP requests for clients
type ClientHandler struct {
	service *service.ClientService
}

// NewClientHandler creates a new client handler
func NewClientHandler(service *service.ClientService) *ClientHandler {
	return &ClientHandler{service: service}
}

// CreateClient handles POST /api/clients
// @Summary Создать нового клиента
// @Description Создает нового клиента в системе
// @Tags Clients
// @Accept json
// @Produce json
// @Param client body models.CreateClientRequest true "Данные клиента"
// @Success 201 {object} models.Client
// @Failure 400 {object} map[string]string
// @Router /clients [post]
func (h *ClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	var req models.CreateClientRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	client, err := h.service.Create(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(client)
}

// GetClients handles GET /api/clients
// @Summary Получить список клиентов
// @Description Возвращает список всех клиентов с пагинацией
// @Tags Clients
// @Produce json
// @Param page query int false "Номер страницы" default(1)
// @Param perPage query int false "Количество элементов на странице" default(10)
// @Success 200 {object} models.ClientListResponse
// @Failure 500 {object} map[string]string
// @Router /clients [get]
func (h *ClientHandler) GetClients(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	perPage, _ := strconv.Atoi(r.URL.Query().Get("perPage"))

	if page < 1 {
		page = 1
	}
	if perPage < 1 || perPage > 100 {
		perPage = 10
	}

	response, err := h.service.GetAll(r.Context(), page, perPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetClientsByType handles GET /api/clients/type/{type}
// @Summary Получить клиентов по типу
// @Description Возвращает список клиентов определенного типа с пагинацией
// @Tags Clients
// @Produce json
// @Param type path string true "Тип клиента" Enums(individual, legal)
// @Param page query int false "Номер страницы" default(1)
// @Param perPage query int false "Количество элементов на странице" default(10)
// @Success 200 {object} models.ClientListResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /clients/type/{type} [get]
func (h *ClientHandler) GetClientsByType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clientTypeStr := vars["type"]

	var clientType models.ClientType
	switch clientTypeStr {
	case "individual":
		clientType = models.ClientTypeIndividual
	case "legal":
		clientType = models.ClientTypeLegalEntity
	default:
		http.Error(w, "Invalid client type", http.StatusBadRequest)
		return
	}

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	perPage, _ := strconv.Atoi(r.URL.Query().Get("perPage"))

	if page < 1 {
		page = 1
	}
	if perPage < 1 || perPage > 100 {
		perPage = 10
	}

	response, err := h.service.GetByType(r.Context(), clientType, page, perPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetClient handles GET /api/clients/{id}
// @Summary Получить клиента по ID
// @Description Возвращает информацию о клиенте по его идентификатору
// @Tags Clients
// @Produce json
// @Param id path string true "ID клиента"
// @Success 200 {object} models.Client
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /clients/{id} [get]
func (h *ClientHandler) GetClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid client ID", http.StatusBadRequest)
		return
	}

	client, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		if err.Error() == "client not found" {
			http.Error(w, "Client not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(client)
}

// UpdateClient handles PUT /api/clients/{id}
// @Summary Обновить клиента
// @Description Обновляет информацию о существующем клиенте
// @Tags Clients
// @Accept json
// @Produce json
// @Param id path string true "ID клиента"
// @Param client body models.UpdateClientRequest true "Данные для обновления"
// @Success 200 {object} models.Client
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /clients/{id} [put]
func (h *ClientHandler) UpdateClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid client ID", http.StatusBadRequest)
		return
	}

	var req models.UpdateClientRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	client, err := h.service.Update(r.Context(), id, &req)
	if err != nil {
		if err.Error() == "client not found" {
			http.Error(w, "Client not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(client)
}

// DeleteClient handles DELETE /api/clients/{id}
// @Summary Удалить клиента
// @Description Удаляет клиента из системы
// @Tags Clients
// @Produce json
// @Param id path string true "ID клиента"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /clients/{id} [delete]
func (h *ClientHandler) DeleteClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid client ID", http.StatusBadRequest)
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		if err.Error() == "client not found" {
			http.Error(w, "Client not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// SearchClients handles GET /api/clients/search
// @Summary Поиск клиентов
// @Description Ищет клиентов по заданным фильтрам
// @Tags Clients
// @Produce json
// @Param query query string false "Поисковый запрос"
// @Param clientType query string false "Тип клиента" Enums(individual, legal_entity)
// @Param status query string false "Статус клиента" Enums(active, inactive, blocked)
// @Param page query int false "Номер страницы" default(1)
// @Param perPage query int false "Количество элементов на странице" default(10)
// @Success 200 {object} models.ClientListResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /clients/search [get]
func (h *ClientHandler) SearchClients(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	clientTypeStr := r.URL.Query().Get("clientType")
	statusStr := r.URL.Query().Get("status")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	perPage, _ := strconv.Atoi(r.URL.Query().Get("perPage"))

	if page < 1 {
		page = 1
	}
	if perPage < 1 || perPage > 100 {
		perPage = 10
	}

	filters := &models.SearchFilters{
		Page:    page,
		PerPage: perPage,
	}

	if query != "" {
		filters.Query = &query
	}

	if clientTypeStr != "" {
		var clientType models.ClientType
		switch clientTypeStr {
		case "individual":
			clientType = models.ClientTypeIndividual
		case "legal_entity":
			clientType = models.ClientTypeLegalEntity
		default:
			http.Error(w, "Invalid client type", http.StatusBadRequest)
			return
		}
		filters.ClientType = &clientType
	}

	if statusStr != "" {
		var status models.ClientStatus
		switch statusStr {
		case "active":
			status = models.ClientStatusActive
		case "inactive":
			status = models.ClientStatusInactive
		case "blocked":
			status = models.ClientStatusBlocked
		default:
			http.Error(w, "Invalid status", http.StatusBadRequest)
			return
		}
		filters.Status = &status
	}

	response, err := h.service.Search(r.Context(), filters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
