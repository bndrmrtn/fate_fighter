package main

import (
	"log"

	"github.com/a-h/templ"
	"github.com/bndrmrtn/fate_fighter/ui/views"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func main() {
	app := fiber.New()

	app.Static("/", "./public")

	component := views.Home()

	app.Get("/", adaptor.HTTPHandler(templ.Handler(component)))

	log.Fatal(app.Listen(":3000"))
}
