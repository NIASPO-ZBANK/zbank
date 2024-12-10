package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"

	"users/internal/infrastructure/database"
)

type Config struct {
	Server server
}

type server struct {
	Addr string `yaml:"addr"`
	Port string `yaml:"port"`

	StaticData struct {
		Connection *database.Config `yaml:"db"`
	} `yaml:"oreshnik_data"`
}

func ReadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
