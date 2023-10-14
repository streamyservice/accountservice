package config

import (
	"accountservice/app/controller"
	"accountservice/app/repository"
	"accountservice/app/service"
)

type Initialization struct {
	userRepo      repository.UserRepository
	userTokenRepo repository.UserTokenRepository
	userSvc       service.UserService
	UserCtrl      controller.UserController
}

func NewInitialization(
	userTokenRepo repository.UserTokenRepository,
	userRepo repository.UserRepository,
	userService service.UserService,
	userCtrl controller.UserController,
) *Initialization {
	return &Initialization{
		userTokenRepo: userTokenRepo,
		userRepo:      userRepo,
		userSvc:       userService,
		UserCtrl:      userCtrl,
	}
}
