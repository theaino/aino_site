package main

import (
	"ainosite/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func NewDB(path string) (db *DB, err error) {
	conn, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return
	}
	db = &DB{conn}
	return
}

func (d DB) Migrate() error {
	return d.AutoMigrate(&models.Mood{})
}
