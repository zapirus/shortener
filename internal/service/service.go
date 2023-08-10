package service

import "gitlab.com/zapirus/shortener/internal/models"

func Hello() models.GetHello {
	hello := models.GetHello{Message: "Hello worlds"}
	return hello
}
