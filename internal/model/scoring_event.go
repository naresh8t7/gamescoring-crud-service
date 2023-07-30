package model

import "time"

type ScoringEvent struct {
	ID        string    `json:"id"`
	GameID    string    `json:"game_id"`
	Timestamp time.Time `json:"timestamp"`
	Data      GameData  `json:"data"`
}

type GameData struct {
	Code       string     `json:"code"`
	Attributes Attributes `json:"attributes"`
}

type Attributes struct {
	AdvancesCount bool   `json:"advances_count"`
	Result        string `json:"result"`
}
