package models

import "github.com/google/uuid"

type Notification struct {
	TicketID uuid.UUID
	NotificationType
	Text string
}

type NotificationType int

const (
	Email NotificationType = iota
)
