package web

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func NewServer(lifecycle fx.Lifecycle) *fiber.App {
	app := fiber.New()
	lifecycle.Append(fx.StartStopHook(func() {
		go app.Listen(":8080")
	}, func() error {
		return app.Shutdown()
	}))
	return app
}
