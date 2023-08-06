FROM postgres:15.3-bullseye

COPY ./init/init.sql /docker-entrypoint-initdb.d/init.sql

ENV POSTGRES_USER mygo-postgres
ENV POSTGRES_PASSWORD mygo-postgres
ENV POSTGRES_DB mygo-postgresdb

EXPOSE 5432
