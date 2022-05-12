package rabbit

import (
	"github.com/DarkSoul94/services"
	queueclient "github.com/DarkSoul94/services/pkg/QueueClient"
	"github.com/DarkSoul94/services/service2/app"
)

func RegisterConsumers(qCli queueclient.QueueClient, uc app.Usecase) {
	h := NewHandler(uc)
	go func() {
		msgChan, err := qCli.Consume(services.NewTicketQueueName)
		if err != nil {
			panic(err)
		}

		for msg := range msgChan {
			h.AcceptTicket(msg)
		}
	}()
}
