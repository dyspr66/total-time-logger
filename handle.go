package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func handleAddActivity(w http.ResponseWriter, r *http.Request) {
	var newActivity activity

	// TODO - validate
	name := r.FormValue("name")
	description := r.FormValue("description")

	newActivity.name = name
	newActivity.description = description

	ttl.activities[len(ttl.activities)+1] = &newActivity

	// TODO - update data storage
	fmt.Fprint(w, "Success!")
}

func handleStart(w http.ResponseWriter, r *http.Request) {
	act, _, err := getActivityFromStringID(r.URL.Query().Get("id"))
	if err != nil {
		slog.Error("Getting activity and id from url", "err", err)
		return
	}

	act.sessions = append(act.sessions, session{startTime: time.Now()})
	fmt.Println("timer start!")
}

func handleEnd(w http.ResponseWriter, r *http.Request) {
	act, _, err := getActivityFromStringID(r.URL.Query().Get("id"))
	if err != nil {
		slog.Error("Getting activity and id from url", "err", err)
		return
	}

	act.sessions[len(act.sessions)-1].endTime = time.Now()
	fmt.Println("timer end!")
}
