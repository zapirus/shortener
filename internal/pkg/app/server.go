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

	"github.com/joho/godotenv"

	"gitlab.com/zapirus/shortener/config"
	"gitlab.com/zapirus/shortener/internal/handlers"
	"gitlab.com/zapirus/shortener/internal/repository"
	"gitlab.com/zapirus/shortener/internal/service"
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
	config  *config.Config
	handler handlers.Handler
	router  *http.ServeMux
	server  *http.Server
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
	s.InitConfig()
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

func (s *APIServer) InitConfig() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error initializing config db: %s", err)
	}

	dbConfig := repository.PostgresConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := repository.NewPostgresDB(dbConfig)
	if err != nil {
		log.Fatalf("Failed to initialize db: %s", err)
	}

	repo, err := repository.NewRepository(db)
	if err != nil {
		log.Fatalf("Failed to initialize repos: %s", err)
	}

	serv := service.NewService(repo)
	s.handler = *handlers.NewHandler(serv)
}
