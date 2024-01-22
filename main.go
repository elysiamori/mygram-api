package main

import (
	"github.com/elysiamori/mygram-api/api/app"
	"github.com/elysiamori/mygram-api/config"
	"github.com/gofiber/fiber/v2"
)

func main() {

	db, err := config.Config()
	if err != nil {
		panic(err)
	}

	app1 := fiber.New()

	app.SetupRoutes(app1, db)

	err = app1.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
