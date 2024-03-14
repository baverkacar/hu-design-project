package model

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type UserCreateModel struct {
	FullName string `json:"fullName" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (u *UserCreateModel) ValidateUser() error {
	if validate == nil {
		validate = validator.New()
	}
	return validate.Struct(u)
}
