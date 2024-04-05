package auth

import (
	"time"

	dto "github.com/greetinc/greet-auth-srv/dto/auth"
	"github.com/greetinc/greet-auth-srv/entity"
	res "github.com/greetinc/greet-auth-srv/util/response"
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
