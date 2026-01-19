package v1

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"techtrack/internal/domain"
)

type TicketHandler struct {
	repo domain.TicketRepository
}

func NewTicketHandler(repo domain.TicketRepository) *TicketHandler {
	return &TicketHandler{repo: repo}
}

func (h *TicketHandler) RegisterRoutes(r chi.Router) {
	r.Post("/tickets", h.Create)
	r.Get("/tickets/{id}", h.GetByID)
}

func (h *TicketHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req domain.Ticket
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.repo.Create(r.Context(), &req); err != nil {
		slog.Error("failed to create ticket", "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(req); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func (h *TicketHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid ticket id", http.StatusBadRequest)
		return
	}

	ticket, err := h.repo.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "ticket not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(ticket); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
