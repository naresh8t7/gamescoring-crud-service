package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type ScorekeepingCode string

const (
	Pitch ScorekeepingCode = "pitch"
	Ball  ScorekeepingCode = "Ball"
)

type ScorekeepingResult string

const (
	BallInPlay ScorekeepingResult = "ball_in_play"
	Strikeout  ScorekeepingResult = "strike_out"
)

type ScoringEvent struct {
	ID        string    `gorm:"primaryKey;column:id" json:"id"`
	GameID    string    `gorm:"primaryKey;column:game_id" json:"game_id"`
	Timestamp time.Time `gorm:"primaryKey;column:timestamp" json:"timestamp"`
	Data      GameData  `gorm:"primaryKey;column:data" json:"data"`
}

type GameData struct {
	Code       ScorekeepingCode `json:"code"`
	Attributes Attributes       `json:"attributes"`
}

type Attributes struct {
	AdvancesCount bool               `json:"advances_count"`
	Result        ScorekeepingResult `json:"result"`
}

func (a *GameData) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *GameData) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed for ConfiantMinionObject")
	}
	return json.Unmarshal(b, &a)
}
