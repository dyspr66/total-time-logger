package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func handleAddActivity(w http.ResponseWriter, r *http.Request) {
	// TODO - validate
	n := r.FormValue("name")
	d := r.FormValue("description")
	ttl.activities = append(ttl.activities, &activity{Name: n, Description: d})

	// TODO - update data storage
	ttl.saveToJson()

	fmt.Fprint(w, "Success!")
}

// TODO - handle pressing start twice.
func handleStart(w http.ResponseWriter, r *http.Request) {
	act, _, err := getActivityFromStringID(r.URL.Query().Get("id"))
	if err != nil {
		slog.Error("Getting activity and id from url", "err", err)
		return
	}

	slog.Info("Timer started.")

	act.Sessions = append(act.Sessions, Sessions{StartTime: time.Now()})
	ttl.saveToJson()
}

// TODO - handle pressing end twice.
func handleEnd(w http.ResponseWriter, r *http.Request) {
	act, _, err := getActivityFromStringID(r.URL.Query().Get("id"))
	if err != nil {
		slog.Error("Getting activity and id from url", "err", err)
		return
	}

	slog.Info("Timer ended.")

	act.Sessions[len(act.Sessions)-1].EndTime = time.Now()
	ttl.saveToJson()
}
