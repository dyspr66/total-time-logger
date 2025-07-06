package main

import (
	"log/slog"
	"net/http"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	err := Home().Render(r.Context(), w)
	if err != nil {
		slog.Error("Rendering home page", "err", err)
		return
	}
}

func serveAddActivity(w http.ResponseWriter, r *http.Request) {
	err := AddActivity().Render(r.Context(), w)
	if err != nil {
		slog.Error("Rendering home page", "err", err)
		return
	}
}
