package entity

type User struct {
	ID    int `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Idade int `json:"idade"`
}