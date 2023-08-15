package app

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gitlab.com/zapirus/shortener/config"
)

//type APIServer struct {
//	config *config.Config
//	router http.Server
//}
//
//func New(config *config.Config) *APIServer {
//	return &APIServer{
//		config: config,
//		router: http.Server{}
//	}
//}

type APIServer struct {
	config *config.Config
	router *http.ServeMux
	server *http.Server
}

func New(config *config.Config) *APIServer {
	router := http.NewServeMux()
	server := &http.Server{
		Addr:    config.HTTPAddr,
		Handler: router,
	}

	return &APIServer{
		config: config,
		router: router,
		server: server,
	}
}

func (s *APIServer) Run() {
	srv := &http.Server{
		Addr:    s.config.HTTPAddr,
		Handler: s.router,
	}

	s.confRouter()
	log.Printf("Завелись на порту %s", s.config.HTTPAddr)
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint
		ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalln(err)
		}
	}()
	if err := srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalln(err)
		}
	}
	log.Println("Всего доброго!")
}
