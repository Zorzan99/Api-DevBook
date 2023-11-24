package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

// Gerar vai retornar um router com as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter() // Cria um novo router usando o pacote gorilla/mux

	return rotas.Configurar(r) // Configura o router chamando a função Configurar do pacote de rotas
}
