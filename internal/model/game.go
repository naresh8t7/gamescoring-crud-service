package model

import "time"

type Game struct {
	ID     string    `gorm:"primaryKey;column:id" json:"id"`
	Start  time.Time `gorm:"column:start_time" json:"start"`
	End    time.Time `gorm:"column:end_time" json:"end"`
	Arrive time.Time `gorm:"column:arrive" json:"arrive"`
}
