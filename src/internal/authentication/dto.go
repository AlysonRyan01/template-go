package authentication

import (
	"errors"
	"net/mail"
)

type RegisterUserDto struct {
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Phone    *string `json:"phone"`
}

func (d RegisterUserDto) validate() error {
	addr, err := mail.ParseAddress(d.Email)
	if err != nil || addr.Address != d.Email {
		return errors.New("e-mail inválido")
	}

	if len(d.Password) < 8 {
		return errors.New("a senha deve ter no mínimo 8 caracteres")
	}
	if len(d.Password) > 72 {
		return errors.New("a senha deve ter no máximo 72 caracteres")
	}

	if d.Phone != nil {
		if len(*d.Phone) < 10 || len(*d.Phone) > 13 {
			return errors.New("telefone inválido")
		}
		for _, c := range *d.Phone {
			if c < '0' || c > '9' {
				return errors.New("telefone deve conter apenas números")
			}
		}
	}

	return nil
}