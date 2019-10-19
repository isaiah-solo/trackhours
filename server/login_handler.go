package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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

type UserSession struct {
	OwnerUsername string `json:"owner_username"`
	SessionKey    string `json:"session_key"`
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
	connection := fmt.Sprintf(
		"%s:%s@/%s",
		os.Getenv("MYSQL_USERNAME_CREDENTIAL"),
		os.Getenv("MYSQL_PASSWORD_CREDENTIAL"),
		os.Getenv("MYSQL_DATABASE_CREDENTIAL"),
	)
	db, err := gorm.Open("mysql", connection)
	db.SingularTable(true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &loginInternalErrorResponse)
		return
	}
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
	db.Create(&user)
	// Generate user session key
	u, err := uuid.NewV4()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &loginInternalErrorResponse)
		return
	}
	sessionKey := u.String()
	userSession := UserSession{
		OwnerUsername: user.Username,
		SessionKey:    sessionKey,
	}
	db.Create(&userSession)
	// Set cookie
	c.SetCookie("trackhours_session_key", sessionKey, 360000, "/", "", false, false)
	c.JSON(http.StatusOK, &successResponse)
}

func CheckLoginHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "http://trackhours.co")
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	cookie, err := c.Cookie("trackhours_session_key")
	if err != nil || cookie == "" {
		c.JSON(http.StatusOK, gin.H{"is_logged_in": false, "error": err})
		return
	}
	connection := fmt.Sprintf(
		"%s:%s@/%s",
		os.Getenv("MYSQL_USERNAME_CREDENTIAL"),
		os.Getenv("MYSQL_PASSWORD_CREDENTIAL"),
		os.Getenv("MYSQL_DATABASE_CREDENTIAL"),
	)
	db, err := gorm.Open("mysql", connection)
	db.SingularTable(true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &loginInternalErrorResponse)
		return
	}
	defer db.Close()
	var userSession UserSession
	db.First(&userSession, "session_key = ?", cookie)
	c.JSON(http.StatusOK, gin.H{"is_logged_in": cookie == userSession.SessionKey, "error": nil})
}

func LoginHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "http://trackhours.co")
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
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

	connection := fmt.Sprintf(
		"%s:%s@/%s",
		os.Getenv("MYSQL_USERNAME_CREDENTIAL"),
		os.Getenv("MYSQL_PASSWORD_CREDENTIAL"),
		os.Getenv("MYSQL_DATABASE_CREDENTIAL"),
	)
	db, err := gorm.Open("mysql", connection)
	db.SingularTable(true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &loginInternalErrorResponse)
		return
	}
	defer db.Close()
	var user User
	db.First(&user, "username = ?", accountInformation.Username)
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
	userSession := UserSession{
		OwnerUsername: user.Username,
		SessionKey:    sessionKey,
	}
	db.Create(&userSession)
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
