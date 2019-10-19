package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func EstablishConnection() (*gorm.DB, error) {
	connection := fmt.Sprintf(
		"%s:%s@/%s",
		os.Getenv("MYSQL_USERNAME_CREDENTIAL"),
		os.Getenv("MYSQL_PASSWORD_CREDENTIAL"),
		os.Getenv("MYSQL_DATABASE_CREDENTIAL"),
	)
	db, err := gorm.Open("mysql", connection)
	db.SingularTable(true)
	return db, err
}
