package routes

import (
	dotenv "marketplace/config/services/Dotenv"
	apicontroller "marketplace/src/controller/ApiController"
	api "marketplace/src/security/Api"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type RoutesApp struct {
	Method        string
	RouteFunction func(*fiber.App) error
	Path          string
}

func configMiddlewareSecurity(app *fiber.App) {
	// Middleware de autenticação
	app.Use("/api", api.MiddlewareAuthenticatorApi)

	// File to logs
	file, err := os.OpenFile(dotenv.MyEnvironmentApp.Log_File, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		panic(err.Error())
	}
}

// Configuração das rotas da aplicação
func ConfigRoutesApp(app *fiber.App) {
	// app.Add(http.MethodGet, "/hello", Hello)
	configMiddlewareSecurity(app)
	apicontroller.RoutesHighlight(app)

	// metrics.GetPrometheusFiber(app) Uma futura ideia
}
