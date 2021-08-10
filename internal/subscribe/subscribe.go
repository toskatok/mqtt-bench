package subscribe

import (
	"fmt"

	"github.com/1995parham/mqtt-bench/internal/option"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/zap"
)

type Subscriber struct {
	logger *zap.Logger
	opts   option.Options
}

func New(logger *zap.Logger, opts option.Options) *Subscriber {
	return &Subscriber{
		logger: logger,
		opts:   opts,
	}
}

// Subscribe start subscribing with goroutines. please note that this isn't a blocking function.
func (s *Subscriber) Subscribe(clients []mqtt.Client) {
	for _, client := range clients {
		s.subscribe(client)
	}
}

// subscribe is called on a goroutine to subscribe messages.
func (s *Subscriber) subscribe(client mqtt.Client) {
	options := client.OptionsReader()

	topic := fmt.Sprintf("%s/%s", option.DefaultTopic, options.ClientID())

	if token := client.Subscribe(topic, s.opts.Qos, func(c mqtt.Client, m mqtt.Message) {
		options := c.OptionsReader()

		s.logger.Info("receive message", zap.String("client", options.ClientID()))
	}); token.Wait() && token.Error() != nil {
		s.logger.Error("subscription failed", zap.Error(token.Error()), zap.String("client", options.ClientID()))
	}
}
