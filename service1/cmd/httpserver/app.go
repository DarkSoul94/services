package server

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/DarkSoul94/services/pkg/QueueClient/rabbitmq"
	"github.com/DarkSoul94/services/service1/app"
	apphttp "github.com/DarkSoul94/services/service1/app/delivery/http"
	appusecase "github.com/DarkSoul94/services/service1/app/usecase"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

type App struct {
	rabbitConn *amqp.Connection
	rabbitChan *amqp.Channel

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

	uc := appusecase.NewUsecase(publisher)

	return &App{
		rabbitConn: con,
		rabbitChan: chn,

		appUC: uc,
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
	defer a.rabbitConn.Close()
	defer a.rabbitChan.Close()

	return a.httpServer.Shutdown(ctx)
}
