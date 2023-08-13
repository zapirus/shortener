package service

import "gitlab.com/zapirus/shortener/internal/models"

func Hello() models.GetHello {
	return models.GetHello{Message: "Hello worlds"}
}

func GetShortUrlHandler(urlRequest models.GetShortURLRequest) models.GetShortURLResponse {
	return models.GetShortURLResponse{AfterURL: "https://ya.ru"}
}
