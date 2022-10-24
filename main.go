package main

import (
	"fmt"
	"gateway/internal/config"
	"gateway/internal/web"
	"os"
	"path/filepath"
)

const (
	DEFAULT_PORT = "3000"
	DEFUALT_MODE = "SIMPLE"
)

func init() {
	prepare := filepath.Join(".", "data")
	os.MkdirAll(prepare, os.ModePerm)
}

func main() {
	var appPort = config.SetDataFromEnv("PORT", DEFAULT_PORT)
	var appMode = config.SetDataFromEnv("PORT", DEFUALT_MODE)
	fmt.Println("Web Server Port:", appPort)
	fmt.Println("Web Server Mode:", appMode)
	web.API(appPort, appMode)
}
