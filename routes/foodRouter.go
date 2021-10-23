package routes

import (
	controller "RestaurantManagement/controllers"
	"github.com/gofiber/fiber/v2"
)

func FoodRoutes(incomingRoutes *fiber.App) error {
	incomingRoutes.Get("/foods", controller.GetFoods())
	incomingRoutes.Get("/foods/:food_id", controller.GetFood())
	incomingRoutes.Post("/foods", controller.CreateFood())
	incomingRoutes.Patch("/foods/:food_id", controller.UpdateFood())

}
