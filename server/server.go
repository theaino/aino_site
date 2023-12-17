package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
  router *gin.Engine
}

func NewServer() *Server {
  server := new(Server)
  server.router = gin.Default()
  server.router.LoadHTMLGlob("templates/**/*")

  server.router.NoRoute(func (c *gin.Context) {
    c.HTML(http.StatusNotFound, "not-found", gin.H{})
  })

  return server
}

func (server *Server) LoadPager(pager *Pager) {
  for _, path := range pager.GetPaths() {
    server.router.GET(path, func (c *gin.Context) {
      c.HTML(http.StatusOK, pager.GetTemplate(path), gin.H{})
    })
  }
}

func (server *Server) Run(address string) {
  server.router.Run(address)
}

