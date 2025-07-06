package main

import (
	"log/slog"
	"net/http"
	"time"
)

type totalTimeLogger struct {
	activities activities
}

type activity struct {
	name        string
	description string
	totalTime   time.Time // TODO - should be a duration
	sessions    []session
}

func (a *activity) viewSessions()   {}
func (a *activity) startActivity()  {}
func (a *activity) stopActivity()   {}
func (a *activity) deleteActivity() {}

type activities map[int]activity

type session struct {
	startTime time.Time
	endTime   time.Time
}

func (a *session) editStartTime() {}
func (a *session) editEndTime()   {}

var ttl totalTimeLogger

func main() {
	test()
	router := http.NewServeMux()

	router.HandleFunc("/", serveHome)
	router.HandleFunc("GET /addActivity", serveAddActivity)
	router.HandleFunc("GET /viewActivities", serveviewActivities)
	router.HandleFunc("GET /selectActivity", serveSelectActivity)

	router.HandleFunc("POST /addActivity", handleAddActivity)

	slog.Info("Starting server")
	err := http.ListenAndServe(":8888", router)
	if err != nil {
		slog.Error("Listening:", "err", err)
		return
	}
}
