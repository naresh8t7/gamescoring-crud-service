package model

import "time"

type Game struct {
	ID     string    `json:"id"`
	Start  time.Time `json:"start"`
	End    time.Time `json:"end"`
	Arrive time.Time `json:"arrive"`
}
