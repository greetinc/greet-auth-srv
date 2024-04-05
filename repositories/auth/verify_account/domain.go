package repositories

import (
	dto "greet-auth-srv/dto/auth"
	"greet-auth-srv/entity"
	"sync"

	"gorm.io/gorm"
)

type DomainRepository interface {
	UpdateUserVerificationStatus(user *entity.UserVerified) error
	VerifyUserByToken(req dto.VerificationRequest) (*entity.UserVerified, error)
	ResendVerifyUserByToken(req dto.ResendVerificationRequest) (*entity.UserVerified, error)
}

type verifyRepository struct {
	DB    *gorm.DB
	mu    sync.Mutex
	users map[string]*entity.User
}

func NewVerifyRepository(DB *gorm.DB) DomainRepository {
	return &verifyRepository{
		DB: DB,
	}
}
