package auth

import (
	dto "greet-auth-srv/dto/auth"
	"greet-auth-srv/entity"
	m "greet-auth-srv/middlewares"
	r "greet-auth-srv/repositories/auth/reset_password"
)

type ResetService interface {
	VerifyOtpReset(req dto.VerifyResetRequest) (*entity.PasswordResetToken, error)
	RequestResetPassword(req dto.ResetPasswordRequest) (dto.ResetPasswordResponse, error)
	ResetPassword(req dto.Reset) error
	ResendReset(req dto.ResendResetRequest) (*entity.PasswordResetToken, error)
}

type verifyService struct {
	Repo r.ResetRepository
	jwt  m.JWTService
}

func NewResetService(Repo r.ResetRepository, jwtS m.JWTService) ResetService {
	return &verifyService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
