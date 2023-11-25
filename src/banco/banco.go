package banco

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver MySQL
)

// Conectar abre a conexão com o banco de dados e a retorna
func Conectar() (*sql.DB, error) {
	// Abre uma conexão com o banco de dados usando o driver MySQL
	db, erro := sql.Open("mysql", config.StringConexaoBanco)

	if erro != nil {
		return nil, erro
	}

	// Verifica se a conexão com o banco de dados é bem-sucedida
	if erro = db.Ping(); erro != nil {
		// Se houver um erro, fecha a conexão e retorna o erro
		db.Close()
		return nil, erro
	}

	// Retorna a conexão com o banco de dados
	return db, nil
}
