package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"gitlab.com/zapirus/shortener/internal/models"
	"gitlab.com/zapirus/shortener/internal/service"
)

var urlMap = make(map[string]string)

type Handler struct {
	service service.Service
}

func NewHandler(service service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	var req models.GetShortURLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shortURL, err := h.service.GenerateShortURL(req.BeforeURL)
	if err != nil {
		log.Fatalf("Failed to generate URL: %s", err)
	}

	urlMap[shortURL] = req.BeforeURL

	resp := models.GetShortURLResponse{
		AfterURL: shortURL,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:]
	fullURL, ok := urlMap[shortURL]
	if !ok {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, fullURL, http.StatusSeeOther)
}
