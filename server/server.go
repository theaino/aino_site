package server

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"strings"

	"aino-spring.com/aino_site/config"
	"aino-spring.com/aino_site/database"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Server struct {
  Router *gin.Engine
  Config *config.Config
  Database *database.Connection
}

func NewServer(db *database.Connection, conf *config.Config) *Server {
  if conf.Release {
    gin.SetMode(gin.ReleaseMode)
  }

  server := new(Server)
  server.Database = db
  server.Config = conf

  server.Router = gin.Default()

  store := cookie.NewStore([]byte(conf.SessionSecret))
  server.Router.Use(sessions.Sessions("ainosession", store))

  server.Router.SetFuncMap(template.FuncMap{
    "dict": func(values ...interface{}) (map[string]interface{}, error) {
      if len(values) % 2 != 0 {
        return nil, errors.New("Invalid dict call")
      }
      dict := make(map[string]interface{}, len(values) / 2)
      for i := 0; i < len(values); i += 2 {
        key, ok := values[i].(string)
        if !ok {
          return nil, errors.New("Dict keys must be strings")
        }
        dict[key] = values[i + 1]
      }
      return dict, nil
    },
    "navitems": func(values ...string) [][]string {
      if len(values) % 3 != 0 {
        return nil
      }
      items := make([][]string, 0)
      currentItem := make([]string, 0)
      for idx, value := range values {
        currentItem = append(currentItem, value)
        if idx % 3 == 2 {
          items = append(items, currentItem)
          currentItem = make([]string, 0)
          continue
        }
      }
      return items
    },
    "navtemplate": func(value []string) string {
      return value[0]
    },
    "navtitle": func(value []string) string {
      return value[1]
    },
    "navhref": func(value []string) string {
      return value[2]
    },
  })

  server.Router.LoadHTMLGlob("templates/**/*")

  pages, err := db.FetchPages()
  if err != nil {
    log.Panic(err)
  }
  pager := NewPagerFromDBPages(pages)
  server.LoadPager(pager)

  server.Router.Static("/static", "static")
  server.Router.StaticFile("/favicon.ico", "favicon.ico")
  server.Router.NoRoute(server.GetHandler(http.StatusNotFound, "not-found", gin.H{}))

  return server
}

func (server *Server) GetValues(template string, c *gin.Context, values gin.H) gin.H {
  caser := cases.Title(language.English)
  title := strings.ReplaceAll(template, "-", " ")
  title = strings.ReplaceAll(title, "_", " ")
  title = caser.String(title)
  values["title"] = title
  values["template"] = template
  isAuthed, isAdmin := server.CheckContext(c)
  values["authed"] = isAuthed
  values["admin"] = isAdmin
  return values
}

func (server *Server) GetHandler(status int, template string, values gin.H) func (*gin.Context) {
  return func (c *gin.Context) {
    c.HTML(http.StatusOK, template, server.GetValues(template, c, values))
  }
}

func (server *Server) GetAdminHandler(status int, template string, values gin.H) func (*gin.Context) {
  return func (c *gin.Context) {
    _, isAdmin := server.CheckContext(c)
    if isAdmin {
      c.HTML(http.StatusOK, template, server.GetValues(template, c, values))
      return
    }
    c.Redirect(http.StatusTemporaryRedirect, "/login")
  }
}

func (server *Server) LoadPager(pager *Pager) {
  for _, path := range pager.GetPaths() {
    var handler func (*gin.Context)
    if pager.IsAdmin(path) {
      handler = server.GetAdminHandler(http.StatusOK, pager.GetTemplate(path), gin.H{})
    } else {
      handler = server.GetHandler(http.StatusOK, pager.GetTemplate(path), gin.H{})
    }
    server.Router.GET(path, handler)
  }
}

func (server *Server) Run(address string) {
  if server.Config.Tls {
    server.Router.RunTLS(address, server.Config.CertPath, server.Config.PrivateKeyPath)
  } else {
    server.Router.Run(address)
  }
}

