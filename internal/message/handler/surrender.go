package handler

import (
	"fmt"
	"github.com/sepuka/focalism/errors"
	domain2 "github.com/sepuka/focalism/internal/domain"
	button2 "github.com/sepuka/focalism/internal/message/button"
	"github.com/sepuka/vkbotserver/api"
	"github.com/sepuka/vkbotserver/api/button"
	"github.com/sepuka/vkbotserver/domain"
	"go.uber.org/zap"
	"strconv"
)

type (
	surrenderHandler struct {
		api            *api.Api
		taskRepository domain2.TaskRepository
		log            *zap.SugaredLogger
	}
)

func NewSurrenderHandler(
	api *api.Api,
	taskRepo domain2.TaskRepository,
	log *zap.SugaredLogger,
) *surrenderHandler {
	return &surrenderHandler{
		api:            api,
		taskRepository: taskRepo,
		log:            log,
	}
}

func (h *surrenderHandler) Handle(req *domain.Request, payload *button.Payload) error {
	var (
		err      error
		task     domain2.Task
		taskId   int64
		peerId   = int(req.Object.Message.FromId)
		keyboard = button.Keyboard{
			OneTime: true,
		}
		msg string
	)

	if taskId, err = strconv.ParseInt(payload.Id, 10, 64); err != nil {
		return errors.NewInvalidJsonError(`could not parse task ID`, err)
	}

	if task, err = h.taskRepository.GetById(taskId); err != nil {
		return errors.NewDatabaseError(fmt.Sprintf(`task "%d" has not found`, taskId), err)
	}

	task.IsCorrect = false
	if err = h.taskRepository.Answer(task); err != nil {
		h.
			log.
			With(
				zap.Int(`peerId`, peerId),
				zap.Int64(`taskId`, task.Id),
				zap.Error(err),
			).
			Error(`answer time update failed`)
	}

	msg = fmt.Sprintf(`The correct answer is "%s"`, task.Vocabulary.Answer)

	keyboard.Buttons = button2.NextWithReturn(fmt.Sprintf(`%d`, task.Vocabulary.TopicId))

	return h.api.SendMessageWithButton(peerId, msg, keyboard)
}
