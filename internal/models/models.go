package models

type GetShortURLRequest struct {
	BeforeURL string `json:"beforeURL"`
}

type GetShortURLResponse struct {
	AfterURL string `json:"afterURL"`
}
