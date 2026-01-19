package v1

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"techtrack/internal/domain"
)

type TenantHandler struct {
	repo domain.TenantRepository
}

func NewTenantHandler(repo domain.TenantRepository) *TenantHandler {
	return &TenantHandler{repo: repo}
}

func (h *TenantHandler) RegisterRoutes(r chi.Router) {
	r.Post("/tenants", h.Create)
	r.Get("/tenants/{id}", h.GetByID)
}

func (h *TenantHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req domain.Tenant
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.repo.Create(r.Context(), &req); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(req); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func (h *TenantHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid tenant id", http.StatusBadRequest)
		return
	}

	tenant, err := h.repo.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "tenant not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tenant); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
