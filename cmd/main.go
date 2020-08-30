package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/kelseyhightower/envconfig"

	"github.com/solympe/solympe-bot/pkg/handler"
	blogger "github.com/solympe/solympe-bot/pkg/logger"
	"github.com/solympe/solympe-bot/pkg/service"
)

type configuration struct {
	BotURL string `envconfig:"BOT_URL" required:"true"`
	Debug  bool   `envconfig:"DEBUG" required:"true"`
}

func main() {
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller) // TODO level

	var cfg configuration
	if err := envconfig.Process("", &cfg); err != nil {
		_ = level.Error(logger).Log("msg", "failed to load configuration", "err", err)
		os.Exit(1)
	}
	if !cfg.Debug {
		logger = level.NewFilter(logger, level.AllowInfo())
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	bot, err := tgbotapi.NewBotAPI(cfg.BotURL)
	if err != nil {
		_ = level.Error(logger).Log("msg", "failed to init bot token", "err", err)
		os.Exit(1)
	}
	bot.Debug = cfg.Debug
	botLogger := blogger.New(logger)
	botService := service.New(bot, botLogger)
	svcHandler := handler.New(bot, botService, botLogger)

	go svcHandler.Handle(ctx)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	defer func(sig os.Signal) {
		_ = level.Info(logger).Log("main", "bye")
	}(<-c)
}
