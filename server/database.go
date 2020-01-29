package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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

func DatabaseWrapper(
	handler func(db *gorm.DB, w http.ResponseWriter, r *http.Request),
) func(http.ResponseWriter, *http.Request) {
	db, err := EstablishConnection()
	return func(w http.ResponseWriter, r *http.Request) {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Response{
				Error: "Issue connection to DB",
			})
			return
		}
		defer db.Close()
		handler(db, w, r)
	}
}
