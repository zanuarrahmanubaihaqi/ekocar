package cmd

import (
	"eko-car/delivery/container"
	"eko-car/delivery/http"
	"fmt"
)

func Execute() {
	// start init container
	container := container.SetupContainer()

	// start http service
	http := http.ServeHttp(container)
	http.Listen(fmt.Sprintf(":%d", container.EnvironmentConfig.App.Port))
}
