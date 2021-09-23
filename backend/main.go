package main

import (
	"flag"
	"fmt"
	"os"
	"webexp/internal/configs"
	"webexp/internal/server/http"
)

type Arguments struct {
	ConfigPath string
}

func NewArguments() *Arguments {
	var configPath string
	flag.StringVar(&configPath, "config-path", "webexp.toml", "specify path to configuration file")
	flag.Parse()

	return &Arguments{ConfigPath: configPath}
}

func main() {
	args := NewArguments()
	config, err := configs.NewConfigLoader(args.ConfigPath).Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load config: %s\n", err.Error())
		os.Exit(111)
	}

	engine := http.Engine()
	engine.Run(config.Server.BindAddress())
}
