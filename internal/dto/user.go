package dto

import (
	"github.com/n0067h/gofiber-auth-server/internal/apperror"
)

type RegisterRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Pass  string `json:"pass"`
}

type LoginRequest struct {
	Email string `json:"email"`
	Pass  string `json:"pass"`
}

func (r RegisterRequest) Validate() error {
	if r.Email == "" || r.Name == "" || r.Pass == "" {
		return apperror.ErrMissingFields
	}

	return nil
}

func (l LoginRequest) Validate() error {
	if l.Email == "" || l.Pass == "" {
		return apperror.ErrMissingFields
	}

	return nil
}
