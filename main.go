package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	fmt.Println(config.SecretKey)

	r := router.Gerar() // Gera as rotas utilizando a função Gerar() do pacote de roteamento

	fmt.Printf("Escutando na porta %d", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r)) // Inicia o servidor na porta 5000 usando as rotas geradas
}
