package models

import "time"

type SessionRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	// TODO
}

type SessionResponse struct {
	Token      string `json:"token"`
	WsHostname string `json:"chat_hostname"`
	WsPort     int    `json:"chat_port"`
}

type SignupRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	// TODO
}

type User struct {
	ID            string
	Username      string
	Email         string
	FirstName     string
	LastName      string
	CreatedAt     time.Time
	LastUpdatedAt time.Time
}
