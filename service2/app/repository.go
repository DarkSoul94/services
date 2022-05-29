package app

import (
	"context"

	"github.com/DarkSoul94/services/models"
)

type Repository interface {
	InsertResult(ctx context.Context, ticket models.Ticket) error
}
