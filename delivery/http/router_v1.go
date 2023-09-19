package http

import (
	"github.com/gofiber/fiber/v2"
)

func RouterGroupV1(app *fiber.App, handler handler) {

	v1 := app.Group("/v1")

	// authProduct := v1.Group("/product")
	// {
	// 	authProduct.Use(middleware.AuthValidations())
	// 	authProduct.Post("/add", handler.logistikHandler.AddProductHandler)
	// 	authProduct.Put("/update/:id", handler.logistikHandler.UpdateProductHandler)
	// 	authProduct.Delete("/delete/:id", handler.logistikHandler.DeleteProductHandler)
	// }

	userFeature := v1.Group("/user")
	{
		userFeature.Post("/add", handler.userHandler.AddUserHandler)
		userFeature.Put("/update/:id", handler.userHandler.UpdateUserHandler)
		userFeature.Delete("/delete/:id", handler.userHandler.DeleteUserHandler)
		userFeature.Get("/lists", handler.userHandler.GetUserListsHandler)
	}

	carFeature := v1.Group("/car")
	{
		carFeature.Post("/add", handler.carHandler.AddCarHandler)
		carFeature.Put("/update/:id", handler.carHandler.UpdateCarHandler)
		carFeature.Delete("/delete/:id", handler.carHandler.DeleteCarHandler)
		carFeature.Get("/lists", handler.carHandler.GetCarListsHandler)
	}

}
