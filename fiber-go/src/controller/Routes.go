package routes

import (
	"github.com/gofiber/fiber/v2"
)

type RoutesApp struct {
	Method        string
	RouteFunction func(*fiber.App) error
	Path          string
}

func configMiddlewareSecurity(app *fiber.App) {
	// Middleware de autenticação
}

// Configuração das rotas da aplicação
func ConfigRoutesApp(app *fiber.App) {
	// app.Add(http.MethodGet, "/hello", Hello)
	configMiddlewareSecurity(app)

	// metrics.GetPrometheusFiber(app) Uma futura ideia
}
