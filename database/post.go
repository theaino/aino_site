package database

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
  gorm.Model
  Title string
  Contents string
  Date time.Time
  Public bool
}

