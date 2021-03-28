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
	progressHandler struct {
		api            *api.Api
		taskRepository domain2.TaskProgressRepository
		log            *zap.SugaredLogger
	}
)

func NewProgressHandler(
	api *api.Api,
	taskRepo domain2.TaskProgressRepository,
	log *zap.SugaredLogger,
) *progressHandler {
	return &progressHandler{
		api:            api,
		taskRepository: taskRepo,
		log:            log,
	}
}

func (h *progressHandler) Handle(req *domain.Request, payload *button.Payload) error {
	var (
		err            error
		topicId        int64
		peerId         = int64(req.Object.Message.FromId)
		success, total int
		msgTmpl        = `сегодня вы верно назвали %d слов из %d`
		keyboard       = button.Keyboard{
			OneTime: true,
			Buttons: button2.NextWithReturn(fmt.Sprintf(`%d`, topicId)),
		}
	)

	if topicId, err = strconv.ParseInt(payload.Id, 10, 64); err != nil {
		return errors.NewInvalidJsonError(`could not parse topic ID`, err)
	}

	success, total, err = h.
		taskRepository.
		GetProgress(topicId, peerId)

	return h.api.SendMessageWithButton(int(peerId), fmt.Sprintf(msgTmpl, success, total), keyboard)
}
