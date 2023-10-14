//go:build wireinject
// +build wireinject

// go:build wireinject
package config

import (
	"accountservice/app/controller"
	"accountservice/app/repository"
	"accountservice/app/service"
	"github.com/google/wire"
)

var db = wire.NewSet(ConnectToDB)

var userServiceSet = wire.NewSet(service.UserServiceInit,
	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
)

var userRepoSet = wire.NewSet(repository.UserRepositoryInit,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
)

var userTokenRepoSet = wire.NewSet(repository.UserTokenRepositoryInit,
	wire.Bind(new(repository.UserTokenRepository), new(*repository.UserTokenRepositoryImpl)),
)

var userCtrlSet = wire.NewSet(controller.UserControllerInit,
	wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)),
)

func Init() *Initialization {
	wire.Build(NewInitialization, db, userCtrlSet, userServiceSet, userRepoSet, userTokenRepoSet)
	return nil
}
