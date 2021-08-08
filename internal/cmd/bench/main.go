package bench

import (
	"fmt"
	"sync"

	"github.com/1995parham/mqtt-bench/internal/option"
	"github.com/1995parham/mqtt-bench/internal/publish"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func main(opts option.Options) {
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
			pterm.Fatal.Printf("cannot connect to %s %s", opts.Broker, token.Error().Error())
		}
	}

	wg := new(sync.WaitGroup)

	for _, client := range clients {
		wg.Add(1)

		publish.Publish(client, opts, wg)
	}

	wg.Wait()
}

// Register benchmark command.
func Register(root *cobra.Command, opts option.Options) {
	// nolint: exhaustivestruct
	cmd :=
		&cobra.Command{
			Use:   "bench",
			Short: "",
			Run: func(cmd *cobra.Command, args []string) {
				main(opts)
			},
		}
	root.AddCommand(cmd)
}
