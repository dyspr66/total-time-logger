package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
)

type totalTimeLogger struct {
	activities []*activity
}

var ttl totalTimeLogger

func main() {
	// test()
	err := ttl.readFromJson()
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
	err = http.ListenAndServe(":8888", router)
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
