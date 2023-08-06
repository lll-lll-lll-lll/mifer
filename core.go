package mifer

import (
	"context"
	"database/sql"
)

type Mifer struct {
	db *sql.DB
}

func (m *Mifer) Inject(ctx context.Context, queries []string) error {
	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	for _, query := range queries {
		if _, err := tx.ExecContext(ctx, query); err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
