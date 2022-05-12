package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/DarkSoul94/services/pkg/config"
	"github.com/DarkSoul94/services/pkg/logger"
	httpserver "github.com/DarkSoul94/services/service1/cmd/httpserver"
	"github.com/spf13/viper"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Fatal(err)
	}
	logger.InitLogger()
	apphttp := httpserver.NewApp()
	apphttp.Run(viper.GetString("app.http_port"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	apphttp.Stop()
}
