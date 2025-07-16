package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

type totalTimeLogger struct {
	activities map[int]*activity // map
}

type activities map[int]*activity

type activity struct {
	name        string
	description string
	timer       time.Ticker
	totalTime   time.Duration
	sessions    []session
}

func (a *activity) viewSessions()   {}
func (a *activity) startActivity()  {}
func (a *activity) stopActivity()   {}
func (a *activity) deleteActivity() {}

type session struct {
	startTime time.Time
	endTime   time.Time
}

func (a *session) editStartTime() {}
func (a *session) editEndTime()   {}

var ttl totalTimeLogger

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", serveHome)

	router.HandleFunc("GET /addActivity", serveAddActivity)
	router.HandleFunc("GET /viewActivities", serveviewActivities)
	router.HandleFunc("GET /selectActivity", serveSelectActivity)

	router.HandleFunc("POST /addActivity", handleAddActivity)
	router.HandleFunc("POST /start", handleStart)
	router.HandleFunc("POST /end", handleEnd)

	slog.Info("Starting server")
	err := http.ListenAndServe(":8888", router)
	if err != nil {
		slog.Error("Listening:", "err", err)
		return
	}
}

func getActivityFromStringID(actID string) (act *activity, id int, err error) {
	id, err = strconv.Atoi(actID)
	if err != nil {
		return act, id, fmt.Errorf("converting id from str to int: %w", err)
	}

	activity := ttl.activities[id]
	return activity, id, nil
}
