package authentication

import "strings"

func RegisterUserDtoToUser(d *RegisterUserDto) User {
	return User{
		Email: strings.ToLower(strings.TrimSpace(d.Email)),
		Phone: d.Phone,
	}
}