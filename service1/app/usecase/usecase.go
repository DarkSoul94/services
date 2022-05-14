package usecase

import (
	"encoding/json"

	"github.com/DarkSoul94/services"
	"github.com/DarkSoul94/services/models"
	queueclient "github.com/DarkSoul94/services/pkg/QueueClient"
	"github.com/DarkSoul94/services/service1/app"
	"github.com/google/uuid"
)

type usecase struct {
	qCli queueclient.QueueClient
	repo app.Repository
}

func NewUsecase(cli queueclient.QueueClient, r app.Repository) app.Usecase {
	return &usecase{
		qCli: cli,
		repo: r,
	}
}

func (u *usecase) AcceptNewTicket(newTicket models.Ticket) (string, error) {
	id := uuid.New().String()
	newTicket.ID = id

	data, err := json.Marshal(newTicket)
	if err != nil {
		return "", err
	}

	err = u.qCli.Publish(services.NewTicketQueueName, data)
	if err != nil {
		return "", err
	}

	return id, nil
}
