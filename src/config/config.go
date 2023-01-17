package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Variaveis de acesso para banco de dados
var (
	StringConexao = ""
	Porta         = 0
	SecretKey     []byte
)

// Carregar vai inicializar as variaveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))

	if erro != nil {
		Porta = 9000
	}

	StringConexao = fmt.Sprintf("%s:%s@/%s?charset=utf8",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))

}
