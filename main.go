package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"ttl/data"
	"ttl/logger"
)

var ttl logger.TotalTimeLogger

func main() {
	err := ttl.ReadFromJson(&ttl)
	if err != nil {
		slog.Error("Reading", "err", err)
		return
	}

	router := http.NewServeMux()

	router.HandleFunc("/", serveHome)

	router.HandleFunc("GET /addActivity", serveAddActivity)
	router.HandleFunc("GET /viewActivities", serveviewActivities)
	router.HandleFunc("GET /selectActivity", serveSelectActivity)

	router.HandleFunc("POST /addActivity", handleAddActivity)
	router.HandleFunc("POST /start", handleStart)
	router.HandleFunc("POST /end", handleEnd)

	slog.Info("Starting server")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		slog.Error("Listening:", "err", err)
		return
	}
}

// TODO - place this somewhere reasonable
func getActivityFromStringID(actID string) (act *data.Activity, id int, err error) {
	id, err = strconv.Atoi(actID)
	if err != nil {
		return act, id, fmt.Errorf("converting id from str to int: %w", err)
	}

	activity := ttl.Activities[id]
	return activity, id, nil
}
