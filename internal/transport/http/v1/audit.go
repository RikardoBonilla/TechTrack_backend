package v1

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"techtrack/internal/domain"
)

type AuditHandler struct {
	repo domain.AuditLogRepository
}

func NewAuditHandler(repo domain.AuditLogRepository) *AuditHandler {
	return &AuditHandler{repo: repo}
}

func (h *AuditHandler) RegisterRoutes(r chi.Router) {
	r.Get("/audit-logs/tenant/{tenantID}", h.GetByTenantID)
}

func (h *AuditHandler) GetByTenantID(w http.ResponseWriter, r *http.Request) {
	tenantIDStr := chi.URLParam(r, "tenantID")
	tenantID, err := uuid.Parse(tenantIDStr)
	if err != nil {
		http.Error(w, "invalid tenant id", http.StatusBadRequest)
		return
	}

	logs, err := h.repo.GetByTenantID(r.Context(), tenantID)
	if err != nil {
		http.Error(w, "failed to get audit logs", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(logs); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
