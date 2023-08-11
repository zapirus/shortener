package service

import "gitlab.com/zapirus/shortener/internal/models"

func Hello() models.GetHello {
	return models.GetHello{Message: "Hello worlds"}
}
