package authentication

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/AlysonRyan01/template-go/internal/response"
)

type UserService interface {
	Register(ctx context.Context, dto *RegisterUserDto) (string, error)
}

type Handler struct {
	service UserService
}

func NewHandler(service UserService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var request RegisterUserDto
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		response.Error(w, http.StatusBadRequest, "JSON inválido")
		return 
	}

	if err := request.validate(); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
	}

	id, err := h.service.Register(r.Context(), &request)
	if err != nil {
		switch {
		case errors.Is(err, ErrEmailAlreadyExists):
			response.Error(w, http.StatusConflict, "e-mail já cadastrado")
		default:
			log.Printf("erro ao registrar usuário: %v", err)
			response.Error(w, http.StatusInternalServerError, "erro interno")
		}
		return
	}

	response.JSON(w, http.StatusCreated, map[string]string{"id": id})
}