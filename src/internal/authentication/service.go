package authentication

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userStore interface {
	Create(ctx context.Context, m *User) error
	GetUserByEmail(ctx context.Context, email string) (User, error)
}

type Service struct {
	repo userStore
}

func NewService(repo userStore) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Register(ctx context.Context, dto *RegisterUserDto) (string, error) {
	_, err := s.repo.GetUserByEmail(ctx, dto.Email)
	if err == nil {
		return "", ErrEmailAlreadyExists
	}
	if !errors.Is(err, ErrUserNotFound) {
		return "", fmt.Errorf("%w: %w", ErrCheckEmail, err)
	}

	user := RegisterUserDtoToUser(dto)
	hash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrHashPassword, err)
	}
	user.PasswordHash = string(hash)
	user.ID = uuid.NewString()

	if err := s.repo.Create(ctx, &user); err != nil {
		return "", fmt.Errorf("%w: %w", ErrCreateUser, err)
	}

	return user.ID, nil
}