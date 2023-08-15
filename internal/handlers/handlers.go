package handlers

import (
	"encoding/json"
	"net/http"

	"gitlab.com/zapirus/shortener/internal/models"
	"gitlab.com/zapirus/shortener/internal/service"
)

var urlMap = make(map[string]string)

func ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	var req models.GetShortURLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shortURL := service.GenerateShortUrl(req.BeforeURL)
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

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:]
	fullURL, ok := urlMap[shortURL]
	if !ok {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, fullURL, http.StatusSeeOther)
}
