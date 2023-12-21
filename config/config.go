package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
  MysqlDsn string `yaml:"mysql_dsn"`
  AdminPassword string `yaml:"admin_password"`
  SessionSecret string `yaml:"session_secret"`
  SessionAge int `yaml:"session_age"`
  Release bool `yaml:"release"`
  Tls bool `yaml:"tls"`
  CertPath string `yaml:"cert_path"`
  PrivateKeyPath string `yaml:"private_key_path"`
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

