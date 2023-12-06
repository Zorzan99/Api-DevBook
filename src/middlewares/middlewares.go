package middlewares

import (
	"api/src/autenticacao"
	"api/src/respostas"
	"log"
	"net/http"
)

// Logger escreve informações da requisição no terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	// A função Logger retorna uma nova função que age como um middleware.
	return func(w http.ResponseWriter, r *http.Request) {
		// Registra informações da requisição no terminal.
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)

		// Chama a função manipuladora original (next) passando os mesmos argumentos.
		next(w, r)
	}
}

// Autenticar verifica se o usuário fazendo a requisição está autenticado
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	// A função Autenticar também retorna uma nova função middleware.
	return func(w http.ResponseWriter, r *http.Request) {
		// Adiciona um log indicando que a autenticação está ocorrendo.
		if erro := autenticacao.ValidarToken(r); erro != nil {
			// Se a autenticação falhar, retorna uma resposta de erro e interrompe a execução.
			respostas.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		// Se a autenticação for bem-sucedida, chama a função manipuladora original (proximaFuncao).
		proximaFuncao(w, r)
	}
}
