package services

import (
	"context"
	"errors"
	"trocup-user/models"
	"trocup-user/repository"
)

// CheckIfUserExists performs checks for both email and pseudo before creating a user
func CheckIfUserExists(ctx context.Context, email, pseudo string) error {
	// Check if a user with the given email already exists
	existingUserByEmail, err := repository.FindUserByEmail(ctx, email)
	if err != nil {
		return err
	}
	if existingUserByEmail != nil {
		return errors.New("email already in use")
	}

	// Check if a user with the given pseudo already exists
	existingUserByPseudo, err := repository.FindUserByPseudo(ctx, pseudo)
	if err != nil {
		return err
	}
	if existingUserByPseudo != nil {
		return errors.New("pseudo already in use")
	}

	// If both checks pass, return nil (no errors)
	return nil
}


// CreateUser creates a new user in the database without password handling
func CreateUser(ctx context.Context, user *models.User) error {
	// Directly create the user in the database without password handling
	return repository.CreateUser(ctx, user)
}