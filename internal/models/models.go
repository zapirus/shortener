package models

type GetHello struct {
	Message string `json:"message"`
}

type ResponseURL struct {
	BeforeURL string `json:"beforeURL"`
	AfterURL  string `json:"afterURL"`
}
