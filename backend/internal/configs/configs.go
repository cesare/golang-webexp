package configs

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type b64string struct {
	bytes []byte
}

func (b64 *b64string) Bytes() []byte {
	return b64.bytes
}

func (b64 *b64string) UnmarshalText(text []byte) error {
	dst := make([]byte, len(text))
	_, err := base64.StdEncoding.Decode(dst, text)
	if err != nil {
		return err
	}

	b64.bytes = dst
	return nil
}

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	App      AppConfig
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

type DatabaseConfig struct {
	Host     string
	Port     uint16
	Database string
	User     string
	Password string
	UseSSL   bool
}

func (c *DatabaseConfig) ConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.User, c.Password, c.Host, c.Port, c.Database, c.sslMode())
}

func (c *DatabaseConfig) sslMode() string {
	if c.UseSSL {
		return "require"
	} else {
		return "disable"
	}
}

type AppConfig struct {
	SessionKey b64string
}

type AuthConfig struct {
	ClientId        string
	ClientSecret    string
	TokenSigningKey b64string
}

type FrontendConfig struct {
	BaseUri      string
	CallbackPath string
}

func (f *FrontendConfig) CallbackUri() string {
	return fmt.Sprintf("%s%s", f.BaseUri, f.CallbackPath)
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
