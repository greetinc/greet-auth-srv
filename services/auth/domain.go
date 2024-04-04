package auth

import (
	dto "greet-auth-srv/dto/auth"
	m "greet-auth-srv/middlewares"
	r "greet-auth-srv/repositories/auth"
)

type AuthService interface {
	Signin(req dto.SigninRequest) (*dto.LoginResponse, error)
	SigninByPhoneNumber(req dto.SigninRequest) (*dto.LoginResponse, error)
}

type authService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewAuthService(Repo r.DomainRepository, jwtS m.JWTService) AuthService {
	return &authService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
