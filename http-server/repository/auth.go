package repository

import (
	"context"
	"http-server/database"
	"http-server/models"

	"github.com/rs/zerolog/log"
)

func GetUserById(userId string) *models.User {
	var user models.User

	sql := "SELECT user_id, username, email, firstname, lastname, created_at, updated_at FROM users WHERE user_id = $1"
	err := database.Get().QueryRow(context.Background(), sql, userId).Scan(
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

func CreateNewUser(username string, hashedPassword string, email string, firstname string, lastname string) *models.User {
	user := models.User{
		Username:  username,
		Email:     email,
		FirstName: firstname,
		LastName:  lastname,
	}

	sql := "INSERT INTO users (username, password_hash, email, firstname, lastname) VALUES ($1, $2, $3, $4, $5) RETURNING user_id, created_at, updated_at"
	err := database.Get().QueryRow(context.Background(), sql,
		username,
		hashedPassword,
		email,
		firstname,
		lastname).Scan(&user.ID, &user.CreatedAt, &user.LastUpdatedAt)

	if err != nil {
		log.Error().Msgf("%s", err.Error())
		return nil
	}

	return &user
}

func GetUserByEmailAndPassword(email string, password string) *models.User {
	var user models.User

	sql := "SELECT user_id, username, email, firstname, lastname, created_at, updated_at FROM users WHERE email = $1 AND password_hash = $2"
	err := database.Get().QueryRow(context.Background(), sql, email, password).Scan(
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
