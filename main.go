package main

import (
	"DevBookApi/src/config"
	"DevBookApi/src/router"

	// "encoding/base64"
	"fmt"
	"log"

	// "math/rand"
	"net/http"
)

// func init() {
// 	chave := make([]byte, 64)
// 	if _, erro := rand.Read(chave); erro != nil {
// 		log.Fatal(erro)
// 	}

// 	stringBase64 := base64.StdEncoding.EncodeToString(chave)
// 	fmt.Println(stringBase64)
// }

func main() {
	config.Carregar()

	// fmt.Println(config.Porta)
	fmt.Printf("Rodando API na url http://localhost:%d!", config.Porta)

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
