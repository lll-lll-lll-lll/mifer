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
	// from database, extract table information and mapping scanned data into `Columns` type
	Scan(ctx context.Context, db *sql.DB, tableName string) (*Column, error)
	// create insert query
	BuildQueries(ctx context.Context, maxDatumNum int, columns Columns, options ...MiferOption) ([]string, error)
}

type Column struct {
	// ex: int, nvarchar, text, date...
	ColumnType string
	// ex: pimary key, foreign key, not null, unique,
	// references table(column name), check(condition), default value.
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
	TableName      string `json:"table_name"`
	ColumnName     string `json:"column_name"`
	DataType       string `json:"data_type"`
	IsNullable     string `json:"is_nullable"`
	ColumnDefault  string `json:"column_default"`
	ConstraintName string `json:"constraint_name"`
	ConstraintType string `json:"constraint_type"`
}

// TODO: from db, get table information and inject into struct
func (psql *PostreSQL) Scan(ctx context.Context, db *sql.DB, tableName string) (*Column, error) {
	var pstis []PostreSQLInfo
	rows, err := db.QueryContext(ctx, `
		SELECT
		table_name,
		column_name,
		data_type,
		is_nullable,
		column_default,
		constraint_name,
		constraint_type
	FROM
		information_schema.columns c
		LEFT JOIN information_schema.constraint_column_usage ccu ON c.table_schema = ccu.table_schema
			AND c.table_name = ccu.table_name
			AND c.column_name = ccu.column_name
		LEFT JOIN information_schema.table_constraints tc ON ccu.constraint_schema = tc.constraint_schema
			AND ccu.constraint_name = tc.constraint_name
	WHERE
		c.table_schema = 'public'
	ORDER BY
		table_name,
		ordinal_position;
		`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var psti PostreSQLInfo
		if err := rows.Scan(&psti.TableName, &psti.ColumnName, &psti.DataType,
			&psti.IsNullable, &psti.ColumnDefault, &psti.ConstraintName, &psti.ConstraintType); err != nil {
			return nil, err
		}
		pstis = append(pstis, psti)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// TODO
	return nil, nil
}

func (psql *PostreSQL) BuildQueries(ctx context.Context, maxDataNum int, columns Columns, options ...MiferOption) ([]string, error) {
	if len(options) == 0 {
		return nil, Error(NoOptionsErr, "Not a option was provided. At least one option must be provided")
	}
	queries := make([]string, maxDataNum)
	columnNames := joinColumnWithComma(options)
	columnNum := len(options)

	for j := 0; j < columnNum; j++ {
		columnDataNum := len(options[j].Datum)
		if isGrater := isGraterThanColumnDataNum(maxDataNum, columnDataNum); !isGrater {
			return nil, Error(NoTypeErr, fmt.Sprintf("maxDatumNum must be greater than RandomDataNum of option. maxDatumNum is %d. columnDataNum is %d. column name is %v", maxDataNum, columnDataNum, options[j].ColumnKey))
		}
		dataFormat := switchFormatByColumnType(columns[options[j].ColumnKey].ColumnType)

		for i := 0; i < columnDataNum; i++ {
			query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (", psql.TableName, columnNames)
			if queries[i] != "" {
				query = queries[i]
			}
			if columnNum == j+1 {
				query += fmt.Sprintf(dataFormat+QUERYEND, options[j].Datum[i])
				queries[i] = query
				continue
			}
			query += fmt.Sprintf(dataFormat+QUERYPERIOD, options[j].Datum[i])
			queries[i] = query
		}
	}
	return queries, nil
}

func isGraterThanColumnDataNum(maxDatumNum, columnRandomDataNum int) bool {
	if maxDatumNum >= columnRandomDataNum {
		return true
	}
	return false
}

func switchFormatByColumnType(columnType string) string {
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

func joinColumnWithComma(options []MiferOption) string {
	tmp := make([]string, 0, len(options))
	for _, opt := range options {
		tmp = append(tmp, opt.ColumnKey)
	}
	clms := strings.Join(tmp, ", ")
	return clms
}
