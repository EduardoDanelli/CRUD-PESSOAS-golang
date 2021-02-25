package models

import (
	"errors"
	"strings"

	"github.com/badoux/checkmail"
)

// Pessoa representa um pessoa no banco de dados
type Pessoa struct {
	ID             int    `json:"id,omitempty" gorm:"primaryKey,serial"`
	Nome           string `json:"nome,omitempty"`
	CPF            string `json:"cpf,omitempty"`
	DataNascimento string `json:"dataNascimento,omitempty" gorm:"column:datanascimento"`
	Telefone       string `json:"telefone,omitempty"`
	Email          string `json:"email,omitempty"`
	Rua            string `json:"rua,omitempty"`
	Bairro         string `json:"bairro,omitempty"`
	Complemento    string `json:"complemento,omitempty"`
	Cidade         string `json:"cidade,omitempty"`
}

// Preparar chama metodos de validar e formatar
func (pessoa *Pessoa) Preparar() error {
	if err := pessoa.validar(); err != nil {
		return err
	}

	if err := pessoa.formatar(); err != nil {
		return err
	}

	return nil
}

func (pessoa *Pessoa) validar() error {
	if pessoa.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}
	if pessoa.CPF == "" {
		return errors.New("O CPF é obrigatório e não pode estar em branco")
	}

	if len(pessoa.CPF) != 11 {
		return errors.New("O CPF é inválido")
	}

	if pessoa.DataNascimento == "" {
		return errors.New("A data de nascimento é obrigatório e não pode estar em branco")
	}
	if pessoa.Telefone == "" {
		return errors.New("O telefone é obrigatório e não pode estar em branco")
	}
	if pessoa.Rua == "" {
		return errors.New("A rua é obrigatório e não pode estar em branco")
	}
	if pessoa.Bairro == "" {
		return errors.New("O bairro é obrigatório e não pode estar em branco")
	}
	if pessoa.Cidade == "" {
		return errors.New("O código da cidade é obrigatório e não pode estar em branco")
	}

	if pessoa.Email != "" {
		if err := checkmail.ValidateFormat(pessoa.Email); err != nil {
			return errors.New("O email inserio não é válido")
		}
	}

	return nil
}

func (pessoa *Pessoa) formatar() error {
	pessoa.Nome = strings.TrimSpace(pessoa.Nome)
	pessoa.CPF = strings.TrimSpace(pessoa.CPF)
	pessoa.DataNascimento = strings.TrimSpace(pessoa.DataNascimento)
	pessoa.Telefone = strings.TrimSpace(pessoa.Telefone)
	pessoa.Email = strings.TrimSpace(pessoa.Email)
	pessoa.Rua = strings.TrimSpace(pessoa.Rua)
	pessoa.Bairro = strings.TrimSpace(pessoa.Bairro)
	pessoa.Complemento = strings.TrimSpace(pessoa.Complemento)
	pessoa.Cidade = strings.TrimSpace(pessoa.Cidade)

	return nil
}
