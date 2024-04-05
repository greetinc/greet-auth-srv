package repositories

import (
	"github.com/greetinc/greet-auth-srv/entity"
)

func (u *verifyResetRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := u.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
