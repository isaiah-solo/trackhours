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

func DBHandlerFor(
	handler func(db *gorm.DB, w http.ResponseWriter, r *http.Request),
) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", BackendOrigin)
		db, err := EstablishConnection()
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
