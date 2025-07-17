package main

import "time"

type Sessions struct {
	StartTime time.Time
	EndTime   time.Time
}

func (a *Sessions) editStartTime() {}
func (a *Sessions) editEndTime()   {}
