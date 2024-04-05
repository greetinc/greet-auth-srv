package auth

import (
	dto "greet-auth-srv/dto/auth"
	m "greet-auth-srv/middlewares"
	r "greet-auth-srv/repositories/auth"
)

type AuthService interface {
	Signin(req dto.SigninRequest) (*dto.LoginResponse, error)
	SigninByPhoneNumber(req dto.SigninRequest) (*dto.LoginResponse, error)
	Signup(req dto.AuthRegisterRequest) (dto.AuthRegisterResponse, error)
	SignupDetail(req dto.RegisterDetailRequest) (dto.RegisterDetailResponse, error)
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
