package mifer

import (
	"context"
	"database/sql"
)

type Mifer struct {
	db *sql.DB
}

type MiferOption struct {
	ColumnKey string
	Datum     []RandomData
}

func (m *Mifer) ExecMigration(ctx context.Context, migration []*MigrationFile) error {
	return nil
}

func (m *Mifer) Inject(ctx context.Context) error {
	return nil
}
