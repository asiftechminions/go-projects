package config

import (
	"log"

	"github.com/spf13/viper"
)

type GitHubConfig struct {
	Token       string `mapstructure:"token"`
	Owner       string `mapstructure:"owner"`
	Repo        string `mapstructure:"repo"`
	Branch      string `mapstructure:"branch"`
	DownloadDir string `mapstructure:"download_dir"`
}

type Config struct {
	GitHub GitHubConfig `mapstructure:"github"`
}

func LoadConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to decode config: %v", err)
	}
	return &config
}
