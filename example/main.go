package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/lll-lll-lll-lll/mifer"
)

func main() {
	connStr := "postgres://mygo-postgres:mygo-postgres@localhost/mygo-postgresdb?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	tableName := "users"
	c := context.Background()
	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()
	builder := mifer.NewPSQLBuilder(db)
	clmns, err := builder.Scan(ctx, tableName)
	if err != nil {
		log.Fatal(err)
	}
	idOpt := mifer.MiferOption{ColumnKey: "id", Datum: mifer.DefaultMiferGenerator{}.Do(100, mifer.DefaultUUIDPrepareDataCallBack)}
	nameOpt := mifer.MiferOption{ColumnKey: "name", Datum: mifer.DefaultMiferGenerator{}.Do(100, mifer.DefaultStringPrepareDataCallBack)}
	queries, err := builder.BuildQueries(ctx, clmns, tableName, idOpt, nameOpt)
	if err != nil {
		log.Fatal(err)
	}
	if err := mifer.Inject(ctx, db, queries); err != nil {
		log.Fatal(err)
	}
}
