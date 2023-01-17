package repositorios

import (
	"DevBookApi/src/model"
	"database/sql"
)

type Publicacoes struct {
	db *sql.DB
}

func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

func (repositorio Publicacoes) Criar(publicacao model.Publicacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare(`
		insert into publicacoes (titulo, conteudo, autor_id) values (? , ? , ?)
	`)

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)

	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()

	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil

}

func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (model.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
		select p.*, u.nick from
		publicacoes p inner join usuarios u
		on u.id = p.autor_id where p.id = ?
	`, publicacaoID)

	var publicacao model.Publicacao

	if erro != nil {
		return publicacao, erro
	}

	if linhas.Next() {
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return publicacao, erro
		}
	}

	return publicacao, nil
}

func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]model.Publicacao, error) {
	linhas, erro := repositorio.db.Query(`
	select distinct p.*, u.nick from publicacoes p
	inner join usuarios u on u.id = p.autor_id
	inner join seguidores s on p.autor_id = s.usuario_id
	where u.id = ? or s.seguidor_id = ?
	`, usuarioID, usuarioID)

	var publicacoes []model.Publicacao

	if erro != nil {
		return publicacoes, erro
	}

	defer linhas.Close()

	for linhas.Next() {
		var publicacao model.Publicacao
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick,
		); erro != nil {
			return []model.Publicacao{}, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	return []model.Publicacao{}, erro
}

func (repositorio Publicacoes) Atualizar(publicacaoID uint64, publicacao model.Publicacao) error {
	statement, erro := repositorio.db.Prepare("update publicacoes set titulo = ?, conteudo = ? where id = ?")

	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacaoID); erro != nil {
		return erro
	}

	return nil
}
