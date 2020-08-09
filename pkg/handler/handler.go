package handler

import (
	"github.com/solympe/solympe-bot/pkg/models"
)

type service interface {
	Info(request models.Update)
}

// Handler ...
type Handler interface {
	Handle()
}

type handler struct {
	pull    <-chan models.Update
	service service
}

// Handle ...
func (h *handler) Handle() {
	var update models.Update
	for {
		select {
		case update = <-h.pull:
			h.handleMessage(update)
		}
	}
}

func (h *handler) handleMessage(update models.Update) {
	switch update.Message.Text {
	case info, infoChat:
		go h.service.Info(update)
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
