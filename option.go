package mifer

import (
	"database/sql"
)

type MiferOption struct {
	ColumnKey string
	Datum     []interface{}
}

func TableNames(db *sql.DB, builder MiferBuilder, dbName string) ([]string, error) {
	var tn = make([]string, 0)
	rows, err := db.Query(`
	SELECT table_name
	FROM information_schema.tables
	WHERE table_schema = 'public'
	ORDER BY table_name;
`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var clm string
		if err := rows.Scan(&clm); err != nil {
			return nil, err
		}
		tn = append(tn, clm)
	}
	return tn, nil
}
