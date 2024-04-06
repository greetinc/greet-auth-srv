package repositories

import (
	dto "github.com/greetinc/greet-auth-srv/dto/auth"
	"github.com/greetinc/greet-auth-srv/entity"
	util "github.com/greetinc/greet-util/s"
)

func (u *authRepository) SignupDetail(req dto.RegisterDetailRequest) (dto.RegisterDetailResponse, error) {

	user := entity.UserDetail{
		ID:        req.ID,
		FullName:  req.FullName,
		UserID:    req.UserID,
		ProfileID: req.ProfileID,
		Age:       req.Age,
		Gender:    req.Gender,
	}

	if err := u.DB.Save(&user).First(&user).Error; err != nil {
		return dto.RegisterDetailResponse{}, err
	}

	age := entity.InterestAge{
		ID:     util.GenerateRandomString(),
		UserID: user.ID,
		MinAge: 16,
		MaxAge: 55,
	}

	if err := u.DB.Save(&age).First(&age).Error; err != nil {
		return dto.RegisterDetailResponse{}, err
	}

	distance := entity.Distance{
		ID:       util.GenerateRandomString(),
		UserID:   user.ID,
		Distance: 80,
	}

	if err := u.DB.Save(&distance).First(&distance).Error; err != nil {
		return dto.RegisterDetailResponse{}, err
	}

	response := dto.RegisterDetailResponse{
		ID:        user.ID,
		FullName:  user.FullName,
		ProfileID: user.ProfileID,
		UserID:    user.UserID,
		Age:       req.Age,
		Gender:    req.Gender,
	}

	return response, nil
}
