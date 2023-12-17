package database

import (
	"aino-spring.com/aino_site/config"
	"aino-spring.com/aino_site/server"
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
  connection.Database.AutoMigrate(&Page{})
}

func (connection *Connection) FetchPager() (*server.Pager, error) {
  pager := server.NewPager()
  var pages []Page
  result := connection.Database.Find(&pages)
  for _, page := range pages {
    pager.AddPage(page.GetCompletePath(connection.Config), page.Template)
  }
  return pager, result.Error
}

