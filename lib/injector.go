package lib

import (
	"github.com/Sirupsen/logrus"
	"github.com/asunaio/charon/config"
	"github.com/asunaio/charon/metrics"
	"github.com/asunaio/charon/riot/client"
	"github.com/simplyianm/inject"
)

// NewInjector builds a new Injector.
func NewInjector() inject.Injector {
	injector := inject.New()
	injector.Map(injector)

	logger := logrus.New()
	injector.Map(logger)

	cfg := config.Initialize()
	injector.Map(cfg)

	metrics := &metrics.Metrics{}
	_, err := injector.ApplyMap(metrics)
	if err != nil {
		logger.Fatalf("Could not inject Metrics: %v", err)
	}
	go metrics.Start()

	_, err = injector.ApplyMap(riot.New())
	if err != nil {
		logger.Fatalf("Could not inject riot: %v", err)
	}

	return injector
}
