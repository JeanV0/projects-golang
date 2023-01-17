package src

import (
	router "marketplace/src/controller"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

var App *fiber.App

// Criar fiber com as rotas
func GetApp() {

	App = fiber.New(fiber.Config{
		// Nome da aplicação
		AppName: "Golang pigz 0.1",
		// No header do html em vez de nginx e tals
		ServerHeader: "Pigz",
		// Validação do ip se é verdadeiro ou não
		EnableIPValidation: true,
		// Pre bifurcação
		Prefork:     true,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// Configuração de rotas da aplicação
	router.ConfigRoutesApp(App)

}
