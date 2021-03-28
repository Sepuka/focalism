package handler

import (
	button2 "github.com/sepuka/focalism/internal/message/button"
	"github.com/sepuka/vkbotserver/api"
	"github.com/sepuka/vkbotserver/api/button"
	"github.com/sepuka/vkbotserver/domain"
)

type (
	returnHandler struct {
		api *api.Api
	}
)

func NewReturnHandler(
	api *api.Api,
) *returnHandler {
	return &returnHandler{
		api: api,
	}
}

// Handles "return" command
// It should show 2 buttons in order to mode choosing: Irregular verbs OR Topic mode
func (h *returnHandler) Handle(req *domain.Request, payload *button.Payload) error {
	const (
		msg = `Пожалуйста, выберете упражнение`
	)

	var (
		peerId   = int(req.Object.Message.FromId)
		keyboard = button.Keyboard{
			OneTime: true,
			Buttons: button2.ModeChoose(),
		}
	)

	return h.api.SendMessageWithButton(peerId, msg, keyboard)
}
