package routes


import (
	controller "RestaurantManagement/controllers"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(incomingroutes *fiber.App) error {
	incomingroutes.Get("/users", controller.GetUsers())
	incomingroutes.Get("/users/:user_id",controller.GetUser())
	incomingroutes.Post("/users/signup", controller.SignUp())
	incomingroutes.Post("/users/login", controller.Login())
	return nil
}


