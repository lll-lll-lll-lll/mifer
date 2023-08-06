build-psql:
	docker build -t mygo-psql .
stop-psql:
	docker stop  mygo-psql-db
remove-psql:
	docker rm mygo-psql-db && docker image rm mygo-psql
run-psql:
	docker run --name mygo-psql-db -p 5432:5432 -d mygo-psql && docker start mygo-psql-db
login-psql:
	psql -h localhost -U mygo-postgres -d mygo-postgresdb