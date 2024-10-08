package main

import (
	"log"

	"github.com/Akkshatt/fiber_go/database"
	"github.com/Akkshatt/fiber_go/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("welcome to my api")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api", welcome)
	// user endpointss
	app.Post(
		"/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)
	// products
	app.Post("/api/products", routes.CreateProduct)
	app.Get("/api/products", routes.GetProducts)
	app.Get("/api/products/:id", routes.GetProduct)
	app.Put("/api/products/:id", routes.UpdateProduct)
	app.Delete("/api/products/:id", routes.DeleteProduct)
	//users
	app.Post("/api/orders",routes.CreateOrder)
	app.Get("/api/orders",routes.GetOrders)
	app.Get("/api/orders/:id",routes.GetOrder)


}

func main() {
	database.ConnectDb()
	app := fiber.New()
	// app.Get(
	// 	"/api",
	// 	welcome,
	// )
	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
