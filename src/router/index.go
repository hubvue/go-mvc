package router

import (
	"github.com/gofiber/fiber/v2"
	"server-test/src/controller"
)

func Init(app *fiber.App) {
	userController := controller.NewUserController()
	userGroup := app.Group("/user")
	userGroup.Get("/getUser", userController.GetUserInfo)
}
