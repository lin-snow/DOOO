// Setting up the configuration for the application
package config

import (
	"log"
	"os"

	yaml "gopkg.in/yaml.v3"
)

type Config struct {
	// Server configuration
	Server struct {
		Listen string `yaml:"listen"`
	} `yaml:"Server"`

	// Database configuration
	Database struct {
		DBtype   string `yaml:"dbtype"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		DBname   string `yaml:"dbname"`
	} `yaml:"Database"`

	// WebPath configuration
	WebPath struct {
		Dist string `yaml:"dist"`
	} `yaml:"WebPath"`
}

// LoadConfig loads the configuration from the config.yaml file.
func LoadConfig() Config {
	file, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		log.Println("Error reading the config file: ", err)
	}

	var config Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Println("Error unmarshalling the config file: ", err)
	}

	return config
}
