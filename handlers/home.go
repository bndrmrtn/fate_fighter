package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"github.com/bndrmrtn/fate_fighter/database"
	"github.com/bndrmrtn/fate_fighter/ui/views"
)

type home struct{}

var Home home

var ctx = context.Background()

func (home) Index(c *fiber.Ctx) error {
	client := clientCtx(c)

	data, err := database.DB.GetFirstServer(ctx)
	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return view(c, views.Home(client, data.Hostname.String, data.Address))
}

func (home) ToggleDarkMode(c *fiber.Ctx) error {
	var val string
	if c.Params("enabled") == "true" {
		val = "1"
	} else {
		val = "0"
	}

	c.Cookie(&fiber.Cookie{
		Name:  "dark_mode",
		Value: val,
	})
	return nil
}
