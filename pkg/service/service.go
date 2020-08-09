package service

import (
	"fmt"

	"github.com/solympe/solympe-bot/pkg/models"
)

type logger interface {
	Log(kind string, message string, err error)
}

type responder interface {
	SendResponse(response models.BotMessage) (err error)
}

// Service ...
type Service interface {
	Info(request models.Update)
}

type service struct {
	responder responder
	logger    logger
}

// Do ...
func (s *service) Info(request models.Update) {
	var (
		botMessage     = models.BotMessage{}
		responseString string
		err            error
	)

	responseString = fmt.Sprintf(
		`Hello, %s!
I was created by @solympe as a test project.
GitHub: https://github.com/solympe/solympe-bot`,
		request.Message.User.FirstName)

	botMessage.ID = request.Message.Chat.ID
	botMessage.Text = responseString

	err = s.responder.SendResponse(botMessage)
	if err != nil {
		s.logger.Log("svc", "info", err)
	}
	return
}

// New ...
func New(
	responder responder,
	logger logger,
) Service {
	return &service{
		responder: responder,
		logger:    logger,
	}
}
