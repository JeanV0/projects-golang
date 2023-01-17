package rotas

import (
	"DevBookApi/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Struct que representa rotas da api
type Rota struct {
	URI    string
	Metodo string
	// Funcao             func(http.ResponseWriter, *http.Request)
	Funcao             http.HandlerFunc
	RequerAutenticacao bool
}

// Colocar todas as rotas dentro do router
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotaLogin)
	rotas = append(rotas, rotasPublicacoes...)

	for _, rota := range rotas {
		if rota.RequerAutenticacao {
			r.HandleFunc(rota.URI, middlewares.Autenticar(rota.Funcao)).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
		}

	}

	return r
}
