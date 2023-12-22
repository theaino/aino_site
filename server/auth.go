package server

import (
	"errors"
	"strconv"

	"aino-spring.com/aino_site/database"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(password string) (string, error) {
  hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
  return string(hash), err
}

func CheckPassword(password, hash string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
  return err == nil
}

func CheckAuth(db *database.Connection, email, password string) (bool, bool) {  // Returns (isAuthed; isAdmin)
  id, err := db.FetchUserByEmail(email)
  if err != nil {
    return false, false
  }
  passwordHash, err := db.FetchUserPassword(id)
  if err != nil {
    return false, false
  }
  isAuthed := CheckPassword(password, passwordHash)
  if !isAuthed {
    return false, false
  }
  admin, err := db.FetchUserIsAdmin(id)
  if err != nil {
    return true, false
  }
  return true, admin
}

func NewUser(db *database.Connection, email, name, password string) (string, error) {
  _, err := db.FetchUserByEmail(email)
  if err == nil {
    return "", errors.New("User with this email already exists")
  }
  passwordHash, err := GenerateHash(password)
  if err != nil {
    return "", err
  }
  id, err := db.NewUser(email, name, passwordHash)
  if err != nil {
    return "", err
  }
  return strconv.Itoa(int(id)), nil
}

func RemoveContextLogin(c *gin.Context) {
  session := sessions.Default(c)
  session.Clear()

  session.Options(sessions.Options{MaxAge: -1})
  session.Save()
}

func (server *Server) CheckContext(c *gin.Context) (bool, bool) {
  queryEmail := c.Query("email")
  queryPassword := c.Query("password")

  session := sessions.Default(c)
  sessionEmail := session.Get("email")
  sessionPassword := session.Get("password")

  email := ""
  password := ""
  if sessionEmail != nil {
    email = sessionEmail.(string)
  }
  if queryEmail != "" {
    email = queryEmail
  }
  if sessionPassword != nil {
    password = sessionPassword.(string)
  }
  if queryPassword != "" {
    password = queryPassword
  }

  if email == "" || password == "" {
    return false, false
  }

  isAuthed, isAdmin := CheckAuth(server.Database, email, password)
  session.Set("email", email)
  session.Set("password", password)
  session.Save()
  return isAuthed, isAdmin
}

