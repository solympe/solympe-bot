package handler

import (
	"context"

	"github.com/solympe/solympe-bot/pkg/models"
)

type service interface {
	Info(request models.Update)
}

// Handler ...
type Handler interface {
	Handle(ctx context.Context)
}

type handler struct {
	pull    <-chan models.Update
	service service
}

// Handle ...
func (h *handler) Handle(ctx context.Context) {
	var update models.Update
	for {
		select {
		case update = <-h.pull:
			go h.handleMessage(update)
		case <-ctx.Done():
			for msg := range h.pull {
				go h.handleMessage(msg)
			}
			return

		}
	}
}

func (h *handler) handleMessage(update models.Update) {
	switch update.Message.Text {
	case info, infoChat:
		h.service.Info(update)
	}
}

func NewHandler(
	pull <-chan models.Update,
	service service,
) Handler {
	return &handler{
		pull:    pull,
		service: service,
	}
}
