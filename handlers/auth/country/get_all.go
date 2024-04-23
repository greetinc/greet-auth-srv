package handlers

import (
	"github.com/labstack/echo/v4"
)

func (b *domainHandler) GetAll(c echo.Context) error {
	get, err := b.domainService.GetAll()
	if err != nil {
		return c.JSON(500, "Error fetching get")
	}

	return c.JSON(200, get)
}
