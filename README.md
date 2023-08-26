![](./public/icon.png)

# Overview

mifer is a library which generate an 'INSERT' QUERY by the random data from a
table info

# Usage

```sh
go get github.com/lll-lll-lll-lll/mifer
```

```go
func main(){
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
	nameOpt := mifer.MiferOption{ColumnKey: "name", Datum: mifer.DefaultMiferGenerator{}.Do(100, mifer.DefaultStringPrepareDataCallBack)}
	queries, err := builder.BuildQueries(ctx, clmns, tableName, nameOpt)
	if err != nil {
		log.Fatal(err)
	}
	if err := mifer.Inject(ctx, db, queries); err != nil {
		log.Fatal(err)
	}
}
```
