package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/solympe/solympe-bot/pkg/handler"
	"github.com/solympe/solympe-bot/pkg/models"
	"github.com/solympe/solympe-bot/pkg/responder"
	"github.com/solympe/solympe-bot/pkg/service"
	"github.com/solympe/solympe-bot/pkg/updater"
)

const (
	configFile = "info.txt"
)

func main() {
	// TODO-------------------------------------------
	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Println("failed to read file", err)
		return
	}
	lines := strings.Split(string(content), "\n")
	if len(lines) < 2 {
		log.Println("data error", err)
		return
	}
	url := lines[0] + lines[1]
	// TODO-------------------------------------------

	messagePull := make(chan models.Update, 100)

	botResponder := responder.NewResponder(url)
	botService := service.New(botResponder)
	svcHandler := handler.NewHandler(messagePull, botService)
	messageUpdater := updater.New(url)
	receiver := handler.NewReceiver(messagePull, messageUpdater)

	go receiver.WaitMessages()
	go svcHandler.Handle()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	defer func(sig os.Signal) {
		close(messagePull)
		log.Println("bye")
	}(<-c)

}
