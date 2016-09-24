package lib

import (
	"github.com/Sirupsen/logrus"
	"github.com/asunaio/charon/config"
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

	return injector
}
