package main

import (
	_ "RestaurantManagement/database"
	_ "RestaurantManagement/middleware"
	_ "RestaurantManagement/routes"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
)

func main() {
	port := os.Getenv("Port")

	if port == "" {
		port = "1988"
	}

	router := fiber.New()
	router.Use(logger.New())
}
