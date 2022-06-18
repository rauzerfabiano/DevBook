package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

// Usuarios é um repositório de usuários
type Usuarios struct {
	db *sql.DB
}

// Função cria um repositório de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um usuário no banco de dados
func (u Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	return 0, nil
}
