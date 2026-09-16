package config

import (
	"fmt"
	"os"
	"path/filepath"
	"gopkg.in/yaml.v3"
)


type Config struct {
	Database DatabaseConfig `yaml:"database"`
}

type DatabaseConfig struct{
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	User string `yaml:"user"`
	Password string `yaml:"password"`
	DBName string `yaml:"dbname"`
	SSLMode string `yaml:"sslmode"`

	URL string `yaml:"-"`

}
func Load() (*Config, error){
	configpath:=filepath.Join("internal", "config", "config.yaml")
	file, err := os.ReadFile(configpath)
	if err != nil{
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(file, &cfg); err != nil{
		return nil, err
	}
	cfg.Database.URL = fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBName,
		cfg.Database.SSLMode,
	)
	return &cfg, nil
}