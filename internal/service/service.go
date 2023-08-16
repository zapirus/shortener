package service

import (
	"crypto/sha256"
	"fmt"
	"math/rand"

	"gitlab.com/zapirus/shortener/internal/repository"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s *Service) GenerateShortURL(beforeURL string) (string, error) {
	hash := sha256.Sum256([]byte(beforeURL))
	//rand.Seed(time.Now().UnixNano())
	randValue := rand.Intn(6) + 4
	shortURL := fmt.Sprintf("%x", hash[:randValue])

	if err := s.repo.InsertURL(beforeURL, shortURL); err != nil {
		return "", fmt.Errorf("error inserting URL: %s", err)
	}

	return shortURL, nil
}
