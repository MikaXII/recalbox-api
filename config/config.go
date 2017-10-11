package configuration

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

// Configuration represents the config file in toml
type Configuration struct {
	Version      string
	Fs           FileSystemInfo
	ListEndpoint Endpoint
}

// FileSystemInfo represents a system
type FileSystemInfo struct {
	SystemsPath string
	BiosPath    string
	Gamelist    string
}

// Endpoint represents the differents path for the API
type Endpoint struct {
	SystemsEndpoint string
	BiosEndpoint    string
}

// LoadConfig load the configuration file
func LoadConfig(mode string) *Configuration {
	// Load config in ./config but only for dev
	// In recalbox should be in /etc/recalbox-api/config.toml
	path := "/etc/recalbox-api/config.toml"
	if mode == "debug" {
		path = "./config/config.toml"
	}

	config := Configuration{}
	_, err := toml.DecodeFile(path, &config)

	if err != nil {
		fmt.Println("error:", err)
	}
	return &config

}
