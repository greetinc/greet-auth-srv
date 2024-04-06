package handlers

import (
	dto "github.com/greetinc/greet-auth-srv/dto/auth"
	res "github.com/greetinc/greet-util/s/response"

	"github.com/labstack/echo/v4"
)

func (u *resetHandler) VerifyResetPassword(c echo.Context) error {
	var req dto.VerifyResetRequest

	err := c.Bind(&req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	token := c.QueryParam("token")
	req.Token = token

	data, err := u.serviceReset.VerifyOtpReset(req)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)
}
