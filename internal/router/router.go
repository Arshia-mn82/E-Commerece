package router

import (
	"E-Commerce/internal/handler"
	"E-Commerce/internal/middleware"

	"github.com/gofiber/fiber/v3"
)

type Deps struct {
	AuthHandler *handler.AuthHandler
	JWTSecret   string
}

func Setup(app *fiber.App, deps Deps) {
	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	api := app.Group("/api/v1")

	// Auth Public
	auth := api.Group("/auth")
	auth.Post("/signup", deps.AuthHandler.SignUp)
	auth.Post("/login", deps.AuthHandler.Login)

	// Auth Portected
	protected := auth.Group("", middleware.AuthRequired(deps.JWTSecret))
	protected.Get("/me", deps.AuthHandler.Me)
}
