package main

import (
	"github.com/gin-gonic/gin"
)

type Event int

type Schedule struct {
	Name       string      `json:"name"`
	TimeEvents []TimeEvent `json:"time_events"`
}

type TimeEvent struct {
	EventType string `json:"event_type"`
	Time      int    `json:"time"`
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
	api := router.Group("/api")
	api.GET("/check_login", CheckLoginHandler)
	api.POST("/account_creation", AccountCreationHandler)
	api.POST("/login", LoginHandler)
	router.Run(":8081")
}
