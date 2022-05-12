package rabbit

import (
	"github.com/DarkSoul94/services"
	queueclient "github.com/DarkSoul94/services/pkg/QueueClient"
	"github.com/DarkSoul94/services/service3/app"
)

func RegisterConsumers(qCli queueclient.QueueClient, notificator app.Notificator) {
	h := NewHandler(notificator)
	go func() {
		msgChan, err := qCli.Consume(services.NotifyQueueName)
		if err != nil {
			panic(err)
		}

		for msg := range msgChan {
			h.AcceptNotificationMsg(msg)
		}
	}()
}
