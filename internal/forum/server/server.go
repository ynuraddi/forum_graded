package server

import (
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Host         string `json:"host"`
	Port         string `json:"port"`
	IdleTimeout  int    `json:"idleTimeout"`
	ReadTimeout  int    `json:"readTimeout"`
	WriteTimeout int    `json:"writeTimeout"`
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

const endTimeLimit = 5 * time.Second

// func (s *server) Serve() error {
// 	srv := http.Server{
// 		Addr: fmt.Sprintf("%s:%s", s.host, s.port),
// 		Handler: s.handler,
// 		IdleTimeout: s.idleTimeout,
// 		ReadTimeout: s.readTimeout,
// 		WriteTimeout: s.writeTimeout,
// 	}

// 	shutdowmChan := make(chan error)

// 	go func ()  {
// 		quit := make(chan os.Signal, 1)
// 		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
// 		s := <- quit

// 		ctx, cancel := context.WithTimeout(context.Background(), endTimeLimit)
// 		defer cancel()

// 	}

// }
