package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Config represents the CLI configuration
type Config struct {
	DefaultServer    string              `mapstructure:"default_server"`
	DefaultWorkspace string              `mapstructure:"default_workspace"`
	Workspaces       map[string]Workspace `mapstructure:"workspaces"`
	KeysDirectory    string              `mapstructure:"keys_directory"`
}

// Workspace configuration
type Workspace struct {
	Name      string `mapstructure:"name"`
	Server    string `mapstructure:"server"`
	KeyFile   string `mapstructure:"key_file"`
	KMSKeyID  string `mapstructure:"kms_key_id"`
}

// Load configuration from file
func Load(cfgFile string) (*Config, error) {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}

		viper.AddConfigPath(filepath.Join(home, ".fedmcp"))
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	viper.SetEnvPrefix("FEDMCP")
	viper.AutomaticEnv()

	// Set defaults
	viper.SetDefault("keys_directory", filepath.Join(os.Getenv("HOME"), ".fedmcp", "keys"))
	viper.SetDefault("default_server", "https://fedmcp.agency.gov")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

// Save configuration to file
func (c *Config) Save() error {
	return viper.WriteConfig()
}

// GetWorkspace returns a workspace configuration
func (c *Config) GetWorkspace(name string) (*Workspace, error) {
	if name == "" {
		name = c.DefaultWorkspace
	}
	
	ws, ok := c.Workspaces[name]
	if !ok {
		return nil, fmt.Errorf("workspace %q not found", name)
	}
	
	return &ws, nil
}
