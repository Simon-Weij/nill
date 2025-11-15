package cli

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var path string

var rootCmd = &cobra.Command{
	Use: "nill",
}

func InitConfigPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	defaultPath := filepath.Join(homeDir, "nill", "config.yaml")

	rootCmd.Flags().StringVarP(
		&path,
		"path",
		"p",
		defaultPath,
		"Path to configuration file",
	)

	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}

	return path
}
