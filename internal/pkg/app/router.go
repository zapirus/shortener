package app

func (s *APIServer) confRouter() {
	s.router.HandleFunc("/short-url/create", s.handler.ShortenURLHandler)
	s.router.HandleFunc("/", s.handler.RedirectHandler)
}
