package main

import (
	"Simon-Weij/nill/src/cli"
	"Simon-Weij/nill/src/router"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	configPath := cli.InitConfigPath()

	_, err := ensureConfigPath(configPath)
	if err != nil {
		panic(err)
	}

	router.DefineRoutes()
}

func ensureConfigPath(path string) (*os.File, error) {
	if strings.HasSuffix(path, ".yml") || strings.HasSuffix(path, ".yaml") {
		dir := filepath.Dir(path)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, err
		}

		file, err := os.Create(path)
		if err != nil {
			return nil, err
		}

		return file, nil
	} else if filepath.Ext(path) == "" {
		if err := os.MkdirAll(path, 0755); err != nil {
			return nil, err
		}
		if file, err := os.Create(path + "/config.yaml"); err != nil {
			return file, nil
		}
	}
	return nil, fmt.Errorf("This file can't be used as configuration file")
}
