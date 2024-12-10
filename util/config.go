package util

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
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
	// read from root for development initially
	data, err := os.ReadFile("config.toml")

	if err != nil {
		// if the file doesn't exist which is likely in general runtime; load bundled config
		if os.IsNotExist(err) {
			bundled, err := resources.LoadToml[Config]("config")

			if err == nil {
				return bundled
			}
		}

		failedToLoadConfig(err)
		return defaultConfig()
	}

	// TODO: refactor, we're doing this in `resources` anyway
	var config *Config

	// parse data
	_, err = toml.Decode(string(data), &config)

	if err != nil {
		failedToLoadConfig(err)
		return defaultConfig()
	}

	return config
}
