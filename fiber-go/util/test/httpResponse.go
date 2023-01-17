package test

import (
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Ler o corpo da http.Response e retornar sua resposta como string
func ReadResponseBody(req *http.Response, app *fiber.App) (string, error) {

	data, _ := ioutil.ReadAll(req.Body)
	dataString := string(data)

	return dataString, nil
}
