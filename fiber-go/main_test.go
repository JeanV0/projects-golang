package main

import (
	"marketplace/Tests/middleware"
	dotenv "marketplace/config/services/Dotenv"
	db "marketplace/config/services/db"
	"marketplace/src"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func getDb() {
	dotenv.LoadEnv()
	db.GetConnectionDB()
}

func getApp() *fiber.App {

	err := db.GetConnectionDB()

	if err != nil {
		panic(err.Error())
	}

	dotenv.LoadEnv()
	db.GetConnectionDB()

	// Pegar o router app do fiber com outra função
	src.GetApp()

	return src.App
}

func TestMiddleware(t *testing.T) { app := getApp(); middleware.TestMiddleware(app, t) }
