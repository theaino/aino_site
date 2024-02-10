package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	MysqlDsn       string `yaml:"mysql_dsn"`
	Release        bool   `yaml:"release"`
	Tls            bool   `yaml:"tls"`
	CertPath       string `yaml:"cert_path"`
	PrivateKeyPath string `yaml:"private_key_path"`
	Address        string `yaml:"address"`
	VerifySalt     string `yaml:"verify_salt"`
	Email          string `yaml:"email"`
	EmailPassword  string `yaml:"email_password"`
	SMTPHost       string `yaml:"smtp_host"`
	SMTPPort       int    `yaml:"smtp_port"`
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
