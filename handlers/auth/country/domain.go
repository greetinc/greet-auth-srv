package handlers

import (
	s "github.com/greetinc/greet-auth-srv/services/auth/country"

	"github.com/labstack/echo/v4"
)

type CountryHandler interface {
	GetAll(c echo.Context) error
	GetById(c echo.Context) error
	// Update(c echo.Context) error
	// Delete(c echo.Context) error
}

type domainHandler struct {
	domainService s.CountryService
}

func NewCountryHandler(CountryS s.CountryService) CountryHandler {
	return &domainHandler{
		domainService: CountryS,
	}
}
