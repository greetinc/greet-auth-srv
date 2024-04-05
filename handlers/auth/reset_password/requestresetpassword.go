package handlers

import (
	dto "greet-auth-srv/dto/auth"
	res "greet-auth-srv/util/response"

	"github.com/labstack/echo/v4"
)

func (u *resetHandler) RequestResetPassword(c echo.Context) error {
	// Bind request body to the ResetPasswordRequest struct
	var req dto.ResetPasswordRequest
	if err := c.Bind(&req); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	// Panggil fungsi service dengan parameter dari DTO
	data, err := u.serviceReset.RequestResetPassword(req)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)

}
