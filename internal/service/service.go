package service

import (
	"crypto/sha256"
	"fmt"
	"math/rand"

	"gitlab.com/zapirus/shortener/internal/models"
)

func GenerateShortUrl(beforeURL string) models.GetShortURLResponse {
	hash := sha256.Sum256([]byte(beforeURL))
	randValue := rand.Intn(6) + 4
	result := fmt.Sprintf("%x", hash[:randValue])
	return models.GetShortURLResponse{AfterURL: "http://localhost:8080/" + result}
}
