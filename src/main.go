package main

import (
	"Simon-Weij/nill/src/cli"
	"Simon-Weij/nill/src/router"
	_ "embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

//go:embed examples/config.yaml
var defaultConfig []byte

func main() {
	godotenv.Load()

	configPath := cli.InitConfigPath()

	if err := ensureConfigFileExists(configPath); err != nil {
		panic(err)
	}

	config, err := router.ParseRoutes(configPath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Loaded %d endpoints from config\n", len(config.Endpoints))

	router.DefineRoutes(config)
}

func ensureConfigFileExists(path string) error {
	if filepath.Ext(path) == "" {
		path = filepath.Join(path, "config.yaml")
	}

	if _, err := os.Stat(path); err == nil {
		return nil
	} else if !os.IsNotExist(err) {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}

	if err := os.WriteFile(path, defaultConfig, 0644); err != nil {
		return err
	}
	return nil
}
