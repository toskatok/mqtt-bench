package bench

import (
	"fmt"

	"github.com/1995parham/mqtt-bench/internal/option"
	"github.com/1995parham/mqtt-bench/internal/publish"
	mqtt "github.com/eclipse/paho.mqtt.golang"
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

	p := publish.New(logger, opts)
	p.Publish(clients)
}

// Register benchmark command.
func Register(root *cobra.Command, opts option.Options, logger *zap.Logger) {
	// nolint: exhaustivestruct
	cmd :=
		&cobra.Command{
			Use:   "bench",
			Short: "",
			Run: func(cmd *cobra.Command, args []string) {
				main(opts, logger)
			},
		}
	root.AddCommand(cmd)
}
