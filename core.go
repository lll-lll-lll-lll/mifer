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
	Datum     []interface{}
}

func (m *Mifer) ExecMigration(ctx context.Context, migration []*SQLFile) error {
	return nil
}

func (m *Mifer) Inject(ctx context.Context) error {
	return nil
}
