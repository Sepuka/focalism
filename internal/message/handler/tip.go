package handler

import (
	"fmt"
	"github.com/sepuka/focalism/errors"
	domain2 "github.com/sepuka/focalism/internal/domain"
	button2 "github.com/sepuka/focalism/internal/message/button"
	"github.com/sepuka/vkbotserver/api"
	"github.com/sepuka/vkbotserver/api/button"
	"github.com/sepuka/vkbotserver/domain"
	"strconv"
	"strings"
)

type (
	tipHandler struct {
		api  *api.Api
		repo domain2.TaskRepository
	}
)

func NewTipHandler(
	api *api.Api,
	taskRepository domain2.TaskRepository,
) *tipHandler {
	return &tipHandler{
		api:  api,
		repo: taskRepository,
	}
}

func (h *tipHandler) Handle(req *domain.Request, payload *button.Payload) error {
	var (
		taskId   int64
		peerId   = int(req.Object.Message.FromId)
		err      error
		task     domain2.Task
		keyboard = button.Keyboard{
			OneTime: true,
		}
	)

	if taskId, err = strconv.ParseInt(payload.Id, 10, 64); err != nil {
		return errors.NewInvalidJsonError(`could not parse task ID`, err)
	}

	if task, err = h.repo.GetById(taskId); err != nil {
		return errors.NewDatabaseError(fmt.Sprintf(`task "%d" has not found`, taskId), err)
	}

	keyboard.Buttons = button2.SurrenderAndReturn(task.GetId())

	return h.api.SendMessageWithButton(peerId, h.tip(task), keyboard)
}

func (h *tipHandler) tip(task domain2.Task) string {
	var (
		length = len(task.Vocabulary.Answer)
		answer string
	)

	if length > 2 {
		stars := strings.Repeat(`*`, length-2)
		answer = fmt.Sprintf(`"%c%s%c"`, task.Vocabulary.Answer[0], stars, task.Vocabulary.Answer[length-1])
	} else {
		answer = fmt.Sprintf(`"%c**"`, task.Vocabulary.Answer[0])
	}

	if len(task.Vocabulary.Example) > 0 {
		answer = fmt.Sprintf("%s\n\n%s", answer, task.Vocabulary.Example)
	}

	return fmt.Sprintf(`%s`, answer)
}
