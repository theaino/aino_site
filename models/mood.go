package models

import "time"

type Mood struct {
	ID uint `gorm:"primaryKey"`
	Value int
	Date time.Time
}
