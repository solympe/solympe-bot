package receiver

import (
	"github.com/solympe/solympe-bot/pkg/models"
)

type logger interface {
	Log(kind string, message string, err error)
}

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
	logger  logger
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
			r.logger.Log("receiver", "failed to get update", err)
		}
		for i := range updates.Result {
			r.pull <- updates.Result[i]
			r.logger.Log("receiver", updates.Result[i].Message.Text, err)
			offset = updates.Result[i].UpdateID + 1
		}
	}
}

// NewReceiver ...
func NewReceiver(
	pull chan models.Update,
	updater updater,
	logger logger,
) Receiver {
	return &receiver{
		pull:    pull,
		updater: updater,
		logger: logger,
	}
}
