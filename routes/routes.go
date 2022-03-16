package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/magicwarms/startup-backend/config"
)

// AppRoutes is all routes in app
func AppRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")
	v1.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON(config.AppResponse{
			Code:    http.StatusOK,
			Message: "SUCCESS",
			Data:    nil,
		})
	})
	v1.Get("/stack", func(c *fiber.Ctx) error {
		return c.JSON(c.App().Stack())
	})

}
