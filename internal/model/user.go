package model

type User struct {
	ID       string `json:"id" ginboot:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
