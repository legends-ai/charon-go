package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	_ "net/http/pprof"

	"github.com/Sirupsen/logrus"
	"github.com/simplyianm/inject"
	"google.golang.org/grpc"

	"github.com/asunaio/charon/config"
	"github.com/asunaio/charon/lib"
	"github.com/asunaio/charon/server"

	apb "github.com/asunaio/charon/gen-go/asuna"
)

func main() {
	injector := lib.NewInjector()
	_, err := injector.Invoke(initServer)
	if err != nil {
		log.Fatalf("Could not invoke init: %v", err)
	}
}

func initServer(injector inject.Injector, logger *logrus.Logger, config *config.AppConfig) {
	// Listen on port
	port := fmt.Sprintf(":%d", config.Port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Setup gRPC server
	s := grpc.NewServer()
	serv := &server.Server{}

	_, err = injector.ApplyMap(serv)
	if err != nil {
		logger.Fatalf("Could not inject server: %v", err)
	}

	go func() {
		http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
		monitorPort := fmt.Sprintf(":%d", config.MonitorPort)
		logger.Infof("Monitoring listening on %s", monitorPort)
		http.ListenAndServe(monitorPort, nil)
	}()

	apb.RegisterCharonServer(s, serv)
	logger.Infof("Listening on %s", port)
	s.Serve(lis)
}
