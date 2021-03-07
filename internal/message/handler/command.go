package handler

import (
	"github.com/sepuka/vkbotserver/api/button"
	"github.com/sepuka/vkbotserver/domain"
)

type MessageHandler interface {
	Handle(*domain.Request, *button.Payload) error
}
