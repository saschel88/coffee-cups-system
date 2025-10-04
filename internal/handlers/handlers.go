package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/your-username/coffee-cups-system/internal/models"
	"github.com/your-username/coffee-cups-system/internal/services"
)

// Handlers holds all HTTP handlers
type Handlers struct {
	services *services.Services
	logger   interface{}
}

// New creates a new Handlers instance
func New(services *services.Services, logger interface{}) *Handlers {
	return &Handlers{
		services: services,
		logger:   logger,
	}
}

// GetUsers handles GET /api/v1/users
func (h *Handlers) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.services.User.GetAllActiveUsers()
	if err != nil {
		http.Error(w, "Failed to get users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetUser handles GET /api/v1/users/{id}
func (h *Handlers) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.services.User.GetUserByTelegramID(int64(id))
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// GetBoxes handles GET /api/v1/boxes
func (h *Handlers) GetBoxes(w http.ResponseWriter, r *http.Request) {
	boxes, err := h.services.Box.GetActiveBoxes()
	if err != nil {
		http.Error(w, "Failed to get boxes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(boxes)
}

// CreateBox handles POST /api/v1/boxes
func (h *Handlers) CreateBox(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name      string  `json:"name"`
		TotalCups int     `json:"total_cups"`
		Price     float64 `json:"price"`
		CreatedBy uint    `json:"created_by"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	box, err := h.services.Box.CreateBox(req.Name, req.TotalCups, req.Price, req.CreatedBy)
	if err != nil {
		http.Error(w, "Failed to create box", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(box)
}

// GetBox handles GET /api/v1/boxes/{id}
func (h *Handlers) GetBox(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid box ID", http.StatusBadRequest)
		return
	}

	box, err := h.services.Box.GetBoxByID(uint(id))
	if err != nil {
		http.Error(w, "Box not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(box)
}

// GetCoffeeLogs handles GET /api/v1/coffee-logs
func (h *Handlers) GetCoffeeLogs(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
		http.Error(w, "user_id parameter is required", http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	logs, err := h.services.Coffee.GetUserCoffeeLogs(uint(userID), 0)
	if err != nil {
		http.Error(w, "Failed to get coffee logs", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logs)
}

// LogCoffee handles POST /api/v1/coffee-logs
func (h *Handlers) LogCoffee(w http.ResponseWriter, r *http.Request) {
	var req struct {
		UserID uint `json:"user_id"`
		BoxID  uint `json:"box_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	log, err := h.services.Coffee.LogCoffee(req.UserID, req.BoxID)
	if err != nil {
		http.Error(w, "Failed to log coffee", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(log)
}

// GetPayments handles GET /api/v1/payments
func (h *Handlers) GetPayments(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	boxIDStr := r.URL.Query().Get("box_id")

	var payments []models.Payment
	var err error

	if userIDStr != "" {
		userID, parseErr := strconv.ParseUint(userIDStr, 10, 32)
		if parseErr != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}
		payments, err = h.services.Payment.GetUserPayments(uint(userID))
	} else if boxIDStr != "" {
		boxID, parseErr := strconv.ParseUint(boxIDStr, 10, 32)
		if parseErr != nil {
			http.Error(w, "Invalid box ID", http.StatusBadRequest)
			return
		}
		payments, err = h.services.Payment.GetBoxPayments(uint(boxID))
	} else {
		http.Error(w, "user_id or box_id parameter is required", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, "Failed to get payments", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(payments)
}
