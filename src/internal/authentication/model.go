package authentication

type User struct {
	ID string
	Email string
	PasswordHash string
	Phone *string
}