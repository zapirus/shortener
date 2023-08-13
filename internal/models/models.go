package models

type GetHello struct {
	Message string `json:"message"`
}

type GetShortURLRequest struct {
	BeforeURL string `json:"beforeURL"`
}

type GetShortURLResponse struct {
	AfterURL string `json:"afterURL"`
}
