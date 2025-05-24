package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config holds the application's configuration settings.
type Config struct {
	DatabaseURL string `mapstructure:"DATABASE_URL"`
	ServerPort  string `mapstructure:"SERVER_PORT"`
	// Add other configuration fields here
}

// LoadConfig reads configuration from a .env file and environment variables
// using Viper.
func LoadConfig() (*Config, error) {
	viper.AddConfigPath(".")    // Search for config in the current directory
	viper.SetConfigName(".env") // Look for .env file
	viper.SetConfigType("env")  // Set config type to env

	// Default values (optional)
	viper.SetDefault("SERVER_PORT", "8080")

	viper.AutomaticEnv() // Read from environment variables

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore this error if using env vars
			fmt.Println("Warning: .env file not found. Using environment variables.")
		} else {
			return nil, fmt.Errorf("failed to read configuration: %w", err)
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal configuration: %w", err)
	}

	return &cfg, nil
}
