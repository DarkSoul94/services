package usecase

import (
	"crypto/rand"
	"encoding/json"
	"math/big"
	"time"

	"github.com/DarkSoul94/services"
	"github.com/DarkSoul94/services/models"
	queueclient "github.com/DarkSoul94/services/pkg/QueueClient"
	"github.com/DarkSoul94/services/service2/app"
)

type usecase struct {
	qCli queueclient.QueueClient
}

func NewUsecase(qCli queueclient.QueueClient) app.Usecase {
	return &usecase{
		qCli: qCli,
	}
}

func (u *usecase) TickerProcessing(ticket models.Ticket) error {
	time.Sleep(1 * time.Second)

	res, err := rand.Int(rand.Reader, big.NewInt(1))
	if err != nil {
		return err
	}

	notification := models.Notification{
		Email: ticket.Email,
	}

	if res.Int64() == 0 {
		ticket.Result = false
		notification.Text = "Sorry, result for your ticket is negative"
	} else {
		ticket.Result = true
		notification.Text = "Congratulations, result for your ticket is positive"
	}

	data, err := json.Marshal(notification)
	if err != nil {
		return err
	}

	return u.qCli.Publish(services.NotifyQueueName, data)
}
