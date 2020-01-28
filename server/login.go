package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

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

func AccountCreationHandler(w http.ResponseWriter, r *http.Request) {
	InitHeader(w)
	db, err := EstablishConnection()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(CheckLoginResponse{
			Error: "Issue connecting to DB",
		})
		return
	}
	defer db.Close()
	accountInformation, err := generateAccountInformation(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(CheckLoginResponse{
			Error: "Issue parsing body json",
		})
		return
	}
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(accountInformation.Password),
		bcrypt.MinCost,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(CheckLoginResponse{
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
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(CheckLoginResponse{
			Error: "Issue generating session key",
		})
		return
	}
	userSession := UserSession{
		OwnerUsername: user.Username,
		SessionKey:    sessionKey,
	}
	db.Create(&userSession)
	http.SetCookie(w, &http.Cookie{
		Name:     "trackhours_session_key",
		Value:    sessionKey,
		MaxAge:   360000,
		Path:     "/",
		Domain:   "",
		Secure:   false,
		HttpOnly: false,
	})
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(CheckLoginResponse{
		Error: "",
	})
}

func CheckLoginHandler(w http.ResponseWriter, r *http.Request) {
	InitHeader(w)
	cookie, err := r.Cookie("trackhours_session_key")
	cookieValue, _ := url.QueryUnescape(cookie.Value)
	if err != nil || cookieValue == "" {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(CheckLoginResponse{
			Error:      err.Error(),
			IsLoggedIn: false,
		})
		return
	}
	db, err := EstablishConnection()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{
			Error: "Issue connection to DB",
		})
		return
	}
	defer db.Close()
	var userSession UserSession
	db.First(&userSession, "session_key = ?", cookieValue)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(CheckLoginResponse{
		Error:      "",
		IsLoggedIn: cookieValue == userSession.SessionKey,
	})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	InitHeader(w)
	accountInformation, err := generateAccountInformation(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{
			Error: "Issue parsing body json",
		})
		return
	}
	db, err := EstablishConnection()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{
			Error: "Issue connecting to DB",
		})
		return
	}
	defer db.Close()
	var user User
	db.First(&user, "username = ?", accountInformation.Username)
	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(accountInformation.Password),
	); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{
			Error: "Issue encrypting",
		})
		return
	}
	sessionKey, err := generateSessionKey()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{
			Error: "Issue generating session key",
		})
		return
	}
	userSession := UserSession{
		OwnerUsername: user.Username,
		SessionKey:    sessionKey,
	}
	db.Create(&userSession)
	http.SetCookie(w, &http.Cookie{
		Name:     "trackhours_session_key",
		Value:    sessionKey,
		MaxAge:   360000,
		Path:     "/",
		Domain:   "",
		Secure:   false,
		HttpOnly: false,
	})
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Error: "",
	})
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	InitHeader(w)
	http.SetCookie(w, &http.Cookie{
		Name:     "trackhours_session_key",
		Value:    "",
		MaxAge:   360000,
		Path:     "/",
		Domain:   "",
		Secure:   false,
		HttpOnly: false,
	})
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Error: "",
	})
}
