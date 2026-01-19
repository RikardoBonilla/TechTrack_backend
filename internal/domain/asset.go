package domain

import (
	"time"

	"github.com/google/uuid"
)

type AssetStatus string

const (
	AssetStatusActive   AssetStatus = "ACTIVE"
	AssetStatusInRepair AssetStatus = "IN_REPAIR"
	AssetStatusRetired  AssetStatus = "RETIRED"
	AssetStatusLost     AssetStatus = "LOST"
)

type Asset struct {
	ID           uuid.UUID              `json:"id"`
	TenantID     uuid.UUID              `json:"tenant_id"`
	Name         string                 `json:"name"`
	QRCode       string                 `json:"qr_code"`
	Status       AssetStatus            `json:"status"`
	Specs        map[string]interface{} `json:"specs"`
	PurchaseDate *time.Time             `json:"purchase_date,omitempty"`
	CreatedAt    time.Time              `json:"created_at"`
	UpdatedAt    time.Time              `json:"updated_at"`
	DeletedAt    *time.Time             `json:"deleted_at,omitempty"`
}
