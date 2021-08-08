package cmd

import (
	"log"
	"os"

	"github.com/1995parham/mqtt-bench/internal/cmd/bench"
	"github.com/1995parham/mqtt-bench/internal/logger"
	"github.com/1995parham/mqtt-bench/internal/option"
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

	var options option.Options

	root.Flags().StringVarP(&options.Broker, "broker", "b",
		"tcp://127.0.0.1:1883", "mqtt broker e.g. tcp://127.0.0.1:1883")
	root.Flags().BoolVarP(&options.Retain, "retain", "r", false, "retain")
	root.Flags().IntVarP(&options.Clients, "clients", "c",
		option.DefaultClients, "number of simultaneous clients")
	root.Flags().IntVarP(&options.Count, "count", "t",
		option.DefaultCounts, "number of send messages")
	root.Flags().StringVarP(&level, "level", "l", "info", "set the logger level")

	logger := logger.New(level)

	bench.Register(root, options, logger)

	if err := root.Execute(); err != nil {
		log.Printf("failed to execute root command: %s", err.Error())
		os.Exit(ExitFailure)
	}
}
