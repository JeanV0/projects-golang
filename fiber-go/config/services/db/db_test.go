package db

import (
	dotenv "marketplace/config/services/Dotenv"
	"testing"
)

// Primeiro teste apenas para conexão
func TestDBConnection(t *testing.T) {
	dotenv.LoadEnv()

	err := GetConnectionDB()

	if err != nil {
		t.Error(err.Error())
	}

}

func TestDBrometheus(t *testing.T) {
	dotenv.LoadEnv()

	err := GetConnectionDB()

	if err != nil {
		t.Error(err.Error())
	}
}
