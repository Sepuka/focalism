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
	"go.uber.org/zap"
)

type (
	Answer struct {
		taskRepository domain.TaskRepository
		api            *api.Api
		log            *zap.SugaredLogger
	}
)

func NewAnswer(repo domain.TaskRepository, api *api.Api, log *zap.SugaredLogger) *Answer {
	return &Answer{
		taskRepository: repo,
		api:            api,
		log:            log,
	}
}

func (h *Answer) Handle(req *domain2.Request) error {
	var (
		err      error
		answer   = `Correct!`
		msg      = req.Object.Message.Text
		peerId   = int64(req.Object.Message.FromId)
		lastTask domain.Task
		keyboard = button.Keyboard{
			OneTime: true,
		}
	)

	if lastTask, err = h.taskRepository.GetLast(); err != nil {
		if errors.Is(err, errors2.NoError) {
			keyboard.Buttons = button2.ModeChoose()
			return h.api.SendMessageWithButton(int(peerId), `выберите режим`, keyboard)
		}
	}

	if lastTask.Vocabulary.Answer != msg {
		answer = `Wrong answer`
	} else {
		lastTask.IsCorrect = true
	}

	if err = h.taskRepository.Answer(lastTask); err != nil {
		h.
			log.
			With(
				zap.Int64(`peerId`, peerId),
				zap.Int64(`taskId`, lastTask.Id),
				zap.Error(err),
			).
			Error(`answer time update failed`)
	}

	keyboard.Buttons = button2.NextWithReturn(fmt.Sprintf(`%d`, lastTask.Vocabulary.TopicId))

	return h.api.SendMessageWithButton(int(peerId), answer, keyboard)
}
