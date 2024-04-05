package repositories

import (
	dto "greet-auth-srv/dto/auth"
	"greet-auth-srv/entity"
	"sync"
	"time"

	"gorm.io/gorm"
)

type ResetRepository interface {
	VerifyOtpReset(req dto.VerifyResetRequest) (*entity.PasswordResetToken, error)
	SavePasswordResetToken(userID string, token string, otp string, expiryDuration time.Duration) (*entity.PasswordResetToken, error)
	GetUserByEmail(email string) (*entity.User, error)
	ValidatePasswordResetToken(token string) (string, error)
	UpdateUserPassword(userID string, newPassword string) error
	ResendReset(req dto.ResendResetRequest) (*entity.PasswordResetToken, error)
}

type verifyResetRepository struct {
	DB    *gorm.DB
	mu    sync.Mutex
	users map[string]*entity.User
}

func NewResetRepository(DB *gorm.DB) ResetRepository {
	return &verifyResetRepository{
		DB: DB,
	}
}
