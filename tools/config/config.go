package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

var configPath = "../configs/config.yaml"

type Config struct {
	Server  Server `yaml:"server"`
	ParsURL string `yaml:"pars_url"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func NewConfig() (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}
