package repositories

import (
	dto "greet-auth-srv/dto/auth"
	"greet-auth-srv/entity"
	"greet-auth-srv/util"
	"time"
)

func (u *authRepository) Signup(req dto.AuthRegisterRequest) (dto.AuthRegisterResponse, error) {

	user := entity.User{
		ID:       req.ID,
		Whatsapp: req.Whatsapp,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := u.DB.Save(&user).First(&user).Error; err != nil {
		return dto.AuthRegisterResponse{}, err
	}

	verified := entity.UserVerified{
		ID:        util.GenerateRandomString(),
		UserID:    user.ID,
		Token:     req.Token,
		Otp:       req.Otp,
		ExpiredAt: time.Now().Add(4 * time.Minute),
	}

	if err := u.DB.Save(&verified).First(&verified).Error; err != nil {
		return dto.AuthRegisterResponse{}, err
	}

	age := entity.InterestAge{
		ID:     util.GenerateRandomString(),
		UserID: user.ID,
		MinAge: 16,
		MaxAge: 55,
	}

	if err := u.DB.Save(&age).First(&age).Error; err != nil {
		return dto.AuthRegisterResponse{}, err
	}

	response := dto.AuthRegisterResponse{
		ID:       user.ID,
		Whatsapp: user.Whatsapp,
		Email:    user.Email,
		Password: user.Password,
		Token:    verified.Token,
	}

	return response, nil
}
