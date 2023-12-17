package server

import (
	"aino-spring.com/aino_site/config"
	"aino-spring.com/aino_site/database"
)

type Pager map[string]string

func NewPager() *Pager {
  pager := make(Pager)
  return &pager
}

func (pager *Pager) GetTemplate(path string) string {
  return (*pager)[path]
}

func (pager *Pager) GetPaths() (paths []string) {
  paths = make([]string, 0)
  for key := range *pager {
    paths = append(paths, key)
  }
  return
}

func (pager *Pager) AddPage(path string, template string) {
  (*pager)[path] = template
}

func NewPagerFromDBPages(conf *config.Config, pages []database.Page) *Pager {
  pager := NewPager()
  for _, page := range pages {
    pager.AddPage(page.GetCompletePath(conf), page.Template)
  }
  return pager
}

