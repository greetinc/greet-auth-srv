package auth

import (
	dto "github.com/greetinc/greet-auth-srv/dto/auth"
	r "github.com/greetinc/greet-auth-srv/repositories/auth"
	m "github.com/greetinc/greet-middlewares/middlewares"
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
