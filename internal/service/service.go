package service

import (
	"crypto/sha256"
	"fmt"
	"math/rand"

	"gitlab.com/zapirus/shortener/internal/models"
)

func GetShortUrl(beforeURL string) models.GetShortURLResponse {
	hash := sha256.Sum256([]byte(beforeURL))
	randValue := rand.Intn(6) + 4
	result := fmt.Sprintf("%x", hash[:randValue])
	return models.GetShortURLResponse{AfterURL: result}
}
