package service

import (
	"fmt"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type logger interface {
	Log(kind string, message string, err error)
}

// Service ...
type Service interface {
	Info(request tgbotapi.Update)
	Roll(request tgbotapi.Update)
	Join(request tgbotapi.Update)
}

type service struct {
	bot    *tgbotapi.BotAPI
	logger logger
}

// key == chat id, value == map of users
var rollMap = map[int64]map[tgbotapi.User]string{}

// Do ...
func (s *service) Info(request tgbotapi.Update) {
	var messageText string

	messageText = fmt.Sprintf(
		`Hello, %s!
I was created by @solympe as a test project.
GitHub: https://github.com/solympe/solympe-bot`,
		request.Message.From.FirstName)
	msg := tgbotapi.NewMessage(request.Message.Chat.ID, messageText)

	s.sendMessage(msg)
	return
}

// Roll ...
func (s *service) Roll(request tgbotapi.Update) {
	msg := tgbotapi.NewMessage(request.Message.Chat.ID, "OK! Lets start! I need more than two participants to start. Press 'Join to roll' on keyboard")

	// check if this game is already started
	if _, ok := rollMap[request.Message.Chat.ID]; ok {
		msg.Text = "This game has already started!"
		s.sendMessage(msg)
		return
	}

	numericKeyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Join to roll"),
		),
	)
	msg.ReplyMarkup = numericKeyboard
	s.sendMessage(msg)

	// init chat map
	users := make(map[tgbotapi.User]string)
	rollMap[request.Message.Chat.ID] = users
	defer delete(rollMap, request.Message.Chat.ID)

	msg.Text = "I will wait 10 seconds and after that the game will start"
	s.sendMessage(msg)

	time.Sleep(10 * time.Second)
	msg.ReplyMarkup = tgbotapi.NewHideKeyboard(true)
	msg.Text = "OK! And now I will randomly determine the winner!"
	s.sendMessage(msg)

	if len(rollMap[request.Message.Chat.ID]) < 2 {
		msg.Text = "Oh.... not enough participants to start the game"
		s.sendMessage(msg)
		return
	}

	// random winner
	for _, v := range rollMap[request.Message.Chat.ID] {
		msg.Text = fmt.Sprintf("Winner is: @%v!", v)
		s.sendMessage(msg)
		return
	}
}

// Join ...
func (s *service) Join(request tgbotapi.Update) {
	// if game is not exists
	if _, ok := rollMap[request.Message.Chat.ID]; !ok {
		return
	}
	// if user exists
	if _, ok := rollMap[request.Message.Chat.ID][*request.Message.From]; ok {
		return
	}

	rollMap[request.Message.Chat.ID][*request.Message.From] = request.Message.From.UserName
	messageText := fmt.Sprintf("%v joins the game!", request.Message.From.UserName)
	msg := tgbotapi.NewMessage(request.Message.Chat.ID, messageText)
	s.sendMessage(msg)
}

// sendMessage - sends response
func (s *service) sendMessage(msg tgbotapi.Chattable) {
	_, err := s.bot.Send(msg)
	if err != nil {
		s.logger.Log("service", "failed to send message", err)
	}
}

// New ...
func New(
	bot *tgbotapi.BotAPI,
	logger logger,
) Service {
	return &service{
		bot:    bot,
		logger: logger,
	}
}
