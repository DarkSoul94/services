package usecase

import (
	"context"
	"encoding/json"
	"errors"

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

func (u *usecase) Registration(ctx context.Context, user models.User) (string, error) {
	if u.repo.CheckEmailExist(ctx, user.Email) {
		return "", errors.New("User with this email already exist")
	}

	user.ID = uuid.New()

	err := u.repo.CreateUser(ctx, user)
	if err != nil {
		return "", err
	}

	return user.ID.String(), nil
}

func (u *usecase) GetUserList(ctx context.Context) ([]models.User, error) {
	return u.repo.GetUserList(ctx)
}

func (u *usecase) AcceptNewTicket(ctx context.Context, newTicket models.Ticket) (string, error) {
	newTicket.ID = uuid.New()
	newTicket.Status = models.Accept

	err := u.repo.CreateTicket(ctx, newTicket)
	if err != nil {
		return "", err
	}

	data, err := json.Marshal(newTicket)
	if err != nil {
		return "", err
	}

	err = u.qCli.Publish(services.NewTicketQueueName, data)
	if err != nil {
		return "", err
	}

	return newTicket.ID.String(), nil
}

func (u *usecase) GetTicketList(ctx context.Context, userID string) ([]models.Ticket, error) {
	return u.repo.GetTicketList(ctx, userID)
}
