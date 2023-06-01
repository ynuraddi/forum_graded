package main

import (
	"context"
	"log"
	"os"
	"os/signal"
)

func main() {
	log.Fatalf("Service shutdown: %s\n", run())
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	gracefullyShutdown(cancel)

	return nil
}

func gracefullyShutdown(cancel context.CancelFunc) {
	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal)

	go func() {
		cancel()
	}()
}
