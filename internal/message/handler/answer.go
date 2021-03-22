package handler

import (
	"errors"
	"fmt"
	errors2 "github.com/sepuka/focalism/errors"
	"github.com/sepuka/focalism/internal/domain"
	button2 "github.com/sepuka/focalism/internal/message/button"
	"github.com/sepuka/vkbotserver/api"
	"github.com/sepuka/vkbotserver/api/button"
	domain2 "github.com/sepuka/vkbotserver/domain"
)

type (
	Answer struct {
		taskRepository domain.TaskRepository
		api            *api.Api
	}
)

func NewAnswer(repo domain.TaskRepository, api *api.Api) *Answer {
	return &Answer{
		taskRepository: repo,
		api:            api,
	}
}

func (h *Answer) Handle(req *domain2.Request) error {
	var (
		err      error
		answer   = `Correct!`
		msg      = req.Object.Message.Text
		peerId   = int64(req.Object.Message.FromId)
		lastTask *domain.Task
		keyboard = button.Keyboard{
			OneTime: true,
		}
	)

	if lastTask, err = h.taskRepository.GetLast(peerId); err != nil {
		if errors.Is(err, errors2.NoError) {
			// чел не получал заданий до этого
		}
	}

	if lastTask.Vocabulary.Answer != msg {
		answer = `Wrong answer`
	}

	keyboard.Buttons = button2.NextWithReturn(fmt.Sprintf(`%d`, lastTask.Vocabulary.TopicId))

	return h.api.SendMessageWithButton(int(peerId), answer, keyboard)
}
