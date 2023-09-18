package http

import (
	"eko-car/delivery/container"
	"eko-car/domain/car"
	"eko-car/domain/health"
	"eko-car/domain/logistik"
	"eko-car/domain/sales"
	"eko-car/domain/user"
)

type handler struct {
	healthHandler   health.HealthHandler
	logistikHandler logistik.LogistikHandler
	salesHandler    sales.SalesHandler
	userHandler     user.UserHandler
	carHandler      car.CarHandler
}

func SetupHandler(container container.Container) handler {
	return handler{
		healthHandler:   health.NewHealthHandler(container.HealthFeature),
		logistikHandler: logistik.NewLogistikHandler(container.LogistikFeature),
		salesHandler:    sales.NewSalesHandler(),
		userHandler:     user.NewUserHandler(container.UserFeature),
		carHandler:      car.NewCarHandler(container.CarFeature),
	}
}
