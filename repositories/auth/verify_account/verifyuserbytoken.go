package repositories

import (
	"fmt"

	dto "github.com/greetinc/greet-auth-srv/dto/auth"
	"github.com/greetinc/greet-auth-srv/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (u *verifyRepository) VerifyUserByToken(req dto.VerificationRequest) (*entity.UserVerified, error) {
	var user entity.UserVerified
	if err := u.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "token"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"verified": true}),
	}).Where("token = ?", req.Token).Where("otp = ?", req.Otp).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("User not found with the given verification token")
		}
		return nil, err
	}

	return &user, nil
}
