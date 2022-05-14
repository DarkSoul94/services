package models

type Ticket struct {
	ID     string
	Email  string
	UserID string
	Result bool
}

//Ticket status
const (
	Accept int = iota
	Processing
	Complete
)
