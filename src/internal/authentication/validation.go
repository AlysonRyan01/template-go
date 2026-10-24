package authentication

import "errors"

var (
	ErrUserNotFound       = errors.New("usuário não encontrado")
	ErrEmailAlreadyExists = errors.New("e-mail já cadastrado")
	ErrCheckEmail         = errors.New("erro ao verificar e-mail")
	ErrCreateUser         = errors.New("erro ao salvar usuário")
	ErrHashPassword       = errors.New("erro ao gerar hash da senha")
)
