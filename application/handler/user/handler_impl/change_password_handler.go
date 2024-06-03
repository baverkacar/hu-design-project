package handler_impl

import (
	"context"
	"errors"
	"hu-design-project/application/model"
	"hu-design-project/application/repository"
)

type ChangePasswordHandler struct {
	repository repository.UserRepository
}

func NewChangePasswordHandler(
	repository repository.UserRepository) *ChangePasswordHandler {
	return &ChangePasswordHandler{
		repository: repository,
	}
}

func (handler *ChangePasswordHandler) Handle(ctx context.Context, model model.ChangePasswordModel) error {
	// Retrieve the user by ID to check the old password.
	user, err := handler.repository.GetUserByEmail(ctx, model.Email)
	if err != nil {
		return err // Error retrieving user
	}

	// Check if the old password matches
	if user.Password != model.OldPassword {
		return errors.New("old password does not match")
	}

	// Change the password to the new password
	err = handler.repository.ChangePassword(ctx, user.UserID.Hex(), model.OldPassword, model.NewPassword)
	if err != nil {
		return err // Error changing password
	}

	return nil
}
