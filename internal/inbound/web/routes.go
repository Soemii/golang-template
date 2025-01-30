package web

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App, monitoringHandler *MonitoringHandler) {
	app.Get("/metrics", monitoringHandler.MetricsEndpoint())
}
