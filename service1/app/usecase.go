package app

import (
	"context"

	"github.com/DarkSoul94/services/models"
)

type Usecase interface {
	Registration(ctx context.Context, user models.User) (string, error)
	GetUserList(ctx context.Context) ([]models.User, error)

	AcceptNewTicket(ctx context.Context, newTicket models.Ticket) (string, error)
	GetTicketList(ctx context.Context, userID string) ([]models.Ticket, error)
}
