package main

import (
  "os"
  "path/filepath"
  "net/http"
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

func main() {
  router := gin.Default()
  router.Use(static.Serve("/static", static.LocalFile("../view/build/static", true)))
  rootDir := "../view/build"
  fileList := []string{}
  filepath.Walk(rootDir, func(path string, f os.FileInfo, err error) error {
    if !(f.IsDir()) {
      fileList = append(fileList, path)
    }
    return nil
  })
  router.LoadHTMLFiles(fileList...)
  router.GET("/", func(c *gin.Context) {
    c.HTML(
      http.StatusOK,
      "index.html",
      gin.H {},
    )
  })
  api := router.Group("/api")
  api.GET("/checklogin", CheckLoginHandler)
  api.POST("/account_creation", AccountCreationHandler)
  api.POST("/login", LoginHandler)
  router.Run(":8081")
}
