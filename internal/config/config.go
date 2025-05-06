package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBConnection string `json:"db_url"`
	UserName     string `json:"current_user_name"`
}

// Read reads the configuration from the JSON file
func Read() (Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return Config{}, fmt.Errorf("failed to get path to home dir: %v", err)
	}

	configPath := filepath.Join(homeDir, configFileName)
	file, err := os.Open(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("failed to open config file: %v", err)
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return Config{}, fmt.Errorf("failed to decode config JSON: %v", err)
	}

	return cfg, nil
}

// SetUser sets the current user name and writes the updated config to the file
func (c *Config) SetUser(newUser string) error {
	c.UserName = newUser

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get path to home dir: %v", err)
	}

	configPath := filepath.Join(homeDir, configFileName)
	file, err := os.Create(configPath) // Overwrites the existing file
	if err != nil {
		return fmt.Errorf("failed to open config file for writing: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Prettify JSON
	if err := encoder.Encode(c); err != nil {
		return fmt.Errorf("failed to encode config to JSON: %v", err)
	}

	return nil
}
