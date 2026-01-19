-- Drop Indexes (Implicitly dropped with tables, but good for completeness if needed individually)
-- DROP INDEX IF EXISTS idx_users_tenant;
-- ... (omitted as DROP TABLE CASCADE handles this)

-- Drop Tables (Order matters due to Foreign Keys)
DROP TABLE IF EXISTS audit_logs;
DROP TABLE IF EXISTS tickets;
DROP TABLE IF EXISTS assets;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS tenants;

-- Drop Enums
DROP TYPE IF EXISTS ticket_status;
DROP TYPE IF EXISTS ticket_priority;
DROP TYPE IF EXISTS asset_status;
DROP TYPE IF EXISTS user_role;

-- Drop Extension
DROP EXTENSION IF EXISTS "pgcrypto";
