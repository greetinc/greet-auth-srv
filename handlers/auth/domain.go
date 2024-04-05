package handlers

import (
	s "greet-auth-srv/services/auth"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Signin(c echo.Context) error //masuk
	Signup(c echo.Context) error //masuk
	SignupDetail(c echo.Context) error
}

type domainHandler struct {
	serviceAuth s.AuthService
}

func NewAuthHandler(service s.AuthService) DomainHandler {
	return &domainHandler{
		serviceAuth: service,
	}
}
