package controllers

import "net/http"

//CriarUsuario insere um usuario no DB
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuario!"))
}

//BuscarUuarios Busca todos os usuarios no DB
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuarios"))
}

//BuscarUuario Busca um usuario no DB
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuario!"))
}

//AtualizarUsuario Atualiza um usuario no DB
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando um usuario!"))
}

//DeletarUusuario Deleta as informacoes de um usuario no DB
func DeletarUusuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuario!"))
}
