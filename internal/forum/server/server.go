package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"forum/pkg/logging"
)

type Config struct {
	Host         string `json:"host"`
	Port         string `json:"port"`
	IdleTimeout  int    `json:"idleTimeout"`
	ReadTimeout  int    `json:"readTimeout"`
	WriteTimeout int    `json:"writeTimeout"`
	EndTimeLimit int    `json:"endTimeLimit"`
}

type Server struct {
	*http.Server
}

func NewServer(cfg *Config, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Handler:      handler,
		IdleTimeout:  time.Duration(cfg.IdleTimeout) * time.Second,
		ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
	}
}

func (s *Server) Serve(cfg Config) error {
	srv := http.Server{
		Addr:         fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Handler:      s.Handler,
		IdleTimeout:  s.IdleTimeout,
		ReadTimeout:  s.ReadTimeout,
		WriteTimeout: s.WriteTimeout,
	}

	logger := logging.GetLoggerInstance()
	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit

		logger.PrintInfo("caught signal:" + s.String())

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.EndTimeLimit)*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}

		logger.PrintInfo("completing background tasks")
		shutdownError <- nil
	}()

	logger.PrintInfo("starting server")

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		logger.PrintFatal(err)
	}

	err = <-shutdownError
	if err != nil {
		logger.PrintError(err)
	}

	return err
}
