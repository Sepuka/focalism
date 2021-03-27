package handler

import (
	"fmt"
	"github.com/sepuka/vkbotserver/api/button"
	"github.com/sepuka/vkbotserver/domain"
)

const (
	irregularTopicId = 1
)

type (
	irregularHandler struct {
		nextHandler *nextHandler
	}
)

func NewIrregularHandler(
	handler *nextHandler,
) *irregularHandler {
	return &irregularHandler{nextHandler: handler}
}

func (h *irregularHandler) Handle(req *domain.Request, payload *button.Payload) error {
	payload.Id = fmt.Sprintf(`%d`, irregularTopicId)

	return h.nextHandler.Handle(req, payload)
}
