package recalapi

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"

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
	SystemsPath  string
	BiosPath     string
	Gamelist     string
	RecalboxConf string
}

// Endpoint represents the differents path for the API
type Endpoint struct {
	SystemsEndpoint string
	BiosEndpoint    string
}

type RecalboxConf struct {
	Options []RecalboxOption
}
type RecalboxOption struct {
	Enable bool
	Option map[string]string
}

// LoadConfig load the configuration file
func LoadConfig(mode string) *Configuration {
	// Load config in ./config but only for dev
	// In recalbox should be in /etc/recalbox-api/config.toml
	path := "/etc/recalbox-api/config.toml"
	if mode == "debug" {
		path = "./recalapi/config.toml"
	}

	config := Configuration{}
	_, err := toml.DecodeFile(path, &config)

	if err != nil {
		fmt.Println("error:", err)
	}

	//config.RecalboxConfParser()
	return &config
}

func (config *Configuration) RecalboxConfParser() {
	buf := bytes.NewBuffer(nil)

	regex := `(;?)([A-Za-z0-9-.]*)=([a-zA-Z0-9_\s/,:.]*)`
	f, _ := os.Open(config.Fs.RecalboxConf)
	io.Copy(buf, f)

	r, err := regexp.Compile(regex)
	if err != nil {
		println(err)
	}
	println(buf.String())
	str := r.FindAllString(buf.String(), -1)
	for _, v := range str {
		if v != "" {
			println(v)
		}
	}

}
