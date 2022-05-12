package app

import "github.com/DarkSoul94/services/models"

type Usecase interface {
	AcceptNewTicket(newTicket models.Ticket) (string, error)
}
