package cmd

import (
	"log"
	"os"

	"github.com/1995parham/mqtt-bench/internal/cmd/bench"
	"github.com/1995parham/mqtt-bench/internal/logger"
	"github.com/spf13/cobra"
)

const (
	// ExitFailure status code.
	ExitFailure = 1
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// nolint: exhaustivestruct
	root := &cobra.Command{
		Use: "mqtt-bench",
	}

	var level string
	root.PersistentFlags().StringVarP(&level, "level", "l", "info", "set the logger level")
	logger := logger.New(level)

	bench.Register(root, logger)

	if err := root.Execute(); err != nil {
		log.Printf("failed to execute root command: %s", err.Error())
		os.Exit(ExitFailure)
	}
}
