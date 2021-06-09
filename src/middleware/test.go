package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func testMiddleware(ctx *fiber.Ctx) error {
	fmt.Println("test Middleware")
	return ctx.Next()
}
