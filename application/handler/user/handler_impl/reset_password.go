package handler_impl

import (
	"context"
	"github.com/labstack/gommon/log"
	"hu-design-project/application/repository"
	"hu-design-project/infrastructure/util"
)

type ResetPasswordHandler struct {
	repository repository.UserRepository
}

func NewResetPasswordHandler(
	repository repository.UserRepository,
) *ResetPasswordHandler {
	return &ResetPasswordHandler{
		repository: repository,
	}
}
func (handler *ResetPasswordHandler) Handle(ctx context.Context, email string) error {
	// Rastgele kod oluşturma
	code, err := util.CreateVerificationCode()
	if err != nil {
		return err
	}

	// Şifreyi veritabanında güncelleme
	err = handler.repository.UpdatePasswordByEmail(ctx, email, code)
	if err != nil {
		return err
	}

	// E-postayı gönderme
	_, err = util.SendMail(email, code)
	if err != nil {
		return err
	}

	log.Printf("Password reset code sent to %s: %s", email, code)
	return nil
}
