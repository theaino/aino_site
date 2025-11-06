package models

import "time"

var DateFormat = "2006-01-02"

type Mood struct {
	ID uint `gorm:"primaryKey"`
	Value int
	Date time.Time
}

