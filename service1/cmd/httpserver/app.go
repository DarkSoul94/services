package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/DarkSoul94/services/pkg/QueueClient/rabbitmq"
	"github.com/DarkSoul94/services/service1/app"
	apphttp "github.com/DarkSoul94/services/service1/app/delivery/http"
	apprepo "github.com/DarkSoul94/services/service1/app/repo/mongo"
	appusecase "github.com/DarkSoul94/services/service1/app/usecase"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	rabbitConn *amqp.Connection
	rabbitChan *amqp.Channel
	dbClient   *mongo.Client
	appUC      app.Usecase
	httpServer *http.Server
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

	publisher := rabbitmq.NewRabbitClient(chn)

	db := mongoDbInit(
		viper.GetString("app.db.host"),
		viper.GetInt("app.db.port"),
	)

	repo := apprepo.NewMongoRepo(db.Database(viper.GetString("app.db.name")))

	uc := appusecase.NewUsecase(publisher, repo)

	return &App{
		rabbitConn: con,
		rabbitChan: chn,
		dbClient:   db,
		appUC:      uc,
	}
}

func (a *App) Run(port string) {

	router := gin.New()
	if viper.GetBool("app.release") {
		gin.SetMode(gin.ReleaseMode)
	} else {
		router.Use(gin.Logger())
	}

	apiRouter := router.Group("/api")
	apphttp.RegisterHTTPEndpoints(apiRouter, a.appUC)

	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	var l net.Listener
	var err error
	l, err = net.Listen("tcp", a.httpServer.Addr)
	if err != nil {
		panic(err)
	}

	if err := a.httpServer.Serve(l); err != nil {
		log.Fatalf("Failed to listen and serve: %+v", err)
	}
}

func (a *App) Stop() error {
	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()
	a.rabbitChan.Close()
	a.rabbitConn.Close()
	a.dbClient.Disconnect(ctx)

	return a.httpServer.Shutdown(ctx)
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
