package handlers

import (
	dto "github.com/greetinc/greet-auth-srv/dto/auth"
	"github.com/greetinc/greet-auth-srv/util"
	res "github.com/greetinc/greet-auth-srv/util/response"

	"github.com/labstack/echo/v4"
)

// Register
// @Summary Register user
// @Description Register user
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.AuthRegisterRequest true "request body"
// @Success 200 {object} dto.AuthRegisterResponse
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /api/auth/register [post]
func (u *domainHandler) SignupDetail(c echo.Context) error {

	var req dto.RegisterDetailRequest
	var resp dto.RegisterDetailResponse

	err := c.Bind(&req)
	if err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	userid := c.QueryParam("user_id")
	req.UserID = userid

	resp, err = u.serviceAuth.SignupDetail(req)
	if err != nil {
		if util.IsDuplicateEntryError(err) {
			return res.ErrorResponse(&res.ErrorConstant.Duplicate).Send(c)
		}
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(resp).Send(c)

}
