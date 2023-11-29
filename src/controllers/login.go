package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"api/src/seguranca"
	"encoding/json"
	"io"
	"net/http"
)

// Login é responsável por autenticar um usuário na API
func Login(w http.ResponseWriter, r *http.Request) {
	// Ler o corpo da requisição
	corpoRequisicao, erro := io.ReadAll(r.Body)

	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	// Criar uma instância da struct Usuario para armazenar os dados do JSON
	var usuario modelos.Usuario

	// Fazer o Unmarshal do corpo da requisição para a struct Usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	// Conectar ao banco de dados
	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close() // Fechar a conexão com o banco de dados no final da função

	// Criar um novo repositório de usuários, passando o banco de dados como parâmetro
	repositorio := repositorios.NovoRepositorioDeUsuarios(db)

	// Buscar o usuário no banco de dados pelo email
	usuarioSalvoNoBanco, erro := repositorio.BuscarPorEmail(usuario.Email)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	// Verificar se a senha fornecida corresponde à senha armazenada no banco de dados
	if erro = seguranca.VerificarSenha(usuarioSalvoNoBanco.Senha, usuario.Senha); erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	// Se tudo estiver correto, enviar uma resposta indicando que o usuário está logado

	token, erro := autenticacao.CriarToken(usuarioSalvoNoBanco.ID)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	w.Write([]byte(token))
}
