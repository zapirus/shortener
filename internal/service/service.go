package service

import "gitlab.com/zapirus/shortener/internal/models"

func Hello() models.GetHello {
	return models.GetHello{Message: "Hello worlds"}
}

func ResponseURL(urlRequest models.ResponseURL) models.ResponseURL {
	return models.ResponseURL{BeforeURL: urlRequest.BeforeURL, AfterURL: "https://ya.ru"}
}
