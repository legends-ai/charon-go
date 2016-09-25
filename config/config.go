package config

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
)

// AppConfig represents the application config
type AppConfig struct {
	Port        int `required:"true" default:"5609"`
	MonitorPort int `required:"true" default:"5610"`

	// APIKey is the Riot API key.
	APIKey string `required:"true"`

	// MaxRate is maximum rate in which requests can be fired.
	MaxRate time.Duration `required:"true" default:"1200ms"`
}

// Initialize initializes the configuration from env vars
func Initialize() *AppConfig {
	var cfg AppConfig
	err := envconfig.Process("charon", &cfg)
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}
	return &cfg
}
