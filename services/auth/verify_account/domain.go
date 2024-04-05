package auth

import (
	dto "github.com/greetinc/greet-auth-srv/dto/auth"
	"github.com/greetinc/greet-auth-srv/entity"
	m "github.com/greetinc/greet-auth-srv/middlewares"
	r "github.com/greetinc/greet-auth-srv/repositories/auth/verify_account"
)

type VerifyService interface {
	VerifyUserByToken(req dto.VerificationRequest) (*entity.UserVerified, error)
	ResendVerifyUserByToken(req dto.ResendVerificationRequest) (*entity.UserVerified, error)
}

type verifyService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewVerifyService(Repo r.DomainRepository, jwtS m.JWTService) VerifyService {
	return &verifyService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
