package handler

import (
	button2 "github.com/sepuka/focalism/internal/message/button"
	"github.com/sepuka/vkbotserver/api"
	"github.com/sepuka/vkbotserver/api/button"
	"github.com/sepuka/vkbotserver/domain"
)

const (
	rules = `Выберете тип заданий`
)

type (
	startHandler struct {
		api *api.Api
	}
)

func NewStartHandler(api *api.Api) *startHandler {
	return &startHandler{api: api}
}

func (h *startHandler) Handle(req *domain.Request, payload *button.Payload) error {
	var (
		peerId   = int(req.Object.Message.FromId)
		keyboard = button.Keyboard{
			OneTime: true,
			Buttons: button2.ModeChoose(),
		}
	)

	return h.api.SendMessageWithButton(peerId, rules, keyboard)
}
