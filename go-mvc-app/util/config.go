package util

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Service struct {
		Port string `yaml:"port"`
	} `yaml:"sevice"`
	Database struct {
		Name     string `yaml:"name"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"database"`
}

var AppConfig *Config

func LoadConfig() {
	f, err := os.Open("config/config.yaml")
	if err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	cfg := &Config{}
	err = decoder.Decode(cfg)

	if err != nil {
		log.Fatalf("Error decoding config file: %v", err)
	}

	AppConfig = cfg
}
