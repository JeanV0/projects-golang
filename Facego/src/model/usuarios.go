package model

import (
	"DevBookApi/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Struct que representa usuario na rede social
type Usuario struct {
	ID       uint      `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

// Validação se o campo estiver vazio
func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O nome é obrigatorio e não pode estar em branco!")
	}

	if usuario.Nick == "" {
		return errors.New("O nick é obrigatorio e não pode estar em branco@")
	}

	if usuario.Email == "" {
		return errors.New("O email é obrigatorio e não pode estar em branco!")
	}

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("O email tem formato invalido")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("O senha é obrigatoria e não pode estar em branco!")
	}

	return nil

}

// Formatar os campos
func (usuario *Usuario) Formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}

		usuario.Senha = string(senhaComHash)
	}

	return nil
}

// Preparar e validar os campos
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	usuario.Formatar(etapa)

	return nil
}
