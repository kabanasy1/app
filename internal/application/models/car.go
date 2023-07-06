package models

type Car struct {
	Id            int64  `json:"id"`
	Brand         string `json:"brand"`
	Model         string `json:"model"`
	Colour        string `json:"colour"`
	LisenceNumber string `json:"lisence_number"`
	Owner         User   `json:"owner"`
}
