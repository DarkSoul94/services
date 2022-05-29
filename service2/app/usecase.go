package app

import (
	"context"

	"github.com/DarkSoul94/services/models"
)

type Usecase interface {
	TickerProcessing(ctx context.Context, ticket models.Ticket) error
}
