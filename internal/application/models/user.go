package models

type User struct {
	Id           int64  `json:"id"`
	FirstName    string `json:"name"`
	LastName     string `json:"surname"`
	OfficeNumber int64  `json:"office"`
	PhoneNumber  string `json:"phone_number"`
	Position     string `json:"position"`
}
