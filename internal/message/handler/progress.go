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
	progressHandler struct {
		api            *api.Api
		taskRepository domain2.TaskProgressRepository
		vocabularyRepo domain2.VocabularyRepository
	}
)

func NewProgressHandler(
	api *api.Api,
	taskRepo domain2.TaskProgressRepository,
	vocabularyRepo domain2.VocabularyRepository,
) *progressHandler {
	return &progressHandler{
		api:            api,
		taskRepository: taskRepo,
		vocabularyRepo: vocabularyRepo,
	}
}

func (h *progressHandler) Handle(req *domain.Request, payload *button.Payload) error {
	var (
		err                      error
		topicId                  int64
		peerId                   = int64(req.Object.Message.FromId)
		success, attempts, total int
		msgTmpl                  = `сегодня вы верно назвали %d слов из %d попыток`
		keyboard                 = button.Keyboard{
			OneTime: true,
			Buttons: button2.NextWithReturn(fmt.Sprintf(`%d`, topicId)),
		}
	)

	if topicId, err = strconv.ParseInt(payload.Id, 10, 64); err != nil {
		return errors.NewInvalidJsonError(`could not parse topic ID`, err)
	}

	if success, attempts, err = h.
		taskRepository.
		GetProgress(topicId, peerId); err != nil {
		return errors.NewDatabaseError(`could not fetch current user's progress`, err)
	}

	if total, err = h.vocabularyRepo.GetTotal(topicId); err != nil {
		return errors.NewDatabaseError(`could not fetch total vocabulary items`, err)
	}

	if attempts == total {
		keyboard.Buttons = button2.Return()
	}

	return h.api.SendMessageWithButton(int(peerId), fmt.Sprintf(msgTmpl, success, attempts), keyboard)
}
