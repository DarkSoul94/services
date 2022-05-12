package rabbitconsumer

import (
	queueclient "github.com/DarkSoul94/services/pkg/QueueClient"
	"github.com/DarkSoul94/services/pkg/QueueClient/rabbitmq"
	"github.com/DarkSoul94/services/service3/app"
	apprabbit "github.com/DarkSoul94/services/service3/app/delivery/rabbit"
	notificator "github.com/DarkSoul94/services/service3/app/notificator/email"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

type App struct {
	rabbitConn *amqp.Connection
	rabbitChan *amqp.Channel
	qCli       queueclient.QueueClient
	notificator         app.Notificator
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

	n := notificator.NewEmailNotifycator()

	return &App{
		rabbitConn: con,
		rabbitChan: chn,
		qCli:       qCli,
		notificator: n,
	}
}

func (a *App) Run() {
	apprabbit.RegisterConsumers(a.qCli, a.notificator)
}

func (a *App) Stop() {
	defer a.rabbitChan.Close()
	defer a.rabbitConn.Close()
}
