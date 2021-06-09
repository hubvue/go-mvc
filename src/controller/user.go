package controller

import (
	"github.com/gofiber/fiber/v2"
	"server-test/src/service"
	"strconv"
)

type UserController interface {
	GetUserInfo(ctx *fiber.Ctx) error
}

type userController struct {
	userService service.UserService
}

func NewUserController() *userController {
	return &userController{
		userService: service.NewUserService(),
	}
}

func (controller *userController) GetUserInfo(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		panic("id is not int type")
	}
	userInfo := controller.userService.GetUser(id)
	return ctx.JSON(userInfo)
}
