package handlers

import (
	"strings"

	dto "github.com/greetinc/greet-auth-srv/dto/auth"
	res "github.com/greetinc/greet-auth-srv/util/response"

	"github.com/labstack/echo/v4"
)

// Login
// @Security BearerAuth
// @Summary Login user
// @Description Login user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body  dto.SigninRequest true "request body"
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /api/auth/login [post]
func (u *domainHandler) Signin(c echo.Context) error {
	var req dto.SigninRequest
	var resp *dto.LoginResponse
	var errResponse error

	err := c.Bind(&req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	switch {
	case req.Email != "":
		req.Email = strings.ToLower(req.Email)
		resp, errResponse = u.serviceAuth.Signin(req)
	case req.Whatsapp != "":
		resp, errResponse = u.serviceAuth.SigninByPhoneNumber(req)
	default:
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	if errResponse != nil {
		return res.ErrorResponse(errResponse).Send(c)
	}

	return res.SuccessResponse(resp).Send(c)
}
