package handler

import (
	domain2 "github.com/sepuka/focalism/internal/domain"
	button2 "github.com/sepuka/focalism/internal/message/button"
	"github.com/sepuka/vkbotserver/api"
	"github.com/sepuka/vkbotserver/api/button"
	"github.com/sepuka/vkbotserver/domain"
	"go.uber.org/zap"
)

type (
	returnHandler struct {
		api      *api.Api
		taskRepo domain2.TaskRepository
		log      *zap.SugaredLogger
	}
)

func NewReturnHandler(
	api *api.Api,
	taskRepo domain2.TaskRepository,
	log *zap.SugaredLogger,
) *returnHandler {
	return &returnHandler{
		api:      api,
		taskRepo: taskRepo,
		log:      log,
	}
}

// Handles "return" command
// It should show 2 buttons in order to mode choosing: Irregular verbs OR Topic mode
func (h *returnHandler) Handle(req *domain.Request, payload *button.Payload) error {
	const (
		msg = `Пожалуйста, выберите упражнение`
	)

	var (
		peerId   = int(req.Object.Message.FromId)
		keyboard = button.Keyboard{
			OneTime: true,
			Buttons: button2.ModeChoose(),
		}
		err error
	)

	if err = h.taskRepo.DeleteLast(int64(peerId)); err != nil {
		h.
			log.
			With(
				zap.Int(`peer_id`, peerId),
			).
			Error(`error while deleting skipped tasks`)
	}

	return h.api.SendMessageWithButton(peerId, msg, keyboard)
}
