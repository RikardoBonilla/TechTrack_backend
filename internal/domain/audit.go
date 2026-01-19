package domain

import (
	"time"

	"github.com/google/uuid"
)

type AuditLog struct {
	ID          uuid.UUID              `json:"id"`
	TenantID    uuid.UUID              `json:"tenant_id"`
	ActorID     *uuid.UUID             `json:"actor_id,omitempty"`
	EntityType  string                 `json:"entity_type"`
	EntityID    uuid.UUID              `json:"entity_id"`
	Action      string                 `json:"action"`
	Changes     map[string]interface{} `json:"changes,omitempty"`
	PerformedAt time.Time              `json:"performed_at"`
}
