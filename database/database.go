package database

import (
	"aino-spring.com/aino_site/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Connection struct {
  Database *gorm.DB
  Config *config.Config
}

func NewConnetion(conf *config.Config) (*Connection, error) {
  connection := new(Connection)
  connection.Config = conf
  var err error
  connection.Database, err = gorm.Open(mysql.Open(conf.MysqlDsn))
  if err != nil {
    return nil, err
  }
  return connection, nil
}

func (connection *Connection) Migrate() {
  connection.Database.AutoMigrate(&Page{}, &Post{})
}

func (connection *Connection) FetchPages() ([]Page, error) {
  var pages []Page
  result := connection.Database.Find(&pages)
  return pages, result.Error
}

func (connection *Connection) FetchPosts() ([]Post, error) {
  var posts []Post
  result := connection.Database.Find(&posts)
  return posts, result.Error
}

func (connection *Connection) FetchPost(id string) (Post, error) {
  var post Post
  result := connection.Database.First(&post, id)
  return post, result.Error
}

