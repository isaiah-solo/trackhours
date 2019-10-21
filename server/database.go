package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func EstablishConnection() (*gorm.DB, error) {
	connection := fmt.Sprintf(
		"%s:%s@/%s",
		MySQLUsernameCredential,
		MySQLPasswordCredential,
		MySQLDatabaseCredential,
	)
	db, err := gorm.Open("mysql", connection)
	db.SingularTable(true)
	return db, err
}
