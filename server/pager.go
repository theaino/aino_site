package server

import (
	"aino-spring.com/aino_site/database"
)

type PagerEntry struct {
  Template string
  IsAdmin bool
}

type Pager map[string]PagerEntry

func NewPager() *Pager {
  pager := make(Pager)
  return &pager
}

func (pager *Pager) GetTemplate(path string) string {
  return (*pager)[path].Template
}

func (pager *Pager) IsAdmin(path string) bool {
  return (*pager)[path].IsAdmin
}

func (pager *Pager) GetPaths() (paths []string) {
  paths = make([]string, 0)
  for key := range *pager {
    paths = append(paths, key)
  }
  return
}

func (pager *Pager) AddPage(path, template string, isAdmin bool) {
  (*pager)[path] = PagerEntry{Template: template, IsAdmin: isAdmin}
}

func NewPagerFromDBPages(pages []database.Page) *Pager {
  pager := NewPager()
  for _, page := range pages {
    pager.AddPage(page.GetCompletePath(), page.Template, page.IsAdminPage)
  }
  return pager
}

