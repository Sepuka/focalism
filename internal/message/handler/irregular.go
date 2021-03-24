package handler

import (
	"github.com/sepuka/focalism/errors"
	domain2 "github.com/sepuka/focalism/internal/domain"
	button2 "github.com/sepuka/focalism/internal/message/button"
	"github.com/sepuka/vkbotserver/api"
	"github.com/sepuka/vkbotserver/api/button"
	"github.com/sepuka/vkbotserver/domain"
)

const (
	irregularTopicId = 1
)

type (
	irregularHandler struct {
		api                  *api.Api
		vocabularyRepository domain2.VocabularyRepository
		taskRepository       domain2.TaskRepository
	}
)

func NewIrregularHandler(
	api *api.Api,
	vocabularyRepo domain2.VocabularyRepository,
	taskRepo domain2.TaskRepository,
) *irregularHandler {
	return &irregularHandler{
		api:                  api,
		vocabularyRepository: vocabularyRepo,
		taskRepository:       taskRepo,
	}
}

func (h *irregularHandler) Handle(req *domain.Request, payload *button.Payload) error {
	var (
		err        error
		vocabulary domain2.Vocabulary
		task       *domain2.Task
		peerId     = int(req.Object.Message.FromId)
		keyboard   = button.Keyboard{
			OneTime: true,
		}
	)

	if vocabulary, err = h.vocabularyRepository.FindActual(irregularTopicId); err != nil {
		return errors.NewDatabaseError(`could not fetch next word`, err)
	}

	defer h.vocabularyRepository.IncrViews(vocabulary)

	if task, err = h.taskRepository.Create(vocabulary, int64(peerId)); err != nil {
		return errors.NewDatabaseError(`could not create new task`, err)
	}

	keyboard.Buttons = button2.Surrender(task.GetId())

	return h.api.SendMessageWithAttachmentAndButton(peerId, vocabulary.Question, vocabulary.Attachment, keyboard)
}
