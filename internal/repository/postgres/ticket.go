package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"techtrack/internal/domain"
)

type TicketRepository struct {
	db *pgxpool.Pool
}

func NewTicketRepository(db *pgxpool.Pool) *TicketRepository {
	return &TicketRepository{db: db}
}

func (r *TicketRepository) Create(ctx context.Context, ticket *domain.Ticket) error {
	query := `
		INSERT INTO tickets (id, tenant_id, asset_id, reporter_id, assigned_to_id, title, description, priority, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`
	if ticket.CreatedAt.IsZero() {
		ticket.CreatedAt = time.Now()
	}
	if ticket.UpdatedAt.IsZero() {
		ticket.UpdatedAt = time.Now()
	}
	if ticket.ID == uuid.Nil {
		ticket.ID = uuid.New()
	}

	_, err := r.db.Exec(ctx, query,
		ticket.ID,
		ticket.TenantID,
		ticket.AssetID,
		ticket.ReporterID,
		ticket.AssignedToID,
		ticket.Title,
		ticket.Description,
		ticket.Priority,
		ticket.Status,
		ticket.CreatedAt,
		ticket.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create ticket: %w", err)
	}
	return nil
}

func (r *TicketRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Ticket, error) {
	query := `
		SELECT id, tenant_id, asset_id, reporter_id, assigned_to_id, title, description, priority, status, created_at, updated_at, closed_at
		FROM tickets
		WHERE id = $1
	`

	ticket := &domain.Ticket{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&ticket.ID,
		&ticket.TenantID,
		&ticket.AssetID,
		&ticket.ReporterID,
		&ticket.AssignedToID,
		&ticket.Title,
		&ticket.Description,
		&ticket.Priority,
		&ticket.Status,
		&ticket.CreatedAt,
		&ticket.UpdatedAt,
		&ticket.ClosedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get ticket: %w", err)
	}
	return ticket, nil
}
