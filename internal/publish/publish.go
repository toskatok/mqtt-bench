package publish

import (
	"fmt"
	"sync"
	"time"

	"github.com/1995parham/mqtt-bench/internal/option"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/zap"
)

type Publisher struct {
	logger *zap.Logger
	opts   option.Options
}

func New(logger *zap.Logger, opts option.Options) *Publisher {
	return &Publisher{
		logger: logger,
		opts:   opts,
	}
}

// Publish start publishing on goroutines.
func (p *Publisher) Publish(clients []mqtt.Client) {
	wg := new(sync.WaitGroup)

	for _, client := range clients {
		wg.Add(1)

		p.publish(client, wg)
	}

	wg.Wait()
}

// publish is called on a goroutine to publish messages.
func (p *Publisher) publish(client mqtt.Client, wg *sync.WaitGroup) {
	for index := 0; index < p.opts.Count; index++ {
		options := client.OptionsReader()

		topic := fmt.Sprintf("%s/%s", option.DefaultTopic, options.ClientID())

		message := make([]byte, p.opts.MessageSize)

		if token := client.Publish(topic, p.opts.Qos, p.opts.Retain, message); token.Wait() && token.Error() != nil {
			p.logger.Error("publish failed", zap.String("topic", topic), zap.Error(token.Error()))
		}

		if index%10 == 0 {
			p.logger.Info("publish message", zap.Int("count", index), zap.String("topic", topic))
		}

		if p.opts.IntervalTime > 0 {
			time.Sleep(time.Duration(p.opts.IntervalTime) * time.Millisecond)
		}
	}

	wg.Done()
}
