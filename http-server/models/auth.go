package models

type SessionRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	// TODO
}

type User struct {
	Username string
	Email    string
	// TODO
}
