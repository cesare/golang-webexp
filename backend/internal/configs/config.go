package configs

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	bind string
	port uint16
}

func (c *ServerConfig) BindAddress() string {
	return fmt.Sprintf("%s:%d", c.bind, c.port)
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
	err = decoder.Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse configuration file %s: %s", cl.path, err.Error())
	}

	return &config, nil
}
