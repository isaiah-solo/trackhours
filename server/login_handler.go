package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type AccountInformation struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Response struct {
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

func generateSessionKey() (string, error) {
	u, err := uuid.NewV4()
	return u.String(), err
}

func AccountCreationHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "http://trackhours.co")
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	db, err := EstablishConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Error: "Issue connecting to DB",
		})
		return
	}
	defer db.Close()
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Error: "Issue parsing body",
		})
		return
	}
	var accountInformation AccountInformation
	if err := json.Unmarshal(body, &accountInformation); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Error: "Issue parsing body json",
		})
		return
	}
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(accountInformation.Password),
		bcrypt.MinCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Error: "Issue encrypting",
		})
		return
	}
	user := User{
		Password: string(hash),
		Username: accountInformation.Username,
	}
	db.Create(&user)
	sessionKey, err := generateSessionKey()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Error: "Issue generating session key",
		})
		return
	}
	userSession := UserSession{
		OwnerUsername: user.Username,
		SessionKey:    sessionKey,
	}
	db.Create(&userSession)
	c.SetCookie("trackhours_session_key", sessionKey, 360000, "/", "", false, false)
	c.JSON(http.StatusOK, Response{
		Error: "",
	})
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
	db, err := EstablishConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Error: "Issue connecting to DB",
		})
		return
	}
	defer db.Close()
	var userSession UserSession
	db.First(&userSession, "session_key = ?", cookie)
	c.JSON(http.StatusOK, gin.H{
		"is_logged_in": cookie == userSession.SessionKey,
		"error":        nil,
	})
}

func LoginHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "http://trackhours.co")
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Error: "Issue parsing body",
		})
		return
	}
	var accountInformation AccountInformation
	if err := json.Unmarshal(body, &accountInformation); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Error: "Issue parsing body json",
		})
		return
	}
	db, err := EstablishConnection()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Error: "Issue connecting to DB",
		})
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
		c.JSON(http.StatusInternalServerError, Response{
			Error: "Issue encrypting",
		})
		return
	}
	// Generate user session key
	sessionKey, err := generateSessionKey()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Error: "Issue generating session key",
		})
		return
	}
	userSession := UserSession{
		OwnerUsername: user.Username,
		SessionKey:    sessionKey,
	}
	db.Create(&userSession)
	c.SetCookie(
		"trackhours_session_key",
		sessionKey,
		360000,
		"/",
		"",
		false,
		false,
	)
	c.JSON(http.StatusOK, Response{
		Error: "",
	})
}

func LogoutHandler(c *gin.Context) {
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Origin", "http://trackhours.co")
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	c.SetCookie(
		"trackhours_session_key",
		"",
		360000,
		"/",
		"",
		false,
		false,
	)
	c.JSON(http.StatusOK, Response{
		Error: "",
	})
}
