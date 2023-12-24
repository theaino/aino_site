package database

import (
	"errors"
	"log"
	"strconv"
)

type Setting struct {
  SettingKey string `gorm:"primaryKey"`
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
  value, err := SettingPresets[setting.SettingKey].Type.Parse(setting.Value)
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
  settingType := SettingPresets[setting.SettingKey].Type
  _, err = settingType.Parse(value)
  if err != nil {
    return err
  }
  setting.Value = value
  result := connection.Database.Save(&setting)
  return result.Error
}

