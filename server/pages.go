package server

import (
	"html/template"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
)

func (server *Server) SetupManualPages() {
  server.Router.GET("/posts", func (c *gin.Context) {
    posts, err := server.Database.FetchPosts()
    if err != nil {
      log.Panic(err)
    }
    sort.Slice(posts, func (i, j int) bool {
      return posts[i].Date.Unix() > posts[j].Date.Unix()
    })
    c.HTML(http.StatusOK, "posts", server.GetValues("posts", c, gin.H{"posts": posts}))
  })

  server.Router.GET("/posts/:id", func (c *gin.Context) {
    id := c.Param("id")
    post, err := server.Database.FetchPost(id)
    if err != nil {
      c.Redirect(http.StatusTemporaryRedirect, "/posts")
      return
    }
    _, isAdmin := server.CheckContext(c)
    if !(post.Public || isAdmin) {
      c.Redirect(http.StatusTemporaryRedirect, "/posts")
      return
    }
    postMap := gin.H{"id": id, "title": post.Title, "date": post.Date.Format("January 02, 2006"), "abstract": post.Abstract, "contents": template.HTML(post.Contents), "public": post.Public}
    c.HTML(http.StatusOK, "post", server.GetValues("post", c, gin.H{"post": postMap}))
  })

  server.Router.GET("/posts/:id/edit", func (c *gin.Context) {
    id := c.Param("id")
    _, isAdmin := server.CheckContext(c)
    if !isAdmin {
      c.Redirect(http.StatusTemporaryRedirect, "/posts/" + id)
      return
    }
    post, err := server.Database.FetchPost(id)
    if err != nil {
      c.Redirect(http.StatusTemporaryRedirect, "/posts")
      return
    }
    postMap := gin.H{"id": id, "title": post.Title, "date": post.Date.Format("January 02, 2006"), "abstract": post.Abstract, "contents": post.Contents, "public": post.Public}
    c.HTML(http.StatusOK, "edit-post", server.GetValues("edit-post", c, gin.H{"post": postMap}))
  })

  server.Router.GET("/new-post", func (c *gin.Context) {
    _, isAdmin := server.CheckContext(c)
    if !isAdmin {
      c.Redirect(http.StatusTemporaryRedirect, "/posts")
      return
    }
    c.HTML(http.StatusOK, "new-post", server.GetValues("new-post", c, gin.H{"date": time.Now().Format("January 02, 2006")}))
  })

  // server.Router.GET("/settings", func (c *gin.Context) {
  //   _, isAdmin := server.CheckContext(c)
  //   if !isAdmin {
  //     c.Redirect(http.StatusTemporaryRedirect, "/home")
  //     return
  //   }
  //   settings, err := server.Database.FetchSettings()
  //   if err != nil {
  //     c.Redirect(http.StatusTemporaryRedirect, "/home")
  //     return
  //   }
  //   c.HTML(http.StatusOK, "settings", server.GetValues("settings", c, gin.H{"settings": settings}))
  // })

  server.Router.GET("/login", func (c *gin.Context) {
    _, isAdmin := server.CheckContext(c)
    allowSignup := server.Database.GetSetting("allow_public_signup").(bool) || isAdmin
    c.HTML(http.StatusOK, "login", server.GetValues("login", c, gin.H{"allowSignup": allowSignup}))
  })

  server.Router.GET("/signup", func (c *gin.Context) {
    if !server.Database.GetSetting("allow_public_signup").(bool) {
      _, isAdmin := server.CheckContext(c)
      if !isAdmin {
        c.Redirect(http.StatusTemporaryRedirect, "/login")
        return
      }
    }
    c.HTML(http.StatusOK, "signup", server.GetValues("signup", c, gin.H{}))
  })

  server.Router.GET("/logout", func (c *gin.Context) {
    redirect := c.DefaultQuery("redirect", "/")
    RemoveContextLogin(c)
    c.Redirect(http.StatusTemporaryRedirect, redirect)
  })
}
