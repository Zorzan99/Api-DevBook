package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// StringConexaoBanco é a string de conexão com o MySQL
	StringConexaoBanco = ""
	// Porta onde a API estará rodando
	Porta = 0
)

// Carregar inicializa as variáveis de ambiente
func Carregar() {
	var erro error

	// Carrega as variáveis de ambiente do arquivo .env
	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	// Converte a variável de ambiente API_PORT para inteiro e armazena em Porta
	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))

	// Se houver erro na conversão ou se a variável não estiver definida, utiliza a porta padrão 9000
	if erro != nil {
		Porta = 9000
	}

	// Constrói a string de conexão com o banco de dados MySQL
	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)
}
