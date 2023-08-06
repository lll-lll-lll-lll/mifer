package mifer

import (
	"context"
	"database/sql"
)

const (
	QUERYEND    = ");"
	QUERYPERIOD = ", "
)

type MiferBuilder interface {
	// from database, extract table information and mapping scanned data into `Columns` type
	Scan(ctx context.Context, db *sql.DB, tableName string) (*Column, error)
	// create insert query
	BuildQueries(ctx context.Context, maxDatumNum int, columns Columns, options ...MiferOption) ([]string, error)
}

type Column struct {
	// ex: int, nvarchar, text, date...
	Type string
	// ex: pimary key, foreign key, not null, unique,
	// references table(column name), check(condition), default value.
	Constraint string
}

// key is the name of the column
type Columns = map[string]Column
