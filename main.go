package main

import (
	"fmt"
	"gateway/internal/config"
	"gateway/internal/web"
)

const (
	DEFAULT_PORT = "3000"
)

func main() {
	var appPort = config.SetDataFromEnv("GATE_APP_PORT", DEFAULT_PORT)
	fmt.Println("Web Server port:", DEFAULT_PORT)
	web.API(appPort)
}
