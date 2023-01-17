package model

import (
	"errors"
	"strings"
	"time"
)

type Publicacao struct {
	ID        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorID   uint64    `json:"autorID,omitempty"`
	AutorNick string    `json:"autorNick,omitempty"`
	Curtidas  uint64    `json:"curtidas"`
	CriadaEm  time.Time `json:"criadaEm,omitempty"`
}

func (publicacao *Publicacao) Preparar() error {

	if erro := publicacao.Validar(); erro != nil {
		return erro
	}

	publicacao.Formatar()
	return nil
}

func (Publicacao *Publicacao) Validar() error {
	if Publicacao.Titulo == "" {
		return errors.New("O titulo é obrigatorio e não pode estar em branco")
	}

	if Publicacao.Conteudo == "" {
		return errors.New("O conteudo é obrigatorio e não pode estar em branco")
	}

	return nil
}

func (publicacao *Publicacao) Formatar() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
}
