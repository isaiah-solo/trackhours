package main

import (
  "github.com/gin-gonic/contrib/static"
  "github.com/gin-gonic/gin"
)

type Event int

type Schedule struct {
  Name string `json:"name"`
  TimeEvents []TimeEvent `json:"time_events"`
}

type TimeEvent struct {
  EventType string `json:"event_type"`
  Time int `json:"time"`
}

const (
  EndBreakEvent Event = iota
  EndDayEvent
  EndLunchEvent
  StartBreakEvent
  StartDayEvent
  StartLunchEvent
)

var eventNames = [...]string{
  "end_break_event",
  "end_day_event",
  "end_lunch_event",
  "start_break_event",
  "start_day_event",
  "start_lunch_event",
}

func test(c *gin.Context) {
  c.Header("Access-Control-Allow-Credentials", "true")
  c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
  _, err := c.Cookie("trackhours_session_key")
  isLoggedIn := true
  if err != nil {
    isLoggedIn = false
  }
  c.JSON(200, gin.H {"is_logged_in": isLoggedIn, "error": err})
}

func main() {
  router := gin.Default()
  router.Use(static.Serve("/", static.LocalFile("../view/build", true)))
  api := router.Group("/api")
  api.GET("/checklogin", test)
  api.POST("/account_creation", AccountCreationHandler)
  api.POST("/login", LoginHandler)
  router.Run(":8081")
}
