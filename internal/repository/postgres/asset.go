package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"techtrack/internal/domain"
)

type AssetRepository struct {
	db *pgxpool.Pool
}

func NewAssetRepository(db *pgxpool.Pool) *AssetRepository {
	return &AssetRepository{db: db}
}

func (r *AssetRepository) Create(ctx context.Context, asset *domain.Asset) error {
	query := `
		INSERT INTO assets (id, tenant_id, name, qr_code, status, specs, purchase_date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	if asset.CreatedAt.IsZero() {
		asset.CreatedAt = time.Now()
	}
	if asset.UpdatedAt.IsZero() {
		asset.UpdatedAt = time.Now()
	}
	if asset.ID == uuid.Nil {
		asset.ID = uuid.New()
	}

	_, err := r.db.Exec(ctx, query,
		asset.ID,
		asset.TenantID,
		asset.Name,
		asset.QRCode,
		asset.Status,
		asset.Specs,
		asset.PurchaseDate,
		asset.CreatedAt,
		asset.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create asset: %w", err)
	}
	return nil
}

func (r *AssetRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Asset, error) {
	query := `
		SELECT id, tenant_id, name, qr_code, status, specs, purchase_date, created_at, updated_at
		FROM assets
		WHERE id = $1 AND deleted_at IS NULL
	`

	asset := &domain.Asset{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&asset.ID,
		&asset.TenantID,
		&asset.Name,
		&asset.QRCode,
		&asset.Status,
		&asset.Specs,
		&asset.PurchaseDate,
		&asset.CreatedAt,
		&asset.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get asset: %w", err)
	}
	return asset, nil
}
