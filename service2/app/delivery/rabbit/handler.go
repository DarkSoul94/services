package rabbit

import (
	"context"
	"encoding/json"

	"github.com/DarkSoul94/services/models"
	"github.com/DarkSoul94/services/service2/app"
	"github.com/streadway/amqp"
)

type Handler struct {
	uc app.Usecase
}

func NewHandler(uc app.Usecase) *Handler {
	return &Handler{
		uc: uc,
	}
}

func (h *Handler) AcceptTicket(msg amqp.Delivery) {
	var ticket models.Ticket

	err := json.Unmarshal(msg.Body, &ticket)
	if err != nil {
		return
	}

	err = h.uc.TickerProcessing(context.Background(), ticket)
	if err != nil {
		return
	}

	msg.Ack(false)
}
