package app

import (
	"context"

	"github.com/DarkSoul94/services/models"
)

type Usecase interface {
	Registration(ctx context.Context, user models.User) (string, error)
	AcceptNewTicket(newTicket models.Ticket) (string, error)
}
