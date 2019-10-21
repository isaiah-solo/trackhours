package main

import (
	"os"
)

var BackendOrigin = os.Getenv("BACKEND_ORIGIN")
var MySQLUsernameCredential = os.Getenv("MYSQL_USERNAME_CREDENTIAL")
var MySQLPasswordCredential = os.Getenv("MYSQL_PASSWORD_CREDENTIAL")
var MySQLDatabaseCredential = os.Getenv("MYSQL_DATABASE_CREDENTIAL")
