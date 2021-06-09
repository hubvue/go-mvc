package service

import (
	"server-test/src/model"
	"server-test/src/repository"
)

type UserService interface {
	GetUser(id int) *model.User
}

type userService struct{
	userRepo repository.UserRepo
}

func NewUserService() *userService {
	return &userService{
		userRepo: repository.NewUserRepo(),
	}
}

func (service *userService) GetUser(id int) *model.User {
	userInfo := service.userRepo.GetUserById(id)
	return userInfo
}
