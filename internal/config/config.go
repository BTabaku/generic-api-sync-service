package config

import (
	"bytes"
	"log"
	"os"
	"text/template"

	"github.com/joho/godotenv"
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
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Read the YAML file content
	file, err := os.ReadFile("config/config.yaml")
	if err != nil {
		return nil, err
	}

	// Replace environment variable placeholders
	tmpl, err := template.New("config").Parse(string(file))
	if err != nil {
		return nil, err
	}

	var resolvedConfig bytes.Buffer
	if err := tmpl.Execute(&resolvedConfig, getEnvMap()); err != nil {
		return nil, err
	}

	// Unmarshal the modified YAML content into the Config struct
	config := &Config{}
	if err := yaml.NewDecoder(&resolvedConfig).Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}

// getEnvMap returns a map of environment variables
func getEnvMap() map[string]string {
	envMap := make(map[string]string)
	for _, env := range os.Environ() {
		pair := bytes.SplitN([]byte(env), []byte{'='}, 2)
		envMap[string(pair[0])] = string(pair[1])
	}
	return envMap
}
