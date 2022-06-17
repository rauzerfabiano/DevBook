package controllers

import "net/http"

// CriarUsuario inserir um usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criar usuario"))
}

// BuscarUsuarios busca todos os usuarios do banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuarios"))
}

// BuscarUsuario busca um usuario do banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Busca um usuário específico"))
}

// AtualizarUsuario altera as informações de um usuario no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuario"))
}

//DeletarUsuario exclui as informações um usuario do banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuario"))
}
