package database

import (
	"aino-spring.com/aino_site/config"
	"gorm.io/gorm"
)

type Page struct {
  gorm.Model
  Path string
  Template string
  IsAdminPage bool
}

func (page Page) GetCompletePath(conf *config.Config) string {
  path := page.Path
  if page.IsAdminPage {
    path = conf.AdminPath + path
  }
  return path
}

