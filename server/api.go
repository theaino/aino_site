package server

import (
	"errors"
	"log"
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
)

type PostPreset struct {
  Title string `json:"title" binding:"required"`
  Abstract string `json:"abstract" binding:"required"`
  Contents string `json:"contents" binding:"required"`
  Public bool `json:"public" binding:"required"`
}

type UserPreset struct {
  Email string `json:"email" binding:"required"`
  Name string `json:"name" binding:"required"`
  Password string `json:"password" binding:"required"`
}

func (server *Server) SetupApiPages() {
  server.Router.POST("/api/posts/:id/edit", func (c *gin.Context) {
    _, isAdmin := server.CheckContext(c)
    if !isAdmin {
      c.JSON(http.StatusForbidden, gin.H{})
      return
    }
    id := c.Param("id")
    var preset PostPreset
    c.BindJSON(&preset)

    err := errors.Join(
      server.Database.SetPostTitle(id, preset.Title),
      server.Database.SetPostAbstract(id, preset.Abstract),
      server.Database.SetPostContents(id, preset.Contents),
      server.Database.SetPostPublic(id, preset.Public),
      )

    if err != nil {
      log.Println(err)
      c.JSON(http.StatusInternalServerError, gin.H{})
      return
    }
    c.JSON(http.StatusOK, gin.H{})
  })

  server.Router.POST("/api/new-post", func (c *gin.Context) {
    _, isAdmin := server.CheckContext(c)
    if !isAdmin {
      c.JSON(http.StatusForbidden, gin.H{})
      return
    }
    var preset PostPreset
    c.BindJSON(&preset)
    id, err := server.Database.NewPost(preset.Title, preset.Abstract, preset.Contents, preset.Public);
    
    if err != nil {
      log.Println(err)
      c.JSON(http.StatusInternalServerError, gin.H{})
      return
    }
    c.JSON(http.StatusOK, gin.H{"id": id})
  })

  server.Router.POST("/api/posts/:id/delete", func (c *gin.Context) {
    _, isAdmin := server.CheckContext(c)
    if !isAdmin {
      c.JSON(http.StatusForbidden, gin.H{})
      return
    }
    id := c.Param("id")

    err := server.Database.DeletePost(id)

    if err != nil {
      log.Println(err)
      c.JSON(http.StatusInternalServerError, gin.H{})
      return
    }
    c.JSON(http.StatusOK, gin.H{})
  })

  server.Router.POST("/api/signup", func (c *gin.Context) {
    var preset UserPreset
    c.BindJSON(&preset)
    _, err := mail.ParseAddress(preset.Email)
    if err != nil {
      c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 1})
      return
    }
    _, err = server.Database.FetchUserByName(preset.Name)
    if err == nil {
      c.JSON(http.StatusConflict, gin.H{"code": 2})
      return
    }
    id, err := NewUser(server.Database, preset.Email, preset.Name, preset.Password)
    if err != nil {
      c.JSON(http.StatusConflict, gin.H{"code": 3})
      return
    }
    c.JSON(http.StatusOK, gin.H{"id": id})
  })

  server.Router.GET("/api/login", func (c *gin.Context) {
    isAuthed, isAdmin := server.CheckContext(c)
    c.JSON(http.StatusOK, gin.H{"authed": isAuthed, "admin": isAdmin})
  })

  server.Router.POST("/api/settings/:key/set/:value", func (c *gin.Context) {
    _, isAdmin := server.CheckContext(c)
    if !isAdmin {
      c.JSON(http.StatusForbidden, gin.H{})
      return
    }
    key := c.Param("key")
    value := c.Param("value")
    err := server.Database.SetSetting(key, value)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
      return
    }
    c.JSON(http.StatusOK, gin.H{})
  })
}

