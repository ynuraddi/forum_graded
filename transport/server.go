package transport

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"graded/config"
	"graded/logger"
	"graded/service"
)

type Server struct {
	http.Server

	config *config.Config
	logger *logger.Logger

	handler http.Handler
	service *service.Manager
}

func Init(cfg *config.Config, lg *logger.Logger, srv *service.Manager) *Server {
	return &Server{
		config: cfg,
		logger: lg,

		handler: initDefaultRouter(),
		service: srv,
	}
}

func (s *Server) Run(ctx context.Context) error {
	go func() {
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalln(err)
		}
	}()

	<-ctx.Done()

	shutdownContext, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(shutdownContext); err != nil {
		log.Fatalf("Server shutdown failed: %s\n", err.Error())
		return err
	}

	return nil
}
