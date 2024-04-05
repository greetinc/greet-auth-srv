package dto

// Register
type RegisterDetailRequest struct {
	ID        string `json:"id"`
	UserID    string `param:"user_id" validate:"required"`
	ProfileID string `json:"profile_id" form:"profile_id" validate:"required,profile_id"`
	FullName  string `json:"full_name" form:"full_name" validate:"required,full_name"`
	Gender    string `json:"gender" form:"gender" validate:"required,gender"`
	Age       int    `json:"age" form:"age" validate:"required,age"`
}

type RegisterDetailResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	FullName  string `json:"full_name" form:"full_name" validate:"required,full_name"`
	ProfileID string `json:"profile_id" form:"profile_id" validate:"required,profile_id"`
	Age       int    `json:"age" form:"age" validate:"required,age"`
	Gender    string `json:"gender" form:"gender" validate:"required,gender"`
	Token     string `json:"token"`
}
