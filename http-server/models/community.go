package models

import (
	"context"
	"http-server/database"
	"time"

	"github.com/rs/zerolog/log"
)

type Community struct {
	ID            string
	Name          string
	Description   string
	ownerId       string
	CreatedAt     time.Time
	LastUpdatedAt time.Time
	Region        string
	IconUrl       string
}

func (c *Community) GetOwner() *User {
	var user User

	sql := "SELECT user_id, username, email, firstname, lastname, created_at, updated_at FROM users WHERE user_id = $1"
	err := database.Get().QueryRow(context.Background(), sql, c.ownerId).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
		&user.LastUpdatedAt,
	)

	if err != nil {
		log.Error().Msgf("%s", err.Error())
		return nil
	}

	return &user
}
