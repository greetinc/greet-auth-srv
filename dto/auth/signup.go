package dto

// Register
type AuthRegisterRequest struct {
	ID string `json:"id"`
	// ProfileID string `json:"profile_id" form:"profile_id" validate:"required,profile_id"`
	// FullName  string `json:"full_name" form:"full_name" validate:"required,full_name"`
	Otp      string `json:"otp" form:"otp" validate:"required,otp"`
	Whatsapp string `json:"whatsapp" form:"whatsapp" validate:"required,whatsapp"`
	// Age       int    `json:"age" form:"age" validate:"required,age"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,password"`
	Token    string `json:"token"`
}

type AuthRegisterResponse struct {
	ID string `json:"id"`
	// FullName  string `json:"full_name" form:"full_name" validate:"required,full_name"`
	// ProfileID string `json:"profile_id" form:"profile_id" validate:"required,profile_id"`
	Whatsapp string `json:"whatsapp" form:"whatsapp" validate:"required,whatsapp"`
	// Age       int    `json:"age" form:"age" validate:"required,age"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"-" form:"password" validate:"required,password"`
	Token    string `json:"token"`
}
