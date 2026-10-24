package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/AlysonRyan01/template-go/internal/authentication"
	"github.com/AlysonRyan01/template-go/internal/config"
	"github.com/go-chi/chi/v5"

	_ "github.com/jackc/pgx/v5/stdlib"
)


func RegisterAuthenticationDependencies(db *sql.DB, r chi.Router) {
	repo := authentication.NewRepository(db)
	service := authentication.NewService(repo)
	handler := authentication.NewHandler(service)
	handler.RegisterRoutes(r)
}

func RegisterDB(ctx context.Context, cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("pgx", cfg.ConnectionString())
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping db: %w", err)
	}

	return db, nil
}