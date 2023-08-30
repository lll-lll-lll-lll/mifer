package mifer

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
)

type Column struct {
	// ex: int, nvarchar, text, date...
	Type string
	// ex: pimary key, foreign key, not null, unique,
	// references table(column name), check(condition), default value.
	Constraint string
}

// key is the name of the column
type Columns = map[string]Column

type MiferBuilder interface {
	// from database, extract table information and mapping scanned data into `Columns` type
	Scan(ctx context.Context, tableName string) (*Column, error)
	// create insert query
	BuildQueries(ctx context.Context, columns Columns, tableName string, options ...MiferOption) ([]string, error)
}

func Inject(ctx context.Context, db *sql.DB, queries []string) error {
	tx, err := db.BeginTx(ctx, nil)
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

func Output(queries []string, dirPath, fileName string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return fmt.Errorf("%v is not a directory path", dirPath)
	}
	f, err := os.Create(filepath.Join(dirPath, fileName))
	if err != nil {
		return err
	}
	for _, query := range queries {
		f.WriteString(query + "\n")
	}
	return nil
}

func maxOptDatum(opts ...MiferOption) int {
	v := -1

	for _, opt := range opts {
		num := len(opt.Datum)
		if v <= num {
			v = num
			continue
		}
	}

	return v
}
