package updater

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/solympe/solympe-bot/pkg/models"
)

// Updater ...
type Updater interface {
	GetUpdates(offset int) (updates models.UpdateResponse, err error)
}

type updater struct {
	botUrl string
}

const (
	getUpdate = "/getUpdates"
)

// GetUpdates ...
func (u *updater) GetUpdates(offset int) (updates models.UpdateResponse, err error) {
	resp, err := http.Get(u.botUrl + getUpdate + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &updates)
	if err != nil {
		return
	}

	return
}

func New(botUrl string) Updater {
	return &updater{
		botUrl: botUrl,
	}
}
