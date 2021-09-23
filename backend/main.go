package main

import (
	"flag"
	"fmt"
	"os"
	"webexp/internal/configs"
	"webexp/internal/server/http"
)

type arguments struct {
	configPath string
}

func newArguments() *arguments {
	var configPath string
	flag.StringVar(&configPath, "config-path", "webexp.toml", "specify path to configuration file")
	flag.Parse()

	return &arguments{configPath: configPath}
}

func main() {
	args := newArguments()
	config, err := configs.NewConfigLoader(args.configPath).Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load config: %s\n", err.Error())
		os.Exit(111)
	}

	engine := http.Engine()
	engine.Run(config.Server.BindAddress())
}
