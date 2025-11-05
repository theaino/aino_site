package server

import (
	"ainosite/models"

	"github.com/glebarez/sqlite"
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

func G[T any](db *DB) gorm.Interface[T] {
	return gorm.G[T](db.DB)
}

func (d DB) Migrate() error {
	return d.AutoMigrate(&models.Mood{})
}
