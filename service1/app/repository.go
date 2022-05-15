package app

import (
	"context"

	"github.com/DarkSoul94/services/models"
)

type Repository interface {
	CreateUser(ctx context.Context, user models.User) error
	CheckEmailExist(ctx context.Context, email string) bool
}
