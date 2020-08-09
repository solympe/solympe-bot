package responder

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/solympe/solympe-bot/pkg/models"
)

// Responder ...
type Responder interface {
	SendResponse(response models.BotMessage) (err error)
}

type responder struct {
	botUrl string
}

// SendResponse ...
func (r *responder) SendResponse(response models.BotMessage) (err error) {
	buf, err := json.Marshal(response)
	if err != nil {
		return
	}
	_, err = http.Post(r.botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	return
}

func NewResponder(botUrl string) Responder {
	return &responder{
		botUrl: botUrl,
	}
}
