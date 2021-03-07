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
)

type (
	surrenderHandler struct {
		api            *api.Api
		taskRepository domain2.TaskRepository
	}
)

func NewSurrenderHandler(
	api *api.Api,
	taskRepo domain2.TaskRepository,
) *surrenderHandler {
	return &surrenderHandler{
		api:            api,
		taskRepository: taskRepo,
	}
}

func (h *surrenderHandler) Handle(req *domain.Request, payload *button.Payload) error {
	var (
		err      error
		task     *domain2.Task
		taskId   int64
		peerId   = int(req.Object.Message.FromId)
		keyboard = button.Keyboard{
			OneTime: true,
			Buttons: button2.Next(button.Payload{
				Command: button2.ButtonIdNext,
			}),
		}
		msg string
	)

	if taskId, err = strconv.ParseInt(payload.Id, 10, 64); err != nil {
		return errors.NewInvalidJsonError(`could not parse task ID`, err)
	}

	if task, err = h.taskRepository.GetById(taskId); err != nil {
		return errors.NewDatabaseError(fmt.Sprintf(`task "%d" has not found`, taskId), err)
	}

	msg = fmt.Sprintf(`The correct answer is "%s"`, task.Vocabulary.Answer)

	return h.api.SendMessageWithButton(peerId, msg, keyboard)
}
