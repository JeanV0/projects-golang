package jwt

import (
	"errors"
	"fmt"
	dotenv "marketplace/config/services/Dotenv"

	"github.com/golang-jwt/jwt/v4"
)

// A base do nosso service
type jwtService struct {
}

// Retornar jwtService
func JwtService() jwtService {
	return jwtService{}
}

// Criação de novo token
func (service jwtService) NewToken(user string, claims jwt.MapClaims) (string, error) {

	// Criar um token com uma claim personalizada
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Chave secreta para criar o token
	secretKey := []byte(dotenv.MyEnvironmentApp.Jwt_Secret)

	// Criação do token para string
	tokenString, err := token.SignedString(secretKey)

	return tokenString, err
}

func (service jwtService) TokenValidate(token string, claim jwt.MapClaims) (*jwt.Token, error) {

	tokenUser, err := jwt.ParseWithClaims(token, claim, service.secretKey)
	_, ok := tokenUser.Claims.(jwt.MapClaims)

	if !ok {
		return &jwt.Token{}, errors.New("Error no claims")
	}

	return tokenUser, err
}

func (service jwtService) secretKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(dotenv.MyEnvironmentApp.Jwt_Secret), nil
}
