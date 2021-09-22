package main

import (
	"fmt"
	"os"
	"webexp/internal/configs"
	"webexp/internal/server/http"
)

func main() {
	config, err := configs.NewConfigLoader("./webexp.toml").Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load config: %s\n", err.Error())
		os.Exit(111)
	}

	engine := http.Engine()
	engine.Run(config.Server.BindAddress())
}
