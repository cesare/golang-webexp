package configs

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Server   ServerConfig
	Auth     AuthConfig
	Frontend FrontendConfig
}

type ServerConfig struct {
	Bind string
	Port uint16
}

func (c *ServerConfig) BindAddress() string {
	return fmt.Sprintf("%s:%d", c.Bind, c.Port)
}

type AuthConfig struct {
	ClientId     string
	ClientSecret string
}

type FrontendConfig struct {
	BaseUri      string
	CallbackPath string
}

type ConfigLoader struct {
	path string
}

func NewConfigLoader(path string) *ConfigLoader {
	return &ConfigLoader{path: path}
}

func (cl *ConfigLoader) Load() (*Config, error) {
	f, err := os.Open(cl.path)
	if err != nil {
		return nil, fmt.Errorf("failed to open configuration file %s: %s", cl.path, err.Error())
	}
	defer f.Close()

	var config Config
	decoder := toml.NewDecoder(f)
	decoder.SetStrict(true)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse configuration file %s: %s", cl.path, err.Error())
	}

	return &config, nil
}
