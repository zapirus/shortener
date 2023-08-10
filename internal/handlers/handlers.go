package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	"gitlab.com/zapirus/shortener/internal/service"
)

type APIServer struct {
	config *Config
	router *mux.Router
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Run() {
	srv := &http.Server{
		Addr:    s.config.HTTPAddr,
		Handler: s.router,
	}

	s.confRouter()
	log.Printf("Завелись на порту %s", s.config.HTTPAddr)
	idConnClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		<-sigint
		ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalln(err)
		}
		close(idConnClosed)
	}()
	if err := srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalln(err)
		}
	}
	<-idConnClosed
	log.Println("Всего доброго!")
}

func (s *APIServer) confRouter() {
	s.router.HandleFunc("/hello", s.GetHello()).Methods("GET")
}

func (s *APIServer) GetHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		hello := service.Hello()
		if err := json.NewEncoder(w).Encode(hello); err != nil {
			log.Printf("Не удалось преаброзовать: %s", err)
			w.WriteHeader(http.StatusConflict)
			return
		}
	}

}
