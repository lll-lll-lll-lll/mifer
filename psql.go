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

	ScanQuery = `
	SELECT column_name, data_type
	FROM information_schema.columns
	WHERE table_name = '%s';`
)

func NewPSQLBuilder(db *sql.DB) *PostresBuilder {
	return &PostresBuilder{db: db}
}

// PostresBuilder represent a table in a database
type PostresBuilder struct {
	DBName  string
	columns *Columns
	db      *sql.DB
}

// Scan from database, extract table information and mapping scanned data into `Columns` type
func (psql *PostresBuilder) Scan(ctx context.Context, tableName string) (Columns, error) {
	clms := Columns{}
	rows, err := psql.db.QueryContext(ctx, fmt.Sprintf(ScanQuery, tableName))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var columnName string
		var column Column
		if err := rows.Scan(columnName, &column.Type); err != nil {
			return nil, err
		}
		clms[columnName] = column
	}

	return clms, nil
}

// BuildQueries create insert queries. from options information.
// return error if options's num is zero.
func (psql *PostresBuilder) BuildQueries(ctx context.Context, columns Columns, tableName string, options ...MiferOption) ([]string, error) {
	if len(options) == 0 {
		return nil, NewErr(NoOptionsErr, "Not a option was provided. At least one option must be provided")
	}
	queryNum := MaxOptDatum(options...)
	queries := make([]string, queryNum)
	columnNames := joinClmnKeys(options)
	columnNum := len(options)

	for j := 0; j < columnNum; j++ {
		columnDataNum := len(options[j].Datum)
		_, ok := columns[options[j].ColumnKey]
		if !ok {
			return nil, NewErr(NoTypeErr, "no column key specified")
		}
		dataFormat := checkType(columns[options[j].ColumnKey].Type)

		buildEachQuery(ctx, columnNum, columnDataNum, tableName, columnNames, dataFormat, &options[j], j+1, queries)
	}

	return queries, nil
}

func buildEachQuery(ctx context.Context, columnNum int, columnDataNum int, tableName string, columnNames string, dataFormat string, option *MiferOption, endIdx int, queries []string) []string {
	for i := 0; i < columnDataNum; i++ {
		q := fmt.Sprintf("INSERT INTO %s (%s) VALUES (", tableName, columnNames)
		if queries[i] != "" {
			q = queries[i]
		}

		if columnNum == endIdx {
			q += fmt.Sprintf(dataFormat+QUERYEND, option.Datum[i])
			queries[i] = q
			continue
		}

		q += fmt.Sprintf(dataFormat+QUERYPERIOD, option.Datum[i])
		queries[i] = q
	}
	return queries
}

func joinClmnKeys(options []MiferOption) string {
	tmp := make([]string, 0, len(options))
	for _, opt := range options {
		tmp = append(tmp, opt.ColumnKey)
	}
	clms := strings.Join(tmp, ", ")
	return clms
}

// checkType switch data format by column type for query's placeholder.
func checkType(columnType string) string {
	switch columnType {
	case "text":
		return "'%s'"
	case "nvarchar":
		return "'%s'"
	case "varchar":
		return "'%s'"
	default:
		return "%v"
	}
}
