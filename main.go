package main

import (
	"fmt"
	"gateway/internal/config"
	"gateway/internal/web"
	"os"
	"path/filepath"
	"strings"
)

const (
	DEFAULT_PORT = "8080"
	DEFUALT_MODE = "NORMAL"
)

var Version = "Dev"

func init() {
	prepare := filepath.Join(".", "data")
	os.MkdirAll(prepare, os.ModePerm)
}

func main() {
	var appPort = config.SetDataFromEnv("PORT", DEFAULT_PORT)
	var appMode = config.SetDataFromEnv("MODE", DEFUALT_MODE)
	// review enabled by default
	if strings.ToUpper(appMode) != "SIMPLE" {
		appMode = DEFUALT_MODE
	}
	fmt.Printf("running soulteary/remote-hosts-server %s\n", Version)
	fmt.Println("Web Server Port:", appPort)
	fmt.Println("Web Server Mode:", appMode)
	web.API(appPort, appMode, Version)
}
