package app

import (
	"gitlab.com/zapirus/shortener/internal/handlers"
)

func (s *APIServer) confRouter() {
	s.router.HandleFunc("/short-url/create", handlers.ShortenURLHandler)
	s.router.HandleFunc("/", handlers.RedirectHandler)
}
