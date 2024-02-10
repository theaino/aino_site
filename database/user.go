package database

import (
	"strconv"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	IsAdmin  bool
	Verified bool
}

func (connection *Connection) FetchUsers() ([]User, error) {
	var users []User
	result := connection.Database.Find(&users)
	return users, result.Error
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

func (connection *Connection) FetchUserEmail(id string) (string, error) {
	var user User
	result := connection.Database.First(&user, id)
	if result.Error != nil {
		return "", result.Error
	}
	return user.Email, nil
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

func (connection *Connection) FetchUserVerified(id string) (bool, error) {
	var user User
	result := connection.Database.First(&user, id)
	if result.Error != nil {
		return false, result.Error
	}
	return user.IsAdmin, nil
}

func (connection *Connection) NewUser(email, name, password string) (uint, error) {
	user := User{Email: email, Name: name, Password: password, IsAdmin: false, Verified: false}
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

func (connection *Connection) SetUserName(email, name string) error {
	id, err := connection.FetchUserByEmail(email)
	if err != nil {
		return err
	}
	var user User
	result := connection.Database.First(&user, id)
	if result.Error != nil {
		return result.Error
	}
	user.Name = name
	result = connection.Database.Save(&user)
	return result.Error
}

func (connection *Connection) SetUserPassword(email, password string) error {
	id, err := connection.FetchUserByEmail(email)
	if err != nil {
		return err
	}
	var user User
	result := connection.Database.First(&user, id)
	if result.Error != nil {
		return result.Error
	}
	user.Password = password
	result = connection.Database.Save(&user)
	return result.Error
}

func (connection *Connection) SetUserEmail(email, newEmail string) error {
	id, err := connection.FetchUserByEmail(email)
	if err != nil {
		return err
	}
	var user User
	result := connection.Database.First(&user, id)
	if result.Error != nil {
		return result.Error
	}
	user.Email = newEmail
	result = connection.Database.Save(&user)
	return result.Error
}

func (connection *Connection) SetUserIsAdmin(email string, isAdmin bool) error {
	id, err := connection.FetchUserByEmail(email)
	if err != nil {
		return err
	}
	var user User
	result := connection.Database.First(&user, id)
	if result.Error != nil {
		return result.Error
	}
	user.IsAdmin = isAdmin
	result = connection.Database.Save(&user)
	return result.Error
}

func (connection *Connection) SetUserVerified(email string, verified bool) error {
	id, err := connection.FetchUserByEmail(email)
	if err != nil {
		return err
	}
	var user User
	result := connection.Database.First(&user, id)
	if result.Error != nil {
		return result.Error
	}
	user.Verified = verified
	result = connection.Database.Save(&user)
	return result.Error
}

func (connection *Connection) DeleteUser(email string) error {
	id, err := connection.FetchUserByEmail(email)
	if err != nil {
		return err
	}

	result := connection.Database.Delete(&User{}, id)
	return result.Error
}
