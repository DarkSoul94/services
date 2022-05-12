package rabbitconsumer

import (
	queueclient "github.com/DarkSoul94/services/pkg/QueueClient"
	"github.com/DarkSoul94/services/pkg/QueueClient/rabbitmq"
	"github.com/DarkSoul94/services/service2/app"
	apprabit "github.com/DarkSoul94/services/service2/app/delivery/rabbit"
	appusecase "github.com/DarkSoul94/services/service2/app/usecase"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

type App struct {
	rabbitConn *amqp.Connection
	rabbitChan *amqp.Channel
	qCli       queueclient.QueueClient
	uc         app.Usecase
}

func NewApp() *App {
	con, chn, err := rabbitmq.Connect(
		viper.GetString("app.rabbit.login"),
		viper.GetString("app.rabbit.pass"),
		viper.GetString("app.rabbit.host"),
		viper.GetInt("app.rabbit.port"),
	)
	if err != nil {
		panic(err)
	}
	qCli := rabbitmq.NewRabbitClient(chn)

	uc := appusecase.NewUsecase(qCli)

	return &App{
		rabbitConn: con,
		rabbitChan: chn,
		qCli:       qCli,
		uc:         uc,
	}
}

func (a *App) Run() {
	apprabit.RegisterConsumers(a.qCli, a.uc)
}

func (a *App) Stop() {
	defer a.rabbitChan.Close()
	defer a.rabbitConn.Close()
}
