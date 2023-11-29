package rotas

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa todas as rotas da api
type Rota struct {
	Uri                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar coloca todas as rotas dentro do router
func Configurar(r *mux.Router) *mux.Router {
	// Obtém a lista de rotas a serem configuradas.
	// Assume que há uma variável chamada rotasUsuarios que contém um array de objetos do tipo Rota.
	rotas := rotasUsuarios
	rotas = append(rotas, rotaLogin)

	// Para cada rota, configura o router com a função associada e outros middlewares, se necessário.
	for _, rota := range rotas {
		// Verifica se a rota requer autenticação.
		if rota.RequerAutenticacao {
			// Se sim, envolve a função manipuladora com o middleware de autenticação e o middleware de log.
			r.HandleFunc(rota.Uri, middlewares.Logger(middlewares.Autenticar(rota.Funcao))).Methods(rota.Metodo)
		} else {
			// Se não, envolve apenas com o middleware de log.
			r.HandleFunc(rota.Uri, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}
		// Para cada rota, adiciona um handler para a URI e método especificados,
		// que chama a função associada à rota quando a rota é acessada.
	}

	// Retorna o router configurado.
	return r
}
