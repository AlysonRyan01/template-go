package authentication

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, m *User) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO users (id, email, password_hash, phone)
		VALUES ($1, $2, $3, $4)`, m.ID, m.Email, m.PasswordHash, m.Phone)
	if err != nil {
		return fmt.Errorf("inserir usuário: %w", err)
	}

	return nil
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (User, error) {
	var user User
	err := r.db.QueryRowContext(ctx, `
		SELECT id, email, password_hash, phone FROM users 
		WHERE email = $1`, email).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Phone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return User{}, ErrUserNotFound
		}
		return User{}, fmt.Errorf("buscar usuário por e-mail: %w", err)
	}
	return user, nil
}