package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
	"ttl/data"
)

func handleAddActivity(w http.ResponseWriter, r *http.Request) {
	// TODO - validate
	n := r.FormValue("name")
	d := r.FormValue("description")

	ttl.Activities = append(ttl.Activities, &data.Activity{Name: n, Description: d})

	ttl.SaveToJSON()

	fmt.Fprint(w, "Success!")
}

// TODO - awful to eyes
func handleStart(w http.ResponseWriter, r *http.Request) {
	msg := "Timer ongoing."
	act, id, err := getActivityFromStringID(r.URL.Query().Get("id"))
	if err != nil {
		slog.Error("Getting activity and id from url", "err", err)
		return
	}

	lastDuration := act.GetTotalTimeSpent()

	// Only start a new timer if the last session has a valid end time.
	if len(act.Sessions) == 0 {
		act.Sessions = append(act.Sessions, data.Session{StartTime: time.Now()})
		ttl.SaveToJSON()
	} else {
		var defaultDuration time.Time
		if act.Sessions[len(act.Sessions)-1].EndTime == defaultDuration { // is the end time not set?
			msg = ""
			w.Header().Add("HX-Trigger", "{\"show\":\"There is an ongoing timer.\"}")
		} else {
			act.Sessions = append(act.Sessions, data.Session{StartTime: time.Now()})
			ttl.SaveToJSON()
		}
	}

	err = SelectActivity(*act, id, lastDuration, msg).Render(r.Context(), w)
	if err != nil {
		slog.Warn("Rendering select activity component", "err", err)
		return
	}
}

// TODO - awful to eyes
func handleEnd(w http.ResponseWriter, r *http.Request) {
	msg := "Timer stopped."
	act, id, err := getActivityFromStringID(r.URL.Query().Get("id"))
	if err != nil {
		slog.Error("Getting activity and id from url", "err", err)
		return
	}

	// Only "end" a timer if the end time has not been set.
	var defaultDuration time.Time
	if len(act.Sessions) == 0 {
		msg = ""
		w.Header().Add("HX-Trigger", "{\"show\":\"There is no timer to be stopped.\"}")
	} else {
		if act.Sessions[len(act.Sessions)-1].EndTime == defaultDuration { // is the end time not set?
			act.Sessions[len(act.Sessions)-1].EndTime = time.Now()
			ttl.SaveToJSON()
		} else {
			msg = ""
			w.Header().Add("HX-Trigger", "{\"show\":\"There is no timer to be stopped.\"}")
		}
	}

	err = SelectActivity(*act, id, act.GetTotalTimeSpent(), msg).Render(r.Context(), w)
	if err != nil {
		slog.Warn("Rendering select activity component", "err", err)
		return
	}
}
