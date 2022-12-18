package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	db_url := "postgres://fuad:fuad1234@localhost:5432/test-db"
	dbpool, err := pgxpool.New(context.Background(), db_url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect Db: %v\n", err)
		os.Exit(1)
	}

	defer dbpool.Close()

	var hello string
	err = dbpool.QueryRow(context.Background(),
		"SELECT 'HELLO!' ").Scan(&hello)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Query Failed %v\n", err)
		os.Exit(1)
	}

	fmt.Println(hello)
}
