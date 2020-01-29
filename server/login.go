package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/jinzhu/gorm"
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

type CheckLoginResponse struct {
	Error      string `json:"error"`
	IsLoggedIn bool   `json:"is_logged_in"`
}

type User struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type UserSession struct {
	OwnerUsername string `json:"owner_username"`
	SessionKey    string `json:"session_key"`
}

func createInternalServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(CheckLoginResponse{
		Error: err.Error(),
	})
}

func createSuccessResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Error: "",
	})
}

func generateAccountInformation(r *http.Request) (AccountInformation, error) {
	var accountInformation AccountInformation
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return accountInformation, err
	}
	if err := json.Unmarshal(body, &accountInformation); err != nil {
		return accountInformation, err
	}
	return accountInformation, nil
}

func generateSessionKey() (string, error) {
	u, err := uuid.NewV4()
	return u.String(), err
}

func removeCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "trackhours_session_key",
		Value:    "",
		MaxAge:   -1,
		Path:     "/",
		Domain:   "",
		Secure:   false,
		HttpOnly: false,
	})
}

func setCookie(w http.ResponseWriter, sessionKey string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "trackhours_session_key",
		Value:    sessionKey,
		MaxAge:   360000,
		Path:     "/",
		Domain:   "",
		Secure:   false,
		HttpOnly: false,
	})
}

func AccountCreationHandler(
	db *gorm.DB,
	w http.ResponseWriter,
	r *http.Request,
) {
	accountInformation, err := generateAccountInformation(r)
	if err != nil {
		createInternalServerError(w, err)
		return
	}
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(accountInformation.Password),
		bcrypt.MinCost,
	)
	if err != nil {
		createInternalServerError(w, err)
		return
	}
	user := User{
		Password: string(hash),
		Username: accountInformation.Username,
	}
	db.Create(&user)
	sessionKey, err := generateSessionKey()
	if err != nil {
		createInternalServerError(w, err)
		return
	}
	userSession := UserSession{
		OwnerUsername: user.Username,
		SessionKey:    sessionKey,
	}
	db.Create(&userSession)
	setCookie(w, sessionKey)
	createSuccessResponse(w)
}

func CheckLoginHandler(
	db *gorm.DB,
	w http.ResponseWriter,
	r *http.Request,
) {
	cookie, err := r.Cookie("trackhours_session_key")
	if err != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(CheckLoginResponse{
			Error:      err.Error(),
			IsLoggedIn: false,
		})
		return
	}
	cookieValue, _ := url.QueryUnescape(cookie.Value)
	if err != nil || cookieValue == "" {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(CheckLoginResponse{
			Error:      err.Error(),
			IsLoggedIn: false,
		})
		return
	}
	var userSession UserSession
	db.First(&userSession, "session_key = ?", cookieValue)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(CheckLoginResponse{
		Error:      "",
		IsLoggedIn: cookieValue == userSession.SessionKey,
	})
}

func LoginHandler(
	db *gorm.DB,
	w http.ResponseWriter,
	r *http.Request,
) {
	accountInformation, err := generateAccountInformation(r)
	if err != nil {
		createInternalServerError(w, err)
		return
	}
	var user User
	db.First(&user, "username = ?", accountInformation.Username)
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(accountInformation.Password),
	); err != nil {
		createInternalServerError(w, err)
		return
	}
	sessionKey, err := generateSessionKey()
	if err != nil {
		createInternalServerError(w, err)
		return
	}
	userSession := UserSession{
		OwnerUsername: user.Username,
		SessionKey:    sessionKey,
	}
	db.Create(&userSession)
	setCookie(w, sessionKey)
	w.WriteHeader(http.StatusOK)
	createSuccessResponse(w)
}

func LogoutHandler(
	db *gorm.DB,
	w http.ResponseWriter,
	r *http.Request,
) {
	removeCookie(w)
	w.WriteHeader(http.StatusOK)
	createSuccessResponse(w)
}
