package handlers

import (
	"net/http"

	dto "github.com/greetinc/greet-auth-srv/dto/auth"
	res "github.com/greetinc/greet-util/s/response"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func (u *resetHandler) ResetPassword(c echo.Context) error {

	var req dto.Reset

	if err := c.Bind(&req); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	token := c.QueryParam("token")
	req.Token = token

	err := u.serviceReset.ResetPassword(req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Password reset request successfully processed",
	})
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func validateStruct(s interface{}) error {
	return validate.Struct(s)
}
