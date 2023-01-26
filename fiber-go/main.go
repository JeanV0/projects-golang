package main

import (
	"fmt"
	"log"
	dotenv "marketplace/config/services/Dotenv"
	"marketplace/config/services/db"
	"marketplace/src"
)

// Inicialização do servidor
func main() {
	dotenv.LoadEnv()
	db.GetConnectionDB()

	// TODO: Em vez de get connection e outros, quero criar uma funcao que ja testa conexao e tudo necessario. Para que se tiver já com algum problema nao inicialize a aplicacao

	// Pegar o router app do fiber com outra função
	src.GetApp()

	// Inicializando servidor
	URI_API := fmt.Sprintf(":%s", dotenv.MyEnvironmentApp.Api_port)
	log.Fatal(src.App.Listen(URI_API))

}
