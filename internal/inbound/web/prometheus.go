package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewMonitoringHandler() *MonitoringHandler {
	return &MonitoringHandler{}
}

type MonitoringHandler struct{}

func (MonitoringHandler) MetricsEndpoint() fiber.Handler {
	return adaptor.HTTPHandler(promhttp.Handler())
}
