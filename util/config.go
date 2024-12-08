package util

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	LogLevel string
	LogMethod string
}

func defaultConfig() *Config {
	return &Config{
		LogLevel: "info",
		LogMethod: "console",
	}
}

func failedToLoadConfig(err error) {
	fmt.Printf("failed to load config, using default config. error: %s\n", err)
}

func LoadConfig() *Config {
	data, err := os.ReadFile("config.toml")

	if err != nil {
		failedToLoadConfig(err)
		return defaultConfig()
	}

	var config *Config

	// parse data
	_, err = toml.Decode(string(data), &config)

	if err != nil {
		failedToLoadConfig(err)
		return defaultConfig()
	}

	return config
}
