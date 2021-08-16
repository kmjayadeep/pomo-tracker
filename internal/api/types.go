package api

import (
  "time"
)

type PomoStatus string

var (
  PomoRunning PomoStatus = "RUNNING"
  PomoStopped PomoStatus = "STOPPED"
)

type Pomo struct {
  ID int `json:"id"`
  StartedAt time.Time `json:"startedAt"`
  StoppedAt time.Time `json:"stoppedAt"`
  Status PomoStatus `json:"status"`
  DurationSeconds int `json:"durationSeconds"`
}
