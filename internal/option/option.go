package option

// Options contains the mqtt-bench options which are loaded from command line flags.
type Options struct {
	Qos               byte
	Retain            bool
	UseDefaultHandler bool
	Broker            string
	Topic             string
	Username          string
	Password          string
	Clients           int
	Count             int
	IntervalTime      int
	MessageSize       int
}

const (
	DefaultClients = 10
	DefaultCounts  = 1000

	DefaultTopic = "/mqtt-bench/benchmark"
)
