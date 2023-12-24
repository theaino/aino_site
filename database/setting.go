package database

import (
	"errors"
	"log"
	"strconv"

	"gorm.io/gorm"
)

type SettingType string

const (
  Int SettingType = "int"
  Str SettingType = "str"
  Bool SettingType = "bool"
)

type Setting struct {
  gorm.Model
  SettingKey string
  Type SettingType
  DefaultValue string
  Value string
}

func (settingType SettingType) Parse(value string) (interface{}, error) {
  switch settingType {
  case Int:
    result, err := strconv.Atoi(value)
    return result, err
  case Str:
    return value, nil
  case Bool: 
    result, err := strconv.ParseBool(value)
    return result, err
  }
  return nil, errors.New("Type not implemented")
}

func (connection *Connection) FetchSetting(key string) (Setting, error) {
  var setting Setting
  result := connection.Database.Where("setting_key = ?", key).First(&setting)
  return setting, result.Error
}

func (connection *Connection) GetSettingSafe(key string) (interface{}, error) {
  setting, err := connection.FetchSetting(key)
  if err != nil {
    return nil, err
  }
  value, err := setting.Type.Parse(setting.Value)
  return value, err
}

func (connection *Connection) GetSetting(key string) interface{} {
  value, err := connection.GetSettingSafe(key)
  if err != nil {
    log.Panic(err)
  }
  return value
}

func (connection *Connection) FetchSettings() ([]Setting, error) {
  var settings []Setting
  result := connection.Database.Find(&settings)
  return settings, result.Error
}

func (connection *Connection) SetSetting(key string, value string) error {
  setting, err := connection.FetchSetting(key)
  if err != nil {
    return err
  }
  settingType := setting.Type
  _, err = settingType.Parse(value)
  if err != nil {
    return err
  }
  setting.Value = value
  result := connection.Database.Save(&setting)
  return result.Error
}

