package main

import (
	"log"
	"net/http"
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

func InitHeader(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", BackendOrigin)
}

func main() {
	http.HandleFunc("/api/check_login", DBHandlerFor(CheckLoginHandler))
	http.HandleFunc("/api/logout", DBHandlerFor(LogoutHandler))
	http.HandleFunc("/api/account_creation", DBHandlerFor(AccountCreationHandler))
	http.HandleFunc("/api/login", DBHandlerFor(LoginHandler))
	log.Fatal(http.ListenAndServe(":8081", nil))
}
