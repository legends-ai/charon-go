package config

import (
	"flag"
	"log"
	"time"
)

// AppConfig represents the application config
type AppConfig struct {
	Port        int
	MonitorPort int

	// RiotKey is the Riot API key.
	RiotKey string

	// MaxRate is maximum rate in which requests can be fired.
	MaxRate time.Duration

	// BufferTime is the buffer time given to retries
	BufferTime time.Duration

	// DefaultRetryAfter is the default retry time when getting a 429 from Riot
	DefaultRetryAfter time.Duration
}

var (
	port = flag.Int("port", 5609,
		"[default 5609] Charon Port.")
	monitorPort = flag.Int("monitor_port", 5610,
		"[default 5610] Charon Monitor Port.")
	riotKey = flag.String("riot_key", "",
		"[required] Riot API Key")
	maxRate = flag.Duration("max_rate", time.Duration(1200*time.Millisecond),
		"[default 1200ms] Max request rate to Riot API.")
	bufferTime = flag.Duration("buffer_time", time.Duration(500*time.Millisecond),
		"[default 500ms] Buffer time for Riot API retries when given a Retry-After header.")
	defaultRetryAfter = flag.Duration("default_retry_after", time.Duration(time.Second),
		"[deafult 1000ms] Default retry time if no Retry-After header is given.")
)

// Initialize initializes the configuration from env vars
func Initialize() *AppConfig {
	flag.Parse()
	if *riotKey == "" {
		log.Fatalf("Must pass Riot Key via command line flags")
	}
	return &AppConfig{
		Port:              *port,
		MonitorPort:       *monitorPort,
		RiotKey:           *riotKey,
		MaxRate:           *maxRate,
		BufferTime:        *bufferTime,
		DefaultRetryAfter: *defaultRetryAfter,
	}
}
