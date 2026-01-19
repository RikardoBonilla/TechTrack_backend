package v1

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"techtrack/internal/domain"
)

type AssetHandler struct {
	repo domain.AssetRepository
}

func NewAssetHandler(repo domain.AssetRepository) *AssetHandler {
	return &AssetHandler{repo: repo}
}

func (h *AssetHandler) RegisterRoutes(r chi.Router) {
	r.Post("/assets", h.Create)
	r.Get("/assets/{id}", h.GetByID)
}

func (h *AssetHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req domain.Asset
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.repo.Create(r.Context(), &req); err != nil {
		slog.Error("failed to create asset", "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(req); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func (h *AssetHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid asset id", http.StatusBadRequest)
		return
	}

	asset, err := h.repo.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "asset not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(asset); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
