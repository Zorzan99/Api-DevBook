package modelos

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Usuario representa um usuário utilizando a rede social
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

// Preparar chamará os métodos validar e formatar para o usuário recebido
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}
	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	}
	return nil
}

// validar verifica se os campos obrigatórios estão preenchidos
func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}
	if usuario.Nick == "" {
		return errors.New("O nick é obrigatório e não pode estar em branco")
	}
	if usuario.Email == "" {
		return errors.New("O email é obrigatório e não pode estar em branco")
	}

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("O e-mail inserido é inválido")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A senha é obrigatória e não pode estar em branco")
	}
	return nil
}

// formatar remove espaços em branco extras dos campos do usuário
// formatar é um método do tipo Usuario que formata os campos do usuário.
// Ele recebe uma etapa como parâmetro para determinar as ações específicas a serem executadas.
func (usuario *Usuario) formatar(etapa string) error {
	// Remover espaços em branco dos campos do usuário
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	// Verificar se a etapa é "cadastro" para realizar a formatação adicional durante o processo de cadastro
	if etapa == "cadastro" {
		// Gerar o hash da senha utilizando a função Hash do pacote seguranca
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}
		// Atualizar a senha do usuário com o hash gerado
		usuario.Senha = string(senhaComHash)
	}

	// Retornar nil para indicar que a formatação foi concluída sem erros
	return nil
}
