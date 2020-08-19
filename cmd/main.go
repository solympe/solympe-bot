package main

import (
	"context"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/solympe/solympe-bot/pkg/handler"
	blogger "github.com/solympe/solympe-bot/pkg/logger"
	"github.com/solympe/solympe-bot/pkg/models"
	receiver2 "github.com/solympe/solympe-bot/pkg/receiver"
	"github.com/solympe/solympe-bot/pkg/responder"
	"github.com/solympe/solympe-bot/pkg/service"
	"github.com/solympe/solympe-bot/pkg/updater"
)

const (
	configFile = "info.txt"
)

func main() {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller) // TODO level

	// TODO------------------------------------------- reading url and token from file
	content, err := ioutil.ReadFile(configFile)
	if err != nil {
		_ = level.Error(logger).Log("failed to read config file")
		return
	}
	lines := strings.Split(string(content), "\n")
	if len(lines) < 1 {
		_ = level.Error(logger).Log("failed to get data from config file")
		return
	}
	url := lines[0]

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	messagePull := make(chan models.Update, 100)

	botLogger := blogger.New(logger)

	botResponder := responder.NewResponder(url)
	botService := service.New(botResponder, botLogger)
	svcHandler := handler.NewHandler(messagePull, botService)
	messageUpdater := updater.New(url)
	receiver := receiver2.NewReceiver(messagePull, messageUpdater, botLogger)

	go receiver.WaitMessages()
	go svcHandler.Handle(ctx)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	defer func(sig os.Signal) {

		close(messagePull)
		_ = level.Info(logger).Log("bye")
	}(<-c)
}
