package dotenv

import (
	"encoding/json"
	"fmt"
	"testing"
)

// Struct para teste. Analisar se as variaveis do dotenv é a mesma dessas
var StructToTestDotEnv = EnvironmentApp{
	Redis_host:        "172.22.0.2",
	Redis_port:        "6379",
	Database_host:     "172.22.0.3",
	Database_port:     "5432",
	Database_user:     "golang",
	Database_DbName:   "pigz_hom",
	Database_password: "golang",
	Api_host:          "127.0.0.1",
	Api_port:          "9000",
	Jwt_Secret:        "Z2JmZGFzamtsdmRocmVhZmR1c3JpZ2J3ZWl3aHNld3FwamlvcGR2Ymd0aWV3dWdpcmdmdWhvZ291aXVhaXVnd3VnT1VFSE9HQUVST1VHSEVPw4dIQUVST0dIT0VSR0/Dh0VSR09JRUhHT0lIR0lISE9oZG9IT0hPRUhHSU9HSE8",
	// Aparti de firebase parei de testar pois ja estava obvio que ja estava estabel a sua utilizaçãp
}

// Teste de usar dotenv se estar certo
func TestLoadEnv(t *testing.T) {
	LoadEnv()

	// Isso será um teste mais manual por sua complexidade e dificuldade de test com structs
	jsonBytes, _ := json.Marshal(StructToTestDotEnv)

	fmt.Println(string(jsonBytes))
	jsonBytes, _ = json.Marshal(MyEnvironmentApp)
	fmt.Println(string(jsonBytes))

	// Usar esses dois sites para analise rapida
	// https://jsonformatter.org/json-pretty-print
	// https://jsoncrack.com/

}
