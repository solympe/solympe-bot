package handler

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type logger interface {
	Log(kind string, message string, err error)
}

type service interface {
	Info(request tgbotapi.Update)
	Roll(request tgbotapi.Update)
	Join(request tgbotapi.Update)
}

// Handler ...
type Handler interface {
	Handle(ctx context.Context)
}

type handler struct {
	bot     *tgbotapi.BotAPI
	service service
	logger  logger
}

// Handle ...
func (h *handler) Handle(ctx context.Context) {
	var update tgbotapi.Update

	updConfig := tgbotapi.NewUpdate(0)
	updChan, err := h.bot.GetUpdatesChan(updConfig)
	if err != nil {
		h.logger.Log("handler", "failed to get upd chan", err)
		return
	}

	for {
		select {
		case <-ctx.Done():
			updChan.Clear()
			return
		case update = <-updChan:
			go h.handleMessage(update)
		}
	}
}

func (h *handler) handleMessage(update tgbotapi.Update) {
	switch update.Message.Text {
	case info, infoChat:
		h.service.Info(update)
	case roll, rollChat:
		h.service.Roll(update)
	case join:
		h.service.Join(update)
	}
}

// New ...
func New(
	bot *tgbotapi.BotAPI,
	service service,
	logger logger,
) Handler {
	return &handler{
		bot:     bot,
		service: service,
		logger:  logger,
	}
}
