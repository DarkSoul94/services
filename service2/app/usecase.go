package app

import "github.com/DarkSoul94/services/models"

type Usecase interface {
	TickerProcessing(ticket models.Ticket) error
}
