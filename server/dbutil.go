package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func EstablishConnection() *sql.DB {
	credentials := fmt.Sprintf(
		"%s:%s@/%s",
		os.Getenv("MYSQL_USERNAME_CREDENTIAL"),
		os.Getenv("MYSQL_PASSWORD_CREDENTIAL"),
		os.Getenv("MYSQL_DATABASE_CREDENTIAL"),
	)
	db, err := sql.Open(
		"mysql",
		credentials,
	)
	if err != nil {
		panic(err)
	}
	return db
}

func PerformQuery(db *sql.DB, query string, args ...interface{}) *sql.Rows {
	rows, err := db.Query(query, args...)
	if err != nil {
		panic(err)
	}
	return rows
}
