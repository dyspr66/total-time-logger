package data

import (
	"time"
)

type Activity struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Sessions    []Session `json:"sessions"`
}

func (a *Activity) GetTotalTimeSpent() time.Duration {
	var defaultTime time.Time

	var totalDuration time.Duration
	for _, session := range a.Sessions {
		if session.EndTime != defaultTime {
			totalDuration += session.EndTime.Sub(session.StartTime)
		}
	}

	return totalDuration
}

type Activities []*Activity
