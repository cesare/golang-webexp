package main

import (
	"webexp/internal/server/http"
)

func main() {
	engine := http.Engine()
	engine.Run(":3000")
}
