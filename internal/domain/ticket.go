package domain

import (
	"time"

	"github.com/google/uuid"
)

type TicketPriority string

const (
	TicketPriorityLow      TicketPriority = "LOW"
	TicketPriorityMedium   TicketPriority = "MEDIUM"
	TicketPriorityHigh     TicketPriority = "HIGH"
	TicketPriorityCritical TicketPriority = "CRITICAL"
)

type TicketStatus string

const (
	TicketStatusOpen       TicketStatus = "OPEN"
	TicketStatusInProgress TicketStatus = "IN_PROGRESS"
	TicketStatusBlocked    TicketStatus = "BLOCKED"
	TicketStatusResolved   TicketStatus = "RESOLVED"
)

type Ticket struct {
	ID           uuid.UUID      `json:"id"`
	TenantID     uuid.UUID      `json:"tenant_id"`
	AssetID      uuid.UUID      `json:"asset_id"`
	ReporterID   *uuid.UUID     `json:"reporter_id,omitempty"`
	AssignedToID *uuid.UUID     `json:"assigned_to_id,omitempty"`
	Title        string         `json:"title"`
	Description  string         `json:"description"`
	Priority     TicketPriority `json:"priority"`
	Status       TicketStatus   `json:"status"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	ClosedAt     *time.Time     `json:"closed_at,omitempty"`
}
