package util

import (
	"fmt"
	"os"

	"github.com/wirdos/resources"
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
	var config *Config

	// read from root for development initially
	data, err := os.ReadFile("config.toml")

	// if the file doesn't exist which is likely in general runtime; load bundled config
	if os.IsNotExist(err) {
		config, err = resources.LoadToml[Config]("config")

		if err == nil {
			return config
		}
	}

	if err == nil {
		config, err = resources.Deserialise[Config](data, resources.TOML)
	}

	if err != nil {
		failedToLoadConfig(err)
		return defaultConfig()
	}

	return config
}
