package models

import "github.com/google/uuid"

type Ticket struct {
	ID     uuid.UUID
	UserID uuid.UUID
	Result bool
	Status TicketStatus
}

type TicketStatus int

const (
	Accept TicketStatus = iota
	Processing
	Complete
)

var TicketStatusString = map[TicketStatus]string{0: "Accept", 1: "Processing", 2: "Complete"}
