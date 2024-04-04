package main

import (
	"greet-auth-srv/configs/seeder"
	"greet-auth-srv/routes"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {

	dbEvent := os.Getenv("DBEVENT")
	if dbEvent == "seeder" {
		seeder.RunSeeder()
	}

	e := routes.New()
	e.Use(CORSMiddleware())

	e.Logger.Fatal(e.Start(":8080"))
}

// CORSMiddleware ..
func CORSMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
			c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
			c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			c.Response().Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, DELETE, OPTIONS, PATCH")

			if c.Request().Method == "OPTIONS" {
				return c.NoContent(204)
			}

			return next(c)
		}
	}
}