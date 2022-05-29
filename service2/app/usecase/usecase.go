package usecase

import (
	"context"
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
	repo app.Repository
}

func NewUsecase(qCli queueclient.QueueClient, repo app.Repository) app.Usecase {
	return &usecase{
		qCli: qCli,
		repo: repo,
	}
}

func (u *usecase) TickerProcessing(ctx context.Context, ticket models.Ticket) error {
	time.Sleep(1 * time.Second)

	res, err := rand.Int(rand.Reader, big.NewInt(1))
	if err != nil {
		return err
	}

	notification := models.Notification{
		TicketID:         ticket.ID,
		NotificationType: models.Email,
	}

	if res.Int64() == 0 {
		ticket.Result = false
		notification.Text = "Sorry, result for your ticket is negative"
	} else {
		ticket.Result = true
		notification.Text = "Congratulations, result for your ticket is positive"
	}

	err = u.repo.InsertResult(ctx, ticket)
	if err != nil {
		return err
	}

	data, err := json.Marshal(notification)
	if err != nil {
		return err
	}

	return u.qCli.Publish(services.NotifyQueueName, data)
}
