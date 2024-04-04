package repositories

import (
	dto "greet-auth-srv/dto/auth"
	"greet-auth-srv/entity"
	"sync"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Signin(req dto.SigninRequest) (*entity.User, error)
	SigninByPhoneNumber(req dto.SigninRequest) (*entity.User, error)
	UpdateTokenVerified(userID string, otp string, token string) (dto.LoginResponse, error)
	UpdateUser(user *entity.User) error
}

type authRepository struct {
	DB    *gorm.DB
	mu    sync.Mutex
	users map[string]*entity.User
}

func NewAuthRepository(DB *gorm.DB) DomainRepository {
	return &authRepository{
		DB: DB,
	}
}
