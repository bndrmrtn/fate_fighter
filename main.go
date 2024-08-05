package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/bndrmrtn/fate_fighter/database"
	"github.com/bndrmrtn/fate_fighter/handlers"
	"github.com/bndrmrtn/fate_fighter/pkg/configs"
)

func main() {
	err := configs.LoadEnv()
	if err != nil {
		log.Fatal(err)
	}

	database.MustConnect()

	app := fiber.New()

	app.Static("/", "./public", fiber.Static{
		CacheDuration: 0 * time.Second,
	})

	app.Use(logger.New())
	app.Get("/", handlers.Home.Index)
	app.Post("/dark-mode/:enabled<bool>", handlers.Home.ToggleDarkMode)

	log.Fatal(app.Listen(":8001"))
}
