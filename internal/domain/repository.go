package domain

import (
	"context"

	"github.com/google/uuid"
)

type TenantRepository interface {
	Create(ctx context.Context, tenant *Tenant) error
	GetByID(ctx context.Context, id uuid.UUID) (*Tenant, error)
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
}

type AssetRepository interface {
	Create(ctx context.Context, asset *Asset) error
	GetByID(ctx context.Context, id uuid.UUID) (*Asset, error)
}

type TicketRepository interface {
	Create(ctx context.Context, ticket *Ticket) error
	GetByID(ctx context.Context, id uuid.UUID) (*Ticket, error)
	// Additional filters will be needed later
}

type AuditLogRepository interface {
	Create(ctx context.Context, log *AuditLog) error
	GetByTenantID(ctx context.Context, tenantID uuid.UUID) ([]*AuditLog, error)
}
