package rabbitconsumer

import (
	"context"
	"fmt"

	queueclient "github.com/DarkSoul94/services/pkg/QueueClient"
	"github.com/DarkSoul94/services/pkg/QueueClient/rabbitmq"
	"github.com/DarkSoul94/services/service2/app"
	apprabit "github.com/DarkSoul94/services/service2/app/delivery/rabbit"
	apprepo "github.com/DarkSoul94/services/service2/app/repo/mongo"
	appusecase "github.com/DarkSoul94/services/service2/app/usecase"

	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	rabbitConn *amqp.Connection
	rabbitChan *amqp.Channel
	qCli       queueclient.QueueClient
	dbClient   *mongo.Client
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

	db := mongoDbInit(
		viper.GetString("app.db.host"),
		viper.GetInt("app.db.port"),
	)

	repo := apprepo.NewMongoRepo(db.Database(viper.GetString("app.db.name")))

	uc := appusecase.NewUsecase(qCli, repo)

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

func mongoDbInit(host string, port int) *mongo.Client {
	ctx := context.Background()

	clientOptions := options.Client().ApplyURI(
		fmt.Sprintf(
			"mongodb://%s:%d/",
			host,
			port,
		),
	)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	return client
}
