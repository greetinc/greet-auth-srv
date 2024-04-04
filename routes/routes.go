package routes

import (
	"greet-auth-srv/configs"
	"greet-auth-srv/docs"
	h_auth "greet-auth-srv/handlers/auth"
	"net/http"

	"greet-auth-srv/middlewares"
	r_auth "greet-auth-srv/repositories/auth"

	s_auth "greet-auth-srv/services/auth"

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
)

func New() *echo.Echo {

	e := echo.New()

	middlewares.LoggerMiddleware(e)

	docs.SwaggerInfo.Title = os.Getenv("APP")
	docs.SwaggerInfo.Version = os.Getenv("VERSION")
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	docs.SwaggerInfo.Schemes = []string{os.Getenv("SCHEME")}

	e.GET("/swagger/*any", echoSwagger.WrapHandler)

	auth := e.Group("api/auth", middlewares.ApiKeyMiddleware)
	{
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
