package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
  MysqlDsn string `yaml:"mysql_dsn"`
  AdminPath string `yaml:"admin_path"`
}

func LoadConfig(path string) (*Config, error) {
  contents, err := os.ReadFile(path)
  if err != nil {
    return nil, err
  }
  config := new(Config)
  err = yaml.Unmarshal(contents, config)
  return config, err
}

