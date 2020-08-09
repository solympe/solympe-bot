package handler

import (
	"log"

	"github.com/solympe/solympe-bot/pkg/models"
)

type service interface {
	Info(request models.Update) (err error)
}

// Handler ...
type Handler interface {
	Handle() ()
}

type handler struct {
	pull    <-chan models.Update
	service service
}

// messages
const (
	info = "/info@SolympeBot"
)

// Handle ...
func (h *handler) Handle() () {
	var update models.Update
	for {
		select {
		case update = <-h.pull:
			log.Println("recieved a message:", update)
			h.handleMessage(update)
		}
	}
}

func (h *handler) handleMessage(update models.Update) {
	var err error
	switch update.Message.Text {
	case info:
		err = h.service.Info(update)
		if err != nil {
			log.Println(err)
		}
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
