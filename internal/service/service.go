package service

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
)

func GenerateShortUrl(beforeURL string) string {
	hash := sha256.Sum256([]byte(beforeURL))
	randValue := rand.Intn(6) + 4
	shortURL := fmt.Sprintf("%x", hash[:randValue])
	return shortURL
}
