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
	nextHandler struct {
		api                  *api.Api
		vocabularyRepository domain2.VocabularyRepository
		taskRepository       domain2.TaskRepository
		log                  *zap.SugaredLogger
	}
)

func NewNextHandler(
	api *api.Api,
	vocabularyRepo domain2.VocabularyRepository,
	taskRepo domain2.TaskRepository,
	log *zap.SugaredLogger,
) *nextHandler {
	return &nextHandler{
		api:                  api,
		vocabularyRepository: vocabularyRepo,
		taskRepository:       taskRepo,
		log:                  log,
	}
}

func (h *nextHandler) Handle(req *domain.Request, payload *button.Payload) error {
	var (
		err        error
		vocabulary domain2.Vocabulary
		task       *domain2.Task
		peerId     = int(req.Object.Message.FromId)
		keyboard   = button.Keyboard{
			OneTime: true,
		}
		topicId                           int64
		tasksPerDay, totalVocabularyItems int
		question                          string
	)

	if topicId, err = strconv.ParseInt(payload.Id, 10, 64); err != nil {
		return errors.NewInvalidJsonError(`could not parse topic ID`, err)
	}

	if tasksPerDay, err = h.taskRepository.GetTodayTasks(topicId, peerId); err != nil {
		return errors.NewDatabaseError(`could not calculate today tasks`, err)
	} else {
		if totalVocabularyItems, err = h.vocabularyRepository.GetTotal(topicId); err != nil {
			return errors.NewDatabaseError(`could not calculate total vocabulary items`, err)
		}
	}

	if tasksPerDay+1 > totalVocabularyItems {
		keyboard.Buttons = button2.ReturnWithProgress(fmt.Sprintf(`%d`, topicId))

		return h.api.SendMessageWithButton(peerId, fmt.Sprintf(`сегодня вы повторили все слова этой темы (%d), приходите к нам завтра`, tasksPerDay), keyboard)
	}

	if vocabulary, err = h.vocabularyRepository.FindActual(topicId, int64(peerId)); err != nil {
		h.log.Debugf(`could not fetch next word: %s`, err)

		keyboard.Buttons = button2.ReturnWithProgress(fmt.Sprintf(`%d`, topicId))

		return h.api.SendMessageWithButton(peerId, `Извините, но слова этого вида закончились. Приходите завтра`, keyboard)
	}

	defer h.vocabularyRepository.IncrViews(vocabulary)

	if task, err = h.taskRepository.Create(vocabulary, int64(peerId)); err != nil {
		return errors.NewDatabaseError(`could not create new task`, err)
	}

	keyboard.Buttons = button2.Surrender(task.GetId())
	question = fmt.Sprintf(`(%d / %d). "%s"`, tasksPerDay+1, totalVocabularyItems, vocabulary.Question)

	return h.api.SendMessageWithAttachmentAndButton(peerId, question, vocabulary.Attachment, keyboard)
}
