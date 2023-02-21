package main

import (
	"customer-api/config"
	"customer-api/injector"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Could not load .env %v: ", err)
	}

	port := os.Getenv("PORT")
	conns := config.NewConn()
	customer := conns.CustomerConn()
	order := conns.OrderConn()
	product := conns.ProductConn()

	app := fiber.New()
	app.Use(cors.New())
	app.Static("/static", "./public", fiber.Static{Compress: true, Browse: true})

	api := app.Group("api")
	injector.NewAppInjector(api, customer)
	injector.NewOrderInjector(api, order)
	injector.NewProductInjector(api, product)

	fmt.Printf("⚡️[server]: Server is running on porting %s\n", port)

	log.Fatal(app.Listen(port))

}
