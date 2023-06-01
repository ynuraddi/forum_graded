package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"graded/config"
	"graded/logger"
	"graded/repository"
	"graded/service"
	"graded/transport"
)

func main() {
	log.Fatalf("Service shutdown: %s\n", run())
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	gracefullyShutdown(cancel)

	config := config.Init()
	logger, err := logger.Init(config)
	if err != nil {
		return err
	}

	repository, err := repository.Init(config, logger)
	if err != nil {
		return err
	}

	service, err := service.Init(config, logger, repository)
	if err != nil {
		return err
	}

	server := transport.Init(config, logger, service)

	return server.Run(ctx)
}

func gracefullyShutdown(cancel context.CancelFunc) {
	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal)

	go func() {
		cancel()
	}()
}
