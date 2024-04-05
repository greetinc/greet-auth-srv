package routes

import (
	"greet-auth-srv/configs"
	"greet-auth-srv/docs"
	h_auth "greet-auth-srv/handlers/auth"
	h_verifyReset "greet-auth-srv/handlers/auth/reset_password"
	h_verify "greet-auth-srv/handlers/auth/verify_account"
	"net/http"

	"greet-auth-srv/middlewares"
	r_auth "greet-auth-srv/repositories/auth"
	r_verifyReset "greet-auth-srv/repositories/auth/reset_password"
	r_verify "greet-auth-srv/repositories/auth/verify_account"

	s_auth "greet-auth-srv/services/auth"
	s_verifyReset "greet-auth-srv/services/auth/reset_password"
	s_verify "greet-auth-srv/services/auth/verify_account"

	"os"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var (
	DB = configs.InitDB()

	JWT = middlewares.NewJWTService()

	authR = r_auth.NewAuthRepository(DB)
	authS = s_auth.NewAuthService(authR, JWT)
	authH = h_auth.NewAuthHandler(authS)

	verifyR = r_verify.NewVerifyRepository(DB)
	verifyS = s_verify.NewVerifyService(verifyR, JWT)
	verifyH = h_verify.NewVerifyHandler(verifyS)

	resetR = r_verifyReset.NewResetRepository(DB)
	resetS = s_verifyReset.NewResetService(resetR, JWT)
	resetH = h_verifyReset.NewResetHandler(resetS)
)

func New() *echo.Echo {

	e := echo.New()

	e.POST("/verify", verifyH.HandleVerification)
	e.PUT("/resend-otp", verifyH.ResendVerification)

	//reset password
	e.POST("/resetpassword", resetH.ResetPassword)
	e.POST("/verify-reset", resetH.VerifyResetPassword)
	e.POST("/request-reset-password", resetH.RequestResetPassword)
	e.PUT("/resend-reset", resetH.ResendVerification)

	middlewares.LoggerMiddleware(e)

	docs.SwaggerInfo.Title = os.Getenv("APP")
	docs.SwaggerInfo.Version = os.Getenv("VERSION")
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	docs.SwaggerInfo.Schemes = []string{os.Getenv("SCHEME")}

	e.GET("/swagger/*any", echoSwagger.WrapHandler)

	auth := e.Group("api/auth", middlewares.ApiKeyMiddleware)
	{
		auth.POST("/register", authH.Signup)
		auth.POST("/register-detail", authH.SignupDetail)
		auth.POST("/login", authH.Signin)
	}

	return e
}

func echoHandlerWrapper(h http.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		h(c.Response(), c.Request())
		return nil
	}
}
