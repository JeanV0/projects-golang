package controllers

import (
	"DevBookApi/src/autenticacao"
	"DevBookApi/src/banco"
	"DevBookApi/src/model"
	"DevBookApi/src/repositorios"
	"DevBookApi/src/respostas"
	"DevBookApi/src/seguranca"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	corpoDaRequisicao, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
	}

	var usuario model.Usuario
	if erro = json.Unmarshal(corpoDaRequisicao, &usuario); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	db, erro := banco.Conectar()

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarioSalvoNoBanco, erro := repositorio.BuscarPorEmail(usuario.Email)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = seguranca.VerificarSenha(usuarioSalvoNoBanco.Senha, usuario.Senha); erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
	}

	token, _ := autenticacao.CriarToken(uint64(usuarioSalvoNoBanco.ID))

	w.Write([]byte(token))

}
