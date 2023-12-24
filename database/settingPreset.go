package database

type SettingType string

const (
  Int SettingType = "int"
  Str SettingType = "str"
  Bool SettingType = "bool"
)

type SettingPreset struct {
  Type SettingType
  DefaultValue string
}

var SettingPresets = map[string]SettingPreset{
  "allow_public_signup": {
    Bool,
    "true",
  },
  "session_secret": {
    Str,
    "b&qmhwrc",
  },
}

func (connection *Connection) ResetSettingPreset(key string, preset SettingPreset) error {
  setting := Setting{SettingKey: key, Value: preset.DefaultValue}
  result := connection.Database.Save(&setting)
  return result.Error
}

func (connection *Connection) SetupSettingPresets() error {
  for key, preset := range SettingPresets {
    var setting Setting
    result := connection.Database.Where("setting_key = ?", key).First(&setting)
    if result.Error != nil || setting.Value == "" {
      err := connection.ResetSettingPreset(key, preset)
      if err != nil {
        return err
      }
    }
  }
  return nil
}

