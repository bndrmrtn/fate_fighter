package handlers

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"

	"github.com/bndrmrtn/fate_fighter/pkg/client"
	"github.com/bndrmrtn/fate_fighter/pkg/i18n"
)

var translator *i18n.I18n

type ClientCookie struct {
	DarkMode bool   `cookie:"dark_mode"`
	Language string `cookie:"language"`
}

func clientCtx(c *fiber.Ctx) *client.Context {
	if translator == nil {
		translator = i18n.NewI18n(&i18n.Config{
			FallbackLanguage: "hu",
		})
	}

	var cookie ClientCookie
	c.CookieParser(&cookie)

	return &client.Context{
		ITranslator: translator.GetTranslator(cookie.Language),
		DarkMode:    cookie.DarkMode,
	}
}

func view(c *fiber.Ctx, component templ.Component) error {
	handler := adaptor.HTTPHandler(templ.Handler(component))
	return handler(c)
}
