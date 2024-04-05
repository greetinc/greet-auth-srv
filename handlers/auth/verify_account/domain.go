package handlers

import (
	s "github.com/greetinc/greet-auth-srv/services/auth/verify_account"

	"github.com/labstack/echo/v4"
)

type VerifyHandler interface {
	HandleVerification(c echo.Context) error //verifikasi email
	ResendVerification(c echo.Context) error //verifikasi resend code
}

type verifyHandler struct {
	serviceVerify s.VerifyService
}

func NewVerifyHandler(service s.VerifyService) VerifyHandler {
	return &verifyHandler{
		serviceVerify: service,
	}
}
