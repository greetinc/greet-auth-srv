package repositories

import (
	"sync"

	dto "github.com/greetinc/greet-auth-srv/dto/auth"
	"github.com/greetinc/greet-auth-srv/entity"

	"gorm.io/gorm"
)

type DomainRepository interface {
	Signup(req dto.AuthRegisterRequest) (dto.AuthRegisterResponse, error)
	SignupDetail(req dto.RegisterDetailRequest) (dto.RegisterDetailResponse, error)
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
