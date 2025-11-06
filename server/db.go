package server

import (
	"ainosite/models"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func NewDB() *DB {
	return new(DB)
}

func (d *DB) OpenSqlite(path string) error {
	conn, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	d.DB = conn
	return err
}

func (d *DB) OpenMysql(dsn string) error {
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	d.DB = conn
	return err
}

func G[T any](db *DB) gorm.Interface[T] {
	return gorm.G[T](db.DB)
}

func (d *DB) Migrate() error {
	return d.AutoMigrate(&models.Mood{})
}
