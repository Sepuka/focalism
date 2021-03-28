package handler

import (
	"fmt"
	"github.com/sepuka/focalism/errors"
	domain2 "github.com/sepuka/focalism/internal/domain"
	button2 "github.com/sepuka/focalism/internal/message/button"
	"github.com/sepuka/vkbotserver/api"
	"github.com/sepuka/vkbotserver/api/button"
	"github.com/sepuka/vkbotserver/domain"
)

type (
	topicHandler struct {
		api       *api.Api
		topicRepo domain2.TopicRepository
	}
)

func NewTopicHandler(
	api *api.Api,
	topicRepo domain2.TopicRepository,
) *topicHandler {
	return &topicHandler{
		api:       api,
		topicRepo: topicRepo,
	}
}

// Handles "topics" command
// It should show all topics buttons
func (h *topicHandler) Handle(req *domain.Request, payload *button.Payload) error {
	const (
		msg = `Выберите тему`
	)

	var (
		err      error
		peerId   = int(req.Object.Message.FromId)
		keyboard = button.Keyboard{
			OneTime: true,
		}
		topics      []domain2.Topic
		topic       domain2.Topic
		buttons     = button2.Return()
		topicButton []button.Button
	)

	if topics, err = h.topicRepo.GetList(domain2.TopicMode); err != nil {
		return errors.NewDatabaseError(`could not fetch any topic`, err)
	}

	for _, topic = range topics {
		topicButton = []button.Button{
			{
				Color: button.PrimaryColor,
				Action: button.Action{
					Type:  button2.TextButtonType,
					Label: button.Text(topic.Title),
					Payload: button.Payload{
						Command: button2.NextIdButton,
						Id:      fmt.Sprintf(`%d`, topic.TopicId),
					}.String(),
				},
			},
		}
		buttons = append(buttons, topicButton)
	}

	keyboard.Buttons = buttons

	return h.api.SendMessageWithButton(peerId, msg, keyboard)
}
