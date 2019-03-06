package main

import (
  "encoding/json"
  "io/ioutil"
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/satori/go.uuid"
  "golang.org/x/crypto/bcrypt"

  _ "github.com/go-sql-driver/mysql"
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
  user := User {
    Password: string(hash),
    Username: accountInformation.Username,
  }
  userInsert, err := db.Prepare(
    "INSERT INTO user (password, username) VALUES (?, ?)",
  )
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

func LoginHandler(c *gin.Context) {
  c.Header("Access-Control-Allow-Credentials", "true")
  c.Header("Access-Control-Allow-Origin", "http://trackhours.co")
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
  // Set cookie
  http.SetCookie(c.Writer, &http.Cookie{
    Name: "trackhours_session_key",
    Value: sessionKey,
    MaxAge: 360000,
    Path: "/",
    Domain: "trackhours.co",
    Secure: false,
    HttpOnly: false,
  })
  //c.SetCookie("trackhours_session_key", sessionKey, 360000, "/", "", false, false)
  c.JSON(http.StatusOK, &successResponse)
}
