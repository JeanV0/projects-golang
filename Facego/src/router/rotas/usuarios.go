package rotas

import (
	"DevBookApi/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuarios,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios/{usuarioID}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUmUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios/{usuarioID}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios/{usuarioID}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios/{usuarioId}/seguir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.SeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{usuarioId}/parar-de-seguir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.PararDeSeguirUsuario,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{usuarioId}/seguidores",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarSeguidores,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{usuarioId}/seguindo",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarSeguindo,
		RequerAutenticacao: true,
	},
	{
		URI:                "/usuarios/{usuarioId}/atualizar-senha",
		Metodo:             http.MethodPost,
		Funcao:             controllers.AtualizarSenha,
		RequerAutenticacao: true,
	},
}
