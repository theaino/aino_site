package database

import (
	"strconv"

	"gorm.io/gorm"
)

type User struct {
  gorm.Model
  Name string
  Email string
  Password string
  IsAdmin bool
}

func (connection *Connection) FetchUserByEmail(email string) (string, error) {
  var user User
  result := connection.Database.First(&user, "email = ?", email)
  if result.Error != nil {
    return "", result.Error
  }
  return strconv.Itoa(int(user.ID)), nil
}

func (connection *Connection) FetchUserByName(name string) (string, error) {
  var user User
  result := connection.Database.First(&user, "name = ?", name)
  if result.Error != nil {
    return "", result.Error
  }
  return strconv.Itoa(int(user.ID)), nil
}

func (connection *Connection) FetchUserName(id string) (string, error) {
  var user User
  result := connection.Database.First(&user, id)
  if result.Error != nil {
    return "", result.Error
  }
  return user.Name, nil
}

func (connection *Connection) FetchUserPassword(id string) (string, error) {
  var user User
  result := connection.Database.First(&user, id)
  if result.Error != nil {
    return "", result.Error
  }
  return user.Password, nil
}

func (connection *Connection) FetchUserIsAdmin(id string) (bool, error) {
  var user User
  result := connection.Database.First(&user, id)
  if result.Error != nil {
    return false, result.Error
  }
  return user.IsAdmin, nil
}

func (connection *Connection) NewUser(email, name, password string) (uint, error) {
  user := User{Email: email, Name: name, Password: password, IsAdmin: false}
  var count int64
  result := connection.Database.Model(&User{}).Count(&count)
  if result.Error != nil {
    return 0, result.Error
  }
  if count == 0 {
    user.IsAdmin = true
  }
  result = connection.Database.Create(&user)
  return user.ID, result.Error
}

