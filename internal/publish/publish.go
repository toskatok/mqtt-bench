package publish

import (
	"fmt"
	"sync"
	"time"

	"github.com/1995parham/mqtt-bench/internal/option"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/pterm/pterm"
)

// Publish function calls on a goroutine to publish messages.
func Publish(client mqtt.Client, opts option.Options, wg *sync.WaitGroup) {
	for index := 0; index < opts.Count; index++ {
		options := client.OptionsReader()

		topic := fmt.Sprintf("%s/%s", option.DefaultTopic, options.ClientID())

		message := make([]byte, opts.MessageSize)

		if token := client.Publish(topic, opts.Qos, opts.Retain, message); token.Wait() && token.Error() != nil {
			pterm.Error.Printf("publish on %s failed %s", topic, token.Error().Error())
		}

		if index%10 == 0 {
			pterm.Info.Printf("publish %d message on %s", index, topic)
		}

		if opts.IntervalTime > 0 {
			time.Sleep(time.Duration(opts.IntervalTime) * time.Millisecond)
		}
	}

	wg.Done()
}
