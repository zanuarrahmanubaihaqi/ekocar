package http

import (
	"eko-car/delivery/container"
	"eko-car/domain/car"
	"eko-car/domain/user"
)

type handler struct {
	userHandler user.UserHandler
	carHandler  car.CarHandler
}

func SetupHandler(container container.Container) handler {
	return handler{
		userHandler: user.NewUserHandler(container.UserFeature),
		carHandler:  car.NewCarHandler(container.CarFeature),
	}
}
