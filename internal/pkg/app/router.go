package app

import (
	"gitlab.com/zapirus/shortener/internal/handlers"
)

func (s *APIServer) confRouter() {
	s.router.HandleFunc("/url", handlers.GetShortUrlHandler()).Methods("POST")
}
