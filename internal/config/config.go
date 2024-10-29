package config

import (
	// "log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	API struct {
		Source struct {
			URL string `yaml:"url"`
		}
		Target struct {
			URL string `yaml:"url"`
		}
	}
	Sync struct {
		Interval int `yaml:"interval"`
	}
}

func LoadConfig() (*Config, error) {
	file, err := os.Open("config/config.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	config := &Config{}
	if err := yaml.NewDecoder(file).Decode(config); err != nil {
		return nil, err
	}
	return config, nil
}
