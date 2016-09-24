package lib

import (
	"github.com/Sirupsen/logrus"
	"github.com/asunaio/charon/config"
	"github.com/asunaio/charon/riot"
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

	_, err := injector.ApplyMap(riot.New())
	if err != nil {
		logger.Fatalf("Could not inject riot: %v", err)
	}

	return injector
}
