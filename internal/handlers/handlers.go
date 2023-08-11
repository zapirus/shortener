package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"gitlab.com/zapirus/shortener/internal/service"
)

func GetHello() http.HandlerFunc {
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
