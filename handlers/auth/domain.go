package handlers

import (
	s "greet-auth-srv/services/auth"

	"github.com/labstack/echo/v4"
)

type DomainHandler interface {
	Signin(c echo.Context) error //masuk
}

type domainHandler struct {
	serviceAuth s.AuthService
}

func NewAuthHandler(service s.AuthService) DomainHandler {
	return &domainHandler{
		serviceAuth: service,
	}
}