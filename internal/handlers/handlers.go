package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sirupsen/logrus"

	"gitlab.com/zapirus/shortener/internal/models"
	"gitlab.com/zapirus/shortener/internal/service"
)

func GetHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		hello := service.Hello()
		if err := json.NewEncoder(w).Encode(hello); err != nil {
			log.Printf("Не удалось преобразовать: %s", err)
			w.WriteHeader(http.StatusConflict)
			return
		}
	}
}

func ResponseURL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var urlRequest models.ResponseURL
		if err := json.NewDecoder(r.Body).Decode(&urlRequest); err != nil {
			logrus.Printf("Не удалось преобразовать: %s", err)
			w.WriteHeader(http.StatusConflict)
			return
		}
		urlResponse := service.ResponseURL(urlRequest)
		if err := json.NewEncoder(w).Encode(urlResponse); err != nil {
			log.Printf("Не удалось преобразовать: %s", err)
			w.WriteHeader(http.StatusConflict)
			return
		}
	}
}
