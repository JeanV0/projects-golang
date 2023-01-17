package dotenv

import (
	"os"

	_ "github.com/joho/godotenv/autoload"

	"github.com/joho/godotenv"
)

// Struct de representação das variaveis da aplicação
type EnvironmentApp struct {
	Redis_host string
	Redis_port string

	Database_host     string
	Database_port     string
	Database_user     string
	Database_password string
	Database_DbName   string

	Api_host string
	Api_port string

	Jwt_Secret string

	Firebase_Key        string
	Firebase_Project_Id string
	Firebase_Storage    string

	Log_File string
}

// Variaveis do dotenv
var MyEnvironmentApp = EnvironmentApp{}

// Carregar arquivo .env ou .env.local
func LoadEnv() {

	err := godotenv.Load(".env.local")

	if err != nil {

		err = godotenv.Load(".env")

		if err != nil {
			panic(err.Error())
		}
	}

	// os.Getenv vai carregar as variaveis e salvar no struct atual

	// Variaveis de acesso ao redis
	MyEnvironmentApp.Redis_host = os.Getenv("REDIS_HOST")
	MyEnvironmentApp.Redis_port = os.Getenv("REDIS_PORT")

	// Variaveis de acesso ao banco de dados
	MyEnvironmentApp.Database_host = os.Getenv("DATABASE_HOST")
	MyEnvironmentApp.Database_port = os.Getenv("DATABASE_PORT")
	MyEnvironmentApp.Database_user = os.Getenv("DATABASE_USER")
	MyEnvironmentApp.Database_password = os.Getenv("DATABASE_PASSWORD")
	MyEnvironmentApp.Database_DbName = os.Getenv("DATABASE_DB")

	// Variaveis para configurar API
	MyEnvironmentApp.Api_host = os.Getenv("API_HOST")
	MyEnvironmentApp.Api_port = os.Getenv("API_PORT")

	// Configurações para jwt
	MyEnvironmentApp.Jwt_Secret = os.Getenv("JWT_SECRET_KEY")

	// Configurações para firebase
	MyEnvironmentApp.Firebase_Key = os.Getenv("FIREBASE_SECRET_API_KEY")
	MyEnvironmentApp.Firebase_Project_Id = os.Getenv("FIREBASE_PROJECT_ID")
	MyEnvironmentApp.Firebase_Storage = os.Getenv("FIREBASE_STORAGE_BUCKET")

	// Configuração de logs do projeto
	MyEnvironmentApp.Log_File = os.Getenv(("LOG_FILE"))
}
