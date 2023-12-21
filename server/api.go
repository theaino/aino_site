package server

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostPreset struct {
  Title string `json:"title" binding:"required"`
  Abstract string `json:"abstract" binding:"required"`
  Contents string `json:"contents" binding:"required"`
  Public bool `json:"public" binding:"required"`
}

func (server *Server) SetupApiPages() {
  server.Router.POST("/api/posts/:id/edit", func (c *gin.Context) {
    isAuthed := server.IsAuthed(c)
    if !isAuthed {
      c.JSON(http.StatusForbidden, gin.H{})
      return;
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
    isAuthed := server.IsAuthed(c)
    if !isAuthed {
      c.JSON(http.StatusForbidden, gin.H{})
      return;
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
    isAuthed := server.IsAuthed(c)
    if !isAuthed {
      c.JSON(http.StatusForbidden, gin.H{})
      return;
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
}

