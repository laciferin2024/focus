package models

import (
	"time"

	"github.com/ayoisaiah/focus/config"
)

type SessionTimeline struct {
	// StartTime is the start of the session including
	// the start of a paused session
	StartTime time.Time `json:"start_time"`
	// EndTime is the end of a session including
	// when a session is paused or stopped prematurely
	EndTime time.Time `json:"end_time"`
}

type Session struct {
	StartTime time.Time         `json:"start_time"`
	EndTime   time.Time         `json:"end_time"`
	Name      config.SessType   `json:"name"`
	Tags      []string          `json:"tags"`
	Timeline  []SessionTimeline `json:"timeline"`
	Duration  time.Duration     `json:"duration"`
	Completed bool              `json:"completed"`
}

type Timer struct {
	Opts       *config.TimerConfig `json:"opts"`
	PausedTime time.Time           `json:"paused_time"`
	StartTime  time.Time           `json:"start_time"`
	SessionKey time.Time           `json:"session_key"`
	WorkCycle  int                 `json:"work_cycle"`
}
