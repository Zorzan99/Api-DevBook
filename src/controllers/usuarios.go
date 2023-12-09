package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CriarUsuario insere um usuario no DB
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	// Ler o corpo da requisição
	corpoRequest, erro := io.ReadAll(r.Body)

	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	// Converter o JSON do corpo da requisição para um objeto Usuario
	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	// Preparar o usuário para cadastro (por exemplo, criptografar senha)
	if erro = usuario.Preparar("cadastro"); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	// Conectar ao banco de dados
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	// Criar um novo repositório de usuários
	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	// Criar o usuário no banco de dados
	usuario.ID, erro = repositorio.Criar(usuario)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	// Responder com o usuário criado em formato JSON
	respostas.JSON(w, http.StatusCreated, usuario)
}

// BuscarUsuarios busca todos os usuarios no DB
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	// Obter o nome ou nick da consulta
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	// Conectar ao banco de dados
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	// Criar um novo repositório de usuários
	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	// Buscar usuários no banco de dados com base no nome ou nick
	usuarios, erro := repositorio.Buscar(nomeOuNick)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	// Responder com a lista de usuários em formato JSON
	respostas.JSON(w, http.StatusOK, usuarios)
}

// BuscarUsuario busca um usuario no DB
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	// Obter parâmetros da URL usando a biblioteca Gorilla Mux
	parametros := mux.Vars(r)

	// Converter o ID do usuário de string para uint64
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)

	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	// Conectar ao banco de dados
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	// Criar um novo repositório de usuários
	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	// Buscar o usuário no banco de dados pelo ID
	usuario, erro := repositorio.BuscarPorID(usuarioID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	// Responder com o usuário encontrado em formato JSON
	respostas.JSON(w, http.StatusOK, usuario)
}

// AtualizarUsuario atualiza um usuario no DB
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	// Obter parâmetros da URL usando a biblioteca Gorilla Mux
	parametros := mux.Vars(r)
	// Converter o ID do usuário de string para uint64
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)

	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	usuarioIDNoToken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	if usuarioID != usuarioIDNoToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possivel atualizar um usuario que não seja o seu"))
		return
	}

	// Ler o corpo da requisição
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	// Converter o JSON do corpo da requisição para um objeto Usuario
	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	// Conectar ao banco de dados
	db, erro := banco.Conectar()

	// Preparar o usuário para edição (por exemplo, criptografar senha)
	if erro = usuario.Preparar("edicao"); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	defer db.Close()

	// Criar um novo repositório de usuários
	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	// Atualizar as informações do usuário no banco de dados
	if erro = repositorio.Atualizar(usuarioID, usuario); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	// Responder com status No Content (204)
	respostas.JSON(w, http.StatusNoContent, nil)
}

// DeletarUsuario deleta as informacoes de um usuario no DB
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	// Obter parâmetros da URL usando a biblioteca Gorilla Mux
	parametros := mux.Vars(r)

	// Converter o ID do usuário de string para uint64
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	usuarioIDNoToken, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if usuarioID != usuarioIDNoToken {
		respostas.Erro(w, http.StatusForbidden, errors.New("Não é possivel deletar um usuário que não seja o seu"))
		return
	}

	// Conectar ao banco de dados
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	// Criar um novo repositório de usuários
	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	// Deletar o usuário no banco de dados
	if erro = repositorio.Deletar(usuarioID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	// Responder com status No Content (204)
	respostas.JSON(w, http.StatusNoContent, nil)
}
