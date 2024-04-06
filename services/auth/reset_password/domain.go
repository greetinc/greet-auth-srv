package auth

import (
	dto "github.com/greetinc/greet-auth-srv/dto/auth"
	"github.com/greetinc/greet-auth-srv/entity"
	r "github.com/greetinc/greet-auth-srv/repositories/auth/reset_password"
	m "github.com/greetinc/greet-middlewares/middlewares"
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
