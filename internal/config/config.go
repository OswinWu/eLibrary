package config

import "gopkg.in/yaml.v3"

type Config struct {
	DB DBConfig `yaml:"db"`
}

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func LoadConfig(path string) (*Config, error) {
	var config Config
	err := yaml.Unmarshal([]byte(path), &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
