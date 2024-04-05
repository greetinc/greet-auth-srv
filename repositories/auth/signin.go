package repositories

import (
	"time"

	dto "github.com/greetinc/greet-auth-srv/dto/auth"
	"github.com/greetinc/greet-auth-srv/entity"
	"github.com/greetinc/greet-auth-srv/util"
)

func (u *authRepository) SigninByPhoneNumber(req dto.SigninRequest) (*entity.User, error) {
	var existingUser entity.User
	err := u.DB.Preload("Verified").Preload("UserDetail").Where("whatsapp = ?", req.Whatsapp).First(&existingUser).Error
	if err != nil {
		return nil, err
	}

	return &existingUser, nil
}

func (u *authRepository) Signin(req dto.SigninRequest) (*entity.User, error) {
	var existingUser entity.User
	err := u.DB.Preload("Verified").Preload("UserDetail").Where("email = ?", req.Email).First(&existingUser).Error
	if err != nil {
		return nil, err
	}

	return &existingUser, nil
}

func (u *authRepository) UpdateUser(user *entity.User) error {
	return u.DB.Model(user).Updates(user).Error
}

func (u *authRepository) UpdateTokenVerified(userID string, otp string, token string) (dto.LoginResponse, error) {
	verified := entity.UserVerified{
		ID:        util.GenerateRandomString(),
		UserID:    userID,
		Otp:       otp,
		Token:     token,
		ExpiredAt: time.Now().Add(4 * time.Minute),
	}

	if err := u.DB.Save(&verified).Error; err != nil {
		return dto.LoginResponse{}, err
	}

	response := dto.LoginResponse{
		VerifiedResp: &dto.AuthUnverifiedResponse{
			TokenVerified: verified.Token,
			Otp:           otp,
		},
	}

	return response, nil
}
