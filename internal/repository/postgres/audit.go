package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"techtrack/internal/domain"
)

type AuditLogRepository struct {
	db *pgxpool.Pool
}

func NewAuditLogRepository(db *pgxpool.Pool) *AuditLogRepository {
	return &AuditLogRepository{db: db}
}

func (r *AuditLogRepository) Create(ctx context.Context, log *domain.AuditLog) error {
	query := `
		INSERT INTO audit_logs (id, tenant_id, actor_id, entity_type, entity_id, action, changes, performed_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	if log.PerformedAt.IsZero() {
		log.PerformedAt = time.Now()
	}
	if log.ID == uuid.Nil {
		log.ID = uuid.New()
	}

	_, err := r.db.Exec(ctx, query,
		log.ID,
		log.TenantID,
		log.ActorID,
		log.EntityType,
		log.EntityID,
		log.Action,
		log.Changes,
		log.PerformedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create audit log: %w", err)
	}
	return nil
}

func (r *AuditLogRepository) GetByTenantID(ctx context.Context, tenantID uuid.UUID) ([]*domain.AuditLog, error) {
	query := `
		SELECT id, tenant_id, actor_id, entity_type, entity_id, action, changes, performed_at
		FROM audit_logs
		WHERE tenant_id = $1
		ORDER BY performed_at DESC
		LIMIT 100
	`

	rows, err := r.db.Query(ctx, query, tenantID)
	if err != nil {
		return nil, fmt.Errorf("failed to list audit logs: %w", err)
	}
	defer rows.Close()

	var logs []*domain.AuditLog
	for rows.Next() {
		log := &domain.AuditLog{}
		err := rows.Scan(
			&log.ID,
			&log.TenantID,
			&log.ActorID,
			&log.EntityType,
			&log.EntityID,
			&log.Action,
			&log.Changes,
			&log.PerformedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan audit log: %w", err)
		}
		logs = append(logs, log)
	}
	return logs, nil
}
