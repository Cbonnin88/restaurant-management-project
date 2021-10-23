package main

import (


	"RestaurantManagement/middleware"
	"RestaurantManagement/routes"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
)

var foodCollection, err = sql.Open("mysql", "root:pwd@tcp("+
	"localhost:3306)/restaurant_management")

func main() {
	port := os.Getenv("Port")

	if port == "" {
		port = "1988"
	}

	router := fiber.New()
	router.Use(logger.New())
	
	err := routes.UserRoutes(router)
	if err != nil {
		panic(err)
		return
	}
	router.Use(middleware.Authentication())

	err = routes.FoodRoutes(router)
	if err != nil {
		fmt.Print(err)
		return
	}
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItems(router)
	routes.InvoiceRoutes(router)


	err = router.Listen(":" + port)
	if err != nil {
		fmt.Print("Error: Unable to show page")
		return
	}



}
