package main

import (
	"time"
)

type activity struct {
	Name        string     `json:name`
	Description string     `json:description`
	Sessions    []Sessions `json:sessions`
}

func (a *activity) getTotalTimeSpent() time.Duration {
	var defaultTime time.Time

	var totalDuration time.Duration
	for _, session := range a.Sessions {
		if session.EndTime != defaultTime {
			totalDuration += session.EndTime.Sub(session.StartTime)
		}
	}
	return totalDuration
}

func (a *activity) viewSessions()   {}
func (a *activity) startActivity()  {}
func (a *activity) stopActivity()   {}
func (a *activity) deleteActivity() {}

type activities []*activity
