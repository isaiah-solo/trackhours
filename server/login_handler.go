package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type AccountInformation struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type LoginResponse struct {
	Error string `json:"error"`
}

type User struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

var loginInternalErrorResponse = LoginResponse{
	Error: "Issue with server rendering",
}

var loginInvalidUserResponse = LoginResponse{
	Error: "User does not exist",
}

var loginIncorrectPasswordResponse = LoginResponse{
	Error: "Incorrect username and password combination",
}

var successResponse = LoginResponse{
	Error: "",
}

func AccountCreationHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "http://trackhours.co")
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	db := EstablishConnection()
	defer db.Close()
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &loginInternalErrorResponse)
		return
	}
	var accountInformation AccountInformation
	if err := json.Unmarshal(body, &accountInformation); err != nil {
		c.JSON(http.StatusInternalServerError, &loginInternalErrorResponse)
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(accountInformation.Password), bcrypt.MinCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &loginInternalErrorResponse)
		return
	}
	user := User{
		Password: string(hash),
		Username: accountInformation.Username,
	}
	userInsert, err := db.Prepare(
		"INSERT INTO user (password, username) VALUES (?, ?)",
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &loginInternalErrorResponse)
		return
	}
	if _, err := userInsert.Exec(
		user.Password,
		user.Username,
	); err != nil {
		c.JSON(http.StatusInternalServerError, &loginInternalErrorResponse)
		return
	}
	// Generate user session key
	u, err := uuid.NewV4()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &loginInternalErrorResponse)
		return
	}
	sessionKey := u.String()
	sessionKeyInsert, err := db.Prepare(
		"INSERT INTO user_session (owner_username, session_key) VALUES (?, ?)",
	)
	if _, err := sessionKeyInsert.Exec(user.Username, sessionKey); err != nil {
		c.JSON(http.StatusInternalServerError, &loginInternalErrorResponse)
		return
	}
	// Set cookie
	c.SetCookie("trackhours_session_key", sessionKey, 360000, "/", "", false, false)
	c.JSON(http.StatusOK, &successResponse)
}

func CheckLoginHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "http://trackhours.co")
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	cookie, err := c.Cookie("trackhours_session_key")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"is_logged_in": false, "error": err})
		return
	}
	db := EstablishConnection()
	defer db.Close()
	rows, err := db.Query("SELECT session_key FROM user_session where session_key = ?", cookie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &loginInternalErrorResponse)
		return
	}
	if rows.Next() != true {
		c.JSON(http.StatusOK, gin.H{"is_logged_in": false, "error": nil})
		return
	}
	var sessionKey string
	if err := rows.Scan(
		&sessionKey,
	); err != nil {
		c.JSON(http.StatusInternalServerError, &loginInternalErrorResponse)
		return
	}
	c.JSON(http.StatusOK, gin.H{"is_logged_in": cookie == sessionKey, "error": nil})
}

func LoginHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "http://trackhours.co")
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	db := EstablishConnection()
	defer db.Close()
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &loginInternalErrorResponse)
		return
	}
	var accountInformation AccountInformation
	if err := json.Unmarshal(body, &accountInformation); err != nil {
		c.JSON(http.StatusInternalServerError, &loginInternalErrorResponse)
		return
	}
	rows := PerformQuery(
		db,
		"SELECT password, username FROM user WHERE username = ?",
		accountInformation.Username,
	)
	if rows.Next() != true {
		c.JSON(http.StatusUnauthorized, &loginInvalidUserResponse)
		return
	}
	var user User
	if err := rows.Scan(
		&user.Password,
		&user.Username,
	); err != nil {
		c.JSON(http.StatusInternalServerError, &loginInternalErrorResponse)
		return
	}
	// Compare credentials with those stored in DB
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(accountInformation.Password),
	); err != nil {
		c.JSON(http.StatusUnauthorized, &loginIncorrectPasswordResponse)
		return
	}
	// Generate user session key
	u, err := uuid.NewV4()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &loginInternalErrorResponse)
		return
	}
	sessionKey := u.String()
	sessionKeyInsert, err := db.Prepare(
		"INSERT INTO user_session (owner_username, session_key) VALUES (?, ?)",
	)
	if _, err := sessionKeyInsert.Exec(user.Username, sessionKey); err != nil {
		c.JSON(http.StatusInternalServerError, &loginInternalErrorResponse)
		return
	}
	c.SetCookie("trackhours_session_key", sessionKey, 360000, "/", "", false, false)
	c.JSON(http.StatusOK, &successResponse)
}

func LogoutHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "http://trackhours.co")
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	c.SetCookie("trackhours_session_key", "", 360000, "/", "", false, false)
	c.JSON(http.StatusOK, &successResponse)
}
