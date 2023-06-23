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

var _ Table = (*PostreSQL)(nil)

// key is the name of the column
type Columns = map[string]Column

type Table interface {
	// column names in the table
	Columns() (Columns, error)
	// from database, extract table information and inject into Columns struct
	Scan(db *sql.DB) error
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
	DBName      string
	TableName   string
	ColumnsInfo Columns
}

func NewTable(dbName string) Table {
	switch dbName {
	case "psql":
		return &PostreSQL{}
	default:
		return nil
	}
}

func (psql *PostreSQL) Columns() (Columns, error) {
	if psql.ColumnsInfo == nil {
		return nil, Error(DBErr, "No column information. calling 'Scan' method first")
	}
	return psql.ColumnsInfo, nil
}

// TODO: dbからテーブル情報を取得し、構造体に落とし込む
func (psql *PostreSQL) Scan(db *sql.DB) error {
	return nil
}

func (psql *PostreSQL) BuildQueries(ctx context.Context, queriesNum int, columns Columns, options ...MiferOption) []string {
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

func ReadRandomDataQuery(filePath string) string {
	return ""
}
