package bench

import (
	"fmt"

	"github.com/1995parham/mqtt-bench/internal/option"
	"github.com/1995parham/mqtt-bench/internal/publish"
	"github.com/1995parham/mqtt-bench/internal/subscribe"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func main(opts option.Options, logger *zap.Logger) {
	clients := make([]mqtt.Client, opts.Clients)

	for i := range clients {
		options := mqtt.NewClientOptions()

		options.SetClientID(fmt.Sprintf("mqtt-bench-%d", i))
		options.AddBroker(opts.Broker)

		if opts.Username != "" {
			options.SetUsername(opts.Username)
		}

		if opts.Password != "" {
			options.SetUsername(opts.Username)
		}

		clients[i] = mqtt.NewClient(options)

		if token := clients[i].Connect(); token.Wait() && token.Error() != nil {
			logger.Fatal("cannot connect", zap.String("broker", opts.Broker), zap.Error(token.Error()))
		}
	}

	s := subscribe.New(logger.Named("subscribe"), opts)
	s.Subscribe(clients)

	p := publish.New(logger.Named("publish"), opts)
	p.Publish(clients)
}

// Register benchmark command.
func Register(root *cobra.Command, logger *zap.Logger) {
	var opts option.Options

	// nolint: exhaustivestruct
	cmd :=
		&cobra.Command{
			Use:   "bench",
			Short: "",
			Run: func(cmd *cobra.Command, args []string) {
				pterm.Info.Printf("loaded configuration %+v\n", opts)

				main(opts, logger)
			},
		}

	cmd.Flags().StringVarP(&opts.Broker, "broker", "b",
		"tcp://127.0.0.1:1883", "mqtt broker e.g. tcp://127.0.0.1:1883")
	cmd.Flags().BoolVarP(&opts.Retain, "retain", "r", false, "retain")
	cmd.Flags().IntVarP(&opts.Clients, "clients", "c",
		option.DefaultClients, "number of simultaneous clients")
	cmd.Flags().IntVarP(&opts.Count, "count", "n",
		option.DefaultCounts, "number of send messages, use 0 for infinite number of messages")
	cmd.Flags().StringVarP(&opts.Username, "username", "u",
		"", "username for connecting to mqtt broker")
	cmd.Flags().StringVarP(&opts.Password, "password", "p",
		"", "password for connecting to mqtt broker")

	root.AddCommand(cmd)
}
