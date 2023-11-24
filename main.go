package main

import (
	"api/src/router" // Importa o pacote de roteamento personalizado
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Rodando api") // Imprime uma mensagem indicando que a API está rodando

	r := router.Gerar() // Gera as rotas utilizando a função Gerar() do pacote de roteamento

	log.Fatal(http.ListenAndServe(":5000", r)) // Inicia o servidor na porta 5000 usando as rotas geradas
}
