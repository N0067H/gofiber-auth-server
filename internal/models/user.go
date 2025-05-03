package models

import (
	"context"

	"github.com/n0067h/gofiber-auth-server/ent"
	"github.com/n0067h/gofiber-auth-server/ent/user"
	"github.com/n0067h/gofiber-auth-server/internal/apperror"
)

func CreateUser(ctx context.Context, client *ent.Client, email, name, password string) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetEmail(email).
		SetName(name).
		SetPassword(password).
		Save(ctx)

	if err != nil {
		if ent.IsConstraintError(err) {
			return nil, apperror.ErrEmailExists
		}

		return nil, apperror.ErrInternalServerError
	}

	return u, nil
}

func ReadUser(ctx context.Context, client *ent.Client, email string) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.EmailEqualFold(email)).
		First(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, apperror.ErrUserNotFound
		}

		return nil, apperror.ErrInternalServerError
	}

	return u, nil
}
