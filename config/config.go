package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

var C Config

type Config struct {
	APP string `yaml:"app"`
	Log Log    `yaml:"log"`
}

func LoadFile(filename string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(content, &C); err != nil {
		return err
	}
	return nil
}
