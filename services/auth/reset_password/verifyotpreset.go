package auth

import (
	dto "greet-auth-srv/dto/auth"
	"greet-auth-srv/entity"
	res "greet-auth-srv/util/response"
	"time"
)

func (u *verifyService) VerifyOtpReset(req dto.VerifyResetRequest) (*entity.PasswordResetToken, error) {
	// Use your repository or service to fetch the user by token from the database
	user, err := u.Repo.VerifyOtpReset(req)
	if err != nil {
		// Handle the error (e.g., database query error)
		return nil, err
	}

	// Pemeriksaan waktu kadaluwarsa OTP
	if time.Now().After(user.ExpiredAt) {
		return nil, res.ErrorBuilder(&res.ErrorConstant.ExpiredToken, err)
	}

	return user, nil
}
