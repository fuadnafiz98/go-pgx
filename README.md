# Golang Postgres with PGX

## Installing

```bash
go get github.com/jackc/pgx/v5/pgxpool # this supports concurrency that is helpful for complex applications
```

- Installing sqlc
  - `brew install sqlc`

## Generating Bulk data

In this repo I will write some scripts to generate large ammount of data.

### Problem description

1. Querying large ammount of data with joining multiple tables
2. joining multiple tables as well as aggigare data
3. effects before and after indexing
4. searching text, fuzzy searching

Let's consider a relatively complex(!) system where we have to manage a large amount of blogs that publish on different websites, websites hosted on different hosting platforms,

```bash
Blog:
  - ID
  - Title
  - URL
  - Status ---> (Requested, Review, Approved, Integrated, Published, Deleted)
  - Date Published
  - Website
  - Hosting
  - Writer             |
  - Editor             |
  - Publisher          | --> Actors
  - Webmaster          |
  - Hosting Admin      |
```

We will add the tags, category etc later

Head over to the `gen` folder for the codes to generate sql query that will prepare the database.

## The Postgres toolkit

there are several tools to handle the database queries and migrations. Coming from django, it's kinda hard to use the different packages to handle the different functionalities.

for example:

- for generating querires: [sqlc](https://sqlc.dev)
- for migration:
- for dynamic queries: [https://github.com/Masterminds/squirrel](https://github.com/Masterminds/squirrel)
