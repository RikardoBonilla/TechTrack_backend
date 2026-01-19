package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"techtrack/internal/domain"
)

type TenantRepository struct {
	db *pgxpool.Pool
}

func NewTenantRepository(db *pgxpool.Pool) *TenantRepository {
	return &TenantRepository{db: db}
}

func (r *TenantRepository) Create(ctx context.Context, tenant *domain.Tenant) error {
	query := `
		INSERT INTO tenants (id, name, subscription_plan, is_active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	if tenant.CreatedAt.IsZero() {
		tenant.CreatedAt = time.Now()
	}
	if tenant.UpdatedAt.IsZero() {
		tenant.UpdatedAt = time.Now()
	}
	if tenant.ID == uuid.Nil {
		tenant.ID = uuid.New()
	}

	_, err := r.db.Exec(ctx, query,
		tenant.ID,
		tenant.Name,
		tenant.SubscriptionPlan,
		tenant.IsActive,
		tenant.CreatedAt,
		tenant.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create tenant: %w", err)
	}
	return nil
}

func (r *TenantRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Tenant, error) {
	query := `
		SELECT id, name, subscription_plan, is_active, created_at, updated_at
		FROM tenants
		WHERE id = $1
	`
	tenant := &domain.Tenant{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&tenant.ID,
		&tenant.Name,
		&tenant.SubscriptionPlan,
		&tenant.IsActive,
		&tenant.CreatedAt,
		&tenant.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get tenant: %w", err)
	}
	return tenant, nil
}
