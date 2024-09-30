package service

import (
	"fmt"
	"http-server/models"
	"http-server/repository"
)

func ValidateSessionRequest(r *models.SessionRequest) (*models.User, bool) {
	foundUser := repository.GetUserByEmailAndPassword(r.Email, r.Password)

	if foundUser != nil {
		return foundUser, true
	} else {
		return nil, false
	}
}

func CreateNewAccount(r *models.SignupRequest) (*models.User, error) {
	newUser := repository.CreateNewUser(r.Username, r.Password, r.Email, r.FirstName, r.LastName)

	if newUser != nil {
		return newUser, nil
	} else {
		return nil, fmt.Errorf("failed to create new account") // should improve this eventually
	}
}
