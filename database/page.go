package database

import (
	"gorm.io/gorm"
)

type Page struct {
  gorm.Model
  Path string
  Template string
  IsAdminPage bool
}

func (page Page) GetCompletePath() string {
  path := page.Path
  if page.IsAdminPage {
    path = "/admin" + path
  }
  return path
}

func (connection *Connection) FetchPages() ([]Page, error) {
  var pages []Page
  result := connection.Database.Find(&pages)
  return pages, result.Error
}

