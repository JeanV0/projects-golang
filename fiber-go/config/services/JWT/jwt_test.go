package jwt

import (
	dotenv "marketplace/config/services/Dotenv"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type testCase struct {
	user   string
	claims jwt.MapClaims
	token  string
}

var MyTestCase = []testCase{
	{
		user: "Pudim",
		claims: jwt.MapClaims{
			"exp":               time.Now().Add(60 * time.Minute).Unix(),
			"authorized":        true,
			"AuthorizationType": "ADMIN",
			"User":              "Pudim",
		},
	},
}

func TestCreateValidateToken(t *testing.T) {
	dotenv.LoadEnv()

	service := JwtService()

	for key, TestActual := range MyTestCase {
		token, err := service.NewToken(TestActual.user, TestActual.claims)

		if err != nil {
			t.Fatalf("\nErro na criação do token para usuario %s \nR: %s", TestActual.user, err.Error())
		}

		MyTestCase[key].token = token

		t.Logf("\nToken create!\nUserio: %s\nToken: %s", TestActual.user, MyTestCase[key].token)
	}

}

func TestValidationToken(t *testing.T) {
	dotenv.LoadEnv()

	service := JwtService()

	for _, TestActual := range MyTestCase {
		_, err := service.TokenValidate(TestActual.token, TestActual.claims)
		if err != nil {
			t.Fatalf("Test token validate error\nUser: %s\nError: %s", TestActual.user, err.Error())
		}
		t.Logf("Token validate!\nUser: %s\nToken: %s", TestActual.token, TestActual.user)
	}
}
