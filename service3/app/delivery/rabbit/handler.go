package rabbit

import (
	"encoding/json"

	"github.com/DarkSoul94/services/models"
	"github.com/DarkSoul94/services/service3/app"
	"github.com/streadway/amqp"
)

type Handler struct {
	notificator app.Notificator
}

func NewHandler(n app.Notificator) *Handler {
	return &Handler{
		notificator: n,
	}
}

func (h *Handler) AcceptNotificationMsg(msg amqp.Delivery) {
	var notification models.Notification

	err := json.Unmarshal(msg.Body, &notification)
	if err != nil {
		return
	}
	h.notificator.Notify(notification)

	msg.Ack(false)
}
