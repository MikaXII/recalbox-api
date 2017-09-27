package configuration

import (
	"fmt"

	"gopkg.in/BurntSushi/toml.v0"
)

type Configuration struct {
	Version      string
	Fs           FileSystemInfo
	ListEndpoint Endpoint
}

type FileSystemInfo struct {
	SystemsPath string
	BiosPath    string
	Gamelist    string
}

type Endpoint struct {
	SystemsEndpoint string
	BiosEndpoint    string
}

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
