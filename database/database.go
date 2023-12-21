package database

import (
	"time"

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

func (connection *Connection) SetPostTitle(id string, title string) error {
  var post Post
  result := connection.Database.First(&post, id)
  if result.Error != nil {
    return result.Error
  }
  post.Title = title
  result = connection.Database.Save(&post)
  return result.Error
}

func (connection *Connection) SetPostAbstract(id string, abstract string) error {
  var post Post
  result := connection.Database.First(&post, id)
  if result.Error != nil {
    return result.Error
  }
  post.Abstract = abstract
  result = connection.Database.Save(&post)
  return result.Error
}

func (connection *Connection) SetPostContents(id string, contents string) error {
  var post Post
  result := connection.Database.First(&post, id)
  if result.Error != nil {
    return result.Error
  }
  post.Contents = contents
  result = connection.Database.Save(&post)
  return result.Error
}

func (connection *Connection) SetPostPublic(id string, public bool) error {
  var post Post
  result := connection.Database.First(&post, id)
  if result.Error != nil {
    return result.Error
  }
  post.Public = public
  result = connection.Database.Save(&post)
  return result.Error
}

func (connection *Connection) NewPost(title, abstract, contents string, public bool) (uint, error) {
  post := Post{Title: title, Abstract: abstract, Contents: contents, Public: public, Date: time.Now()}
  result := connection.Database.Create(&post)
  return post.ID, result.Error
}

func (connection *Connection) DeletePost(id string) error {
  result := connection.Database.Delete(&Post{}, id)
  return result.Error
}

