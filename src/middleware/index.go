package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func Init(app *fiber.App) {
	app.Use(testMiddleware)
}
