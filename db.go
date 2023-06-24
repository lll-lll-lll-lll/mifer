package mifer

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

const (
	QUERYEND    = ");"
	QUERYPERIOD = ", "
)

var _ MiferDB = (*PostreSQL)(nil)

// key is the name of the column
type Columns = map[string]Column

type MiferDB interface {
	// from database, extract table information and inject into Columns struct
	Scan(ctx context.Context, db *sql.DB, tableName string) (*Column, error)
	// create insert query
	BuildQueries(ctx context.Context, queriesNum int, columns Columns, options ...MiferOption) []string
}

type Column struct {
	// column type. int, nvarchar, text, date...
	ColumnType string
	// ex: PimaryKey, Foreign Key, NOT NULL, UNIQUE,
	// REFERENCES table(column name), CHECK(condition), DEFAULT value.
	Constraint string
}

// PostreSQL represent a table in a database
type PostreSQL struct {
	DBName    string
	TableName string
	Info      PostreSQLInfo
}

func NewTable(dbName string) MiferDB {
	switch dbName {
	case "psql":
		return &PostreSQL{}
	default:
		return nil
	}
}

type PostreSQLInfo struct {
	ColumnName      string `json:"column_name"`
	DateType        string `json:"data_type"`
	OrdinalPosition string `json:"ordinal_position"`
	IsNullable      string `json:"is_nullable"`
}

// TODO: from db, get table information and inject into struct
func (psql *PostreSQL) Scan(ctx context.Context, db *sql.DB, tableName string) (*Column, error) {
	var psti PostreSQLInfo
	if err := db.QueryRowContext(ctx, `
		SELECT column_name,data_type, ordinal_position, is_nullable 
		FROM information_schema.columns
		WHERE table_name = ?
		ORDER BY ordinal_position
		`, tableName).Scan(&psti.ColumnName, &psti.DateType, &psti.OrdinalPosition, &psti.IsNullable); err != nil {
		return nil, Error(DBErr, fmt.Sprintf("failed to prepare table information. check table_name. table_name is %s", tableName))
	}
	return nil, nil
}

func (psql *PostreSQL) BuildQueries(ctx context.Context, queriesNum int, columns Columns, options ...MiferOption) []string {
	if len(options) == 0 {
		return nil
	}
	queries := make([]string, queriesNum)
	clms := joinOptions(options)
	clmsNum := len(options)
	for j := 0; j < clmsNum; j++ {
		queryHolder := queryHolder(columns[options[j].ColumnKey].ColumnType)
		for i := 0; i < len(options[j].Datum); i++ {
			query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (", psql.TableName, clms)
			if queries[i] != "" {
				query = queries[i]
			}
			if clmsNum == j+1 {
				query += fmt.Sprintf(queryHolder+QUERYEND, options[j].Datum[i])
				queries[i] = query
				continue
			}
			query += fmt.Sprintf(queryHolder+QUERYPERIOD, options[j].Datum[i])
			queries[i] = query
		}
	}
	return queries
}

func queryHolder(columnType string) string {
	switch columnType {
	case "character":
		return "'%s'"
	case "nvarchar":
		return "'%s'"
	case "text":
		return "'%s'"
	case "varchar":
		return "'%s'"
	case "char":
		return "'%s'"
	default:
		return "%v"
	}
}

func joinOptions(options []MiferOption) string {
	tmp := make([]string, 0, len(options))
	for _, opt := range options {
		tmp = append(tmp, opt.ColumnKey)
	}
	clms := strings.Join(tmp, ", ")
	return clms
}
