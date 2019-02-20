package main

import (
  "fmt"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

func EstablishConnection() *sql.DB {
  db, err := sql.Open(
    "mysql",
    fmt.Sprintf(
      "%s:%s@/%s",
      MySQLUsernameCredential,
      MySQLPasswordCredential,
      MySQLDatabaseCredential,
    ),
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
