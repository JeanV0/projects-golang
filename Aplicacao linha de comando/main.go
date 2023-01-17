package main

import (
	"linhaDeComando/app"
	"log"
	"os"
)

func main() {
	application := app.Gerar()
	erro := application.Run(os.Args)
	if erro != nil {
		log.Fatal(erro)
	}

}
