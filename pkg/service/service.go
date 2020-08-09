package service

import (
	"fmt"

	"github.com/solympe/solympe-bot/pkg/models"
)

type responder interface {
	SendResponse(response models.BotMessage) (err error)
}

// Service ...
type Service interface {
	Info(request models.Update) (err error)
}

type service struct {
	responder responder
}

// Do ...
func (s *service) Info(request models.Update) (err error) {
	var (
		botMessage     = models.BotMessage{}
		responseString string
	)

	responseString = fmt.Sprintf(
		`Hello, %s! My name is SolympeBot!

	I was created by @solympe as a test project.
	GitHub: https://github.com/solympe/solympe-bot`,
		request.Message.User.FirstName)

	botMessage.ID = request.Message.Chat.ID
	botMessage.Text = responseString

	err = s.responder.SendResponse(botMessage)
	return
}

// New ...
func New(responder responder) Service {
	return &service{
		responder: responder,
	}
}
