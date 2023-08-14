package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"

	"gitlab.com/zapirus/shortener/internal/models"
	"gitlab.com/zapirus/shortener/internal/service"
)

func GetShortUrlHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var urlRequest models.GetShortURLRequest
		if err := json.NewDecoder(r.Body).Decode(&urlRequest); err != nil {
			logrus.Printf("Не удалось преобразовать: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		urlResponse := service.GetShortUrl(urlRequest.BeforeURL)
		if err := json.NewEncoder(w).Encode(urlResponse); err != nil {
			log.Printf("Не удалось преобразовать: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
