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
	c := context.Background()
	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()
	builder := mifer.NewPSQLBuilder(db)
	clmns, err := builder.Scan(ctx, "users")
	if err != nil {
		log.Fatal(err)
	}
	randomData := mifer.DefaultMiferGenerator{}.Do(100, mifer.DefaultUUIDPrepareDataCallBack)
	idOpt := mifer.MiferOption{ColumnKey: "id", Datum: randomData}
	queries, err := builder.BuildQueries(ctx, clmns, "users", idOpt)
	if err != nil {
		log.Fatal(err)
	}
	if err := mifer.Inject(ctx, db, queries); err != nil {
		log.Fatal(err)
	}

}
