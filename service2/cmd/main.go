package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/DarkSoul94/services/pkg/config"
	"github.com/DarkSoul94/services/service2/cmd/rabbitconsumer"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Fatal(err)
	}

	apprabbit := rabbitconsumer.NewApp()
	apprabbit.Run()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	apprabbit.Stop()
}
