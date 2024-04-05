package dto

import "time"

type SigninRequest struct {
	Whatsapp string `json:"whatsapp" validate:"required,whatsapp"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,password"`
}

type AuthVerified struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Token     string    `json:"token"`
	Verified  bool      `json:"verified"`
	Otp       string    `json:"otp"`
	ExpiredAt time.Time `json:"expired_at"`
}

type LoginResponse struct {
	ID            string `json:"id"`
	FullName      string `json:"full_name"`
	ProfileID     string `json:"profile_id"`
	Email         string `json:"email"`
	Token         string `json:"token"`
	TokenVerified string `json:"token_verified"`
	VerifiedResp  *AuthUnverifiedResponse
}

type AuthUnverifiedResponse struct {
	Email         string `json:"email"`
	Otp           string `json:"otp"`
	TokenVerified string `json:"token_verified"`
}

type UserRequest struct {
	ID             string `param:"id" validate:"required"`
	ProfileID      string `query:"profile_id" validate:"required"`
	SwitcherAccess string `param:"switcher_access" validate:"required"`
}
