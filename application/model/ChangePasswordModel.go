package model

type ChangePasswordModel struct {
	Email       string `json:"email"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
