package autenticacao

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// CriarToken cria um token para o usuário com as permissões do usuário
func CriarToken(usuarioID uint64) (string, error) {
	// Criar um mapa de claims (reivindicações) para o token
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix() // Token expira em 6 horas
	permissoes["usuarioId"] = usuarioID

	// Criar um novo token com as claims e o método de assinatura HMAC
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)

	// Assinar o token com a chave secreta configurada no arquivo de configuração
	return token.SignedString([]byte(config.SecretKey))
}

// ValidarToken verifica se o token passado na requisição é válido
func ValidarToken(r *http.Request) error {
	// Extrair o token da requisição
	tokenString := extrairToken(r)

	// Parse do token usando a função de verificação de chave
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return erro
	}

	// Verificar se o token é válido
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("Token inválido")
}

// extrairToken extrai o token da requisição HTTP
func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	// Verificar se o token está no formato esperado (Bearer <token>)
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

// retornarChaveDeVerificacao retorna a chave secreta para verificar a assinatura do token
func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	// Verificar se o método de assinatura é HMAC
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura inesperado! %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}
