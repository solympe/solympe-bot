package handler

import (
	"log"

	"github.com/solympe/solympe-bot/pkg/models"
)

type updater interface {
	GetUpdates(offset int) (updates models.UpdateResponse, err error)
}

// Receiver ...
type Receiver interface {
	WaitMessages()
}

type receiver struct {
	pull    chan models.Update
	updater updater
}

// WaitMessages ...
func (r *receiver) WaitMessages() {
	var (
		offset  int
		updates models.UpdateResponse
		err     error
	)

	for {
		updates, err = r.updater.GetUpdates(offset)
		if err != nil {
			log.Println(err)
		}
		for i := range updates.Result {
			r.pull <- updates.Result[i]
			offset = updates.Result[i].UpdateID + 1
		}
	}
}

// NewReceiver ...
func NewReceiver(
	pull chan models.Update,
	updater updater,
) Receiver {
	return &receiver{
		pull:    pull,
		updater: updater,
	}
}
