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
	act, id, err := getActivityFromStringID(r.URL.Query().Get("id"))
	if err != nil {
		slog.Error("Getting activity and id from url", "err", err)
		return
	}

	lastDuration := act.getTotalTimeSpent()

	act.Sessions = append(act.Sessions, Sessions{StartTime: time.Now()})
	ttl.saveToJson()

	err = SelectActivity(*act, id, lastDuration, "Timer ongoing.").Render(r.Context(), w)
	if err != nil {
		slog.Error("Rendering select activity component", "err", err)
		return
	}
}

// TODO - handle pressing end twice.
func handleEnd(w http.ResponseWriter, r *http.Request) {
	act, id, err := getActivityFromStringID(r.URL.Query().Get("id"))
	if err != nil {
		slog.Error("Getting activity and id from url", "err", err)
		return
	}

	act.Sessions[len(act.Sessions)-1].EndTime = time.Now()
	ttl.saveToJson()

	err = SelectActivity(*act, id, act.getTotalTimeSpent(), "Timer stoppped.").Render(r.Context(), w)
	if err != nil {
		slog.Error("Rendering select activity component", "err", err)
		return
	}
}
