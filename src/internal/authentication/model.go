package authentication

import "time"

type User struct {
	ID           string
	Email        string
	PasswordHash string
	Phone        *string
	Role         Role
}

type RefreshToken struct {
	ID        string
	TokenHash string
	UserID    string
	ExpiresAt time.Time
	RevokedAt *time.Time
	CreatedAt time.Time
}

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)