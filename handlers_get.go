package main

import (
	"log/slog"
	"net/http"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	err := Home().Render(r.Context(), w)
	if err != nil {
		slog.Warn("Rendering home page", "err", err)
		return
	}
}

func serveAddActivity(w http.ResponseWriter, r *http.Request) {
	err := AddActivity().Render(r.Context(), w)
	if err != nil {
		slog.Warn("Rendering add activity component", "err", err)
		return
	}
}

func serveviewActivities(w http.ResponseWriter, r *http.Request) {
	err := ViewActivities(ttl.Activities).Render(r.Context(), w)
	if err != nil {
		slog.Warn("Rendering view activities component", "err", err)
		return
	}
}

func serveSelectActivity(w http.ResponseWriter, r *http.Request) {
	act, id, err := getActivityFromStringID(r.URL.Query().Get("id"))
	if err != nil {
		slog.Error("Getting activity and id from url", "err", err)
		return
	}

	err = SelectActivity(*act, id, act.GetTotalTimeSpent(), "").Render(r.Context(), w)
	if err != nil {
		slog.Warn("Rendering select activity component", "err", err)
		return
	}
}
