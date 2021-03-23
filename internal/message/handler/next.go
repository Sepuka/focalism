package handler

import (
	"github.com/sepuka/focalism/errors"
	domain2 "github.com/sepuka/focalism/internal/domain"
	button2 "github.com/sepuka/focalism/internal/message/button"
	"github.com/sepuka/vkbotserver/api"
	"github.com/sepuka/vkbotserver/api/button"
	"github.com/sepuka/vkbotserver/domain"
	"strconv"
)

type (
	nextHandler struct {
		api                  *api.Api
		vocabularyRepository domain2.VocabularyRepository
		taskRepository       domain2.TaskRepository
	}
)

func NewNextHandler(
	api *api.Api,
	vocabularyRepo domain2.VocabularyRepository,
	taskRepo domain2.TaskRepository,
) *nextHandler {
	return &nextHandler{
		api:                  api,
		vocabularyRepository: vocabularyRepo,
		taskRepository:       taskRepo,
	}
}

func (h *nextHandler) Handle(req *domain.Request, payload *button.Payload) error {
	var (
		err        error
		vocabulary *domain2.Vocabulary
		task       *domain2.Task
		peerId     = int(req.Object.Message.FromId)
		keyboard   = button.Keyboard{
			OneTime: true,
		}
		topicId int64
	)

	if topicId, err = strconv.ParseInt(payload.Id, 10, 64); err != nil {
		return errors.NewInvalidJsonError(`could not parse topic ID`, err)
	}

	if vocabulary, err = h.vocabularyRepository.FindActual(topicId); err != nil {
		return errors.NewDatabaseError(`could not fetch next word`, err)
	}

	defer h.vocabularyRepository.IncrViews(vocabulary)

	if task, err = h.taskRepository.Create(vocabulary, int64(peerId)); err != nil {
		return errors.NewDatabaseError(`could not create new task`, err)
	}

	keyboard.Buttons = button2.Answer(button.Payload{
		Command: button2.ButtonIdSurrender,
		Id:      task.GetId(),
	})

	return h.api.SendMessageWithAttachmentAndButton(peerId, vocabulary.Question, vocabulary.Attachment, keyboard)
}