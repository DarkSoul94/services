package app

import (
	"context"

	"github.com/DarkSoul94/services/models"
)

type Repository interface {
	CreateUser(ctx context.Context, user models.User) error
	CheckEmailExist(ctx context.Context, email string) bool
	GetUserList(ctx context.Context) ([]models.User, error)

	CreateTicket(ctx context.Context, ticket models.Ticket) error
	GetTicketList(ctx context.Context, userID string) ([]models.Ticket, error)
}
