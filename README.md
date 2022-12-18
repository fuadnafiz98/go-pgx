# Golang Postgres with PGX

## Installing

```bash
go get github.com/jackc/pgx/v5/pgxpool # this supports concurrency that is helpful for complex applications
```

- Installing sqlc
  - `brew install sqlc`
  - ``

## The Postgres toolkit

there are several tools to handle the database queries and migrations. Coming from django, it's kinda hard to use the different packages to handle the different functionalities.

for example:

- for generating querires: [sqlc](https://sqlc.dev)
- for migration:
- for dynamic queries: [https://github.com/Masterminds/squirrel](https://github.com/Masterminds/squirrel)
