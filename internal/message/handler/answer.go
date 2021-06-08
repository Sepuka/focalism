package handler

import (
	"fmt"
	errors2 "github.com/sepuka/focalism/errors"
	"github.com/sepuka/focalism/internal/domain"
	button2 "github.com/sepuka/focalism/internal/message/button"
	"github.com/sepuka/focalism/internal/message/handler/comparator"
	progress2 "github.com/sepuka/focalism/internal/progress"
	"github.com/sepuka/vkbotserver/api"
	"github.com/sepuka/vkbotserver/api/button"
	domain2 "github.com/sepuka/vkbotserver/domain"
	"go.uber.org/zap"
)

type (
	Answer struct {
		taskRepo     domain.TaskRepository
		progressRepo domain.ProgressRepository
		api          *api.Api
		log          *zap.SugaredLogger
	}
)

// NewAnswer is the answers handler
func NewAnswer(
	taskRepository domain.TaskRepository,
	progressRepository domain.ProgressRepository,
	api *api.Api,
	log *zap.SugaredLogger,
) *Answer {
	return &Answer{
		taskRepo:     taskRepository,
		progressRepo: progressRepository,
		api:          api,
		log:          log,
	}
}

func (h *Answer) Handle(req *domain2.Request) error {
	var (
		err      error
		answer   = `правильно! ещё?`
		msg      = req.Object.Message.Text
		peerId   = int(req.Object.Message.FromId)
		lastTask *domain.Task
		keyboard = button.Keyboard{
			OneTime: true,
		}
		vocabulary *domain.Vocabulary
	)

	if lastTask, err = h.taskRepo.GetLastUnanswered(int64(peerId)); err != nil {
		if err.(errors2.FocalismError).Is(errors2.NoRowsDatabaseError) {
			keyboard.Buttons = button2.ModeChoose()
			return h.api.SendMessageWithButton(peerId, `выберите режим для занятий`, keyboard)
		}

		h.
			log.
			With(
				zap.Int(`peerId`, peerId),
				zap.Error(err),
			).
			Error(`cannot find last task`)

		keyboard.Buttons = button2.Return()

		return h.api.SendMessageWithButton(peerId, `что-то пошло не так`, keyboard)
	}

	vocabulary = lastTask.Vocabulary
	if h.Comparator(vocabulary).Compare(vocabulary, msg) {
		lastTask.IsCorrect = true
	} else {
		answer = fmt.Sprintf(`вы ошиблись, правильный ответ: "%s"`, vocabulary.Answer)
	}

	if err = h.taskRepo.Answer(lastTask); err != nil {
		h.
			log.
			With(
				zap.Int(`peerId`, peerId),
				zap.Int64(`taskId`, lastTask.Id),
				zap.Error(err),
			).
			Error(`answer time update failed`)
	}

	go h.schedule(lastTask, vocabulary)

	keyboard.Buttons = button2.NextWithReturnAndProgress(fmt.Sprintf(`%d`, lastTask.Vocabulary.TopicId))

	return h.api.SendMessageWithButton(peerId, answer, keyboard)
}

func (h *Answer) schedule(task *domain.Task, vocabulary *domain.Vocabulary) {
	var (
		progress *domain.Progress
		err      error
		peerId   = task.PeerId
	)

	if progress, err = h.progressRepo.Fetch(task); err != nil {
		if err.(errors2.FocalismError).Is(errors2.NoRowsDatabaseError) {
			progress = progress2.NewProgress(vocabulary, peerId)
		} else {
			h.
				log.
				With(
					zap.Int64(`peerId`, peerId),
					zap.Int64(`taskId`, task.Id),
					zap.Error(err),
				).
				Error(`could fetch progress row`)
			return
		}
	}

	if progress, err = progress2.ScheduleProgress(progress); err != nil {
		if err.(errors2.FocalismError).Is(errors2.LearnerLimitError) {
			h.
				log.
				With(
					zap.Int64(`peerId`, peerId),
					zap.Int64(`taskId`, task.Id),
				).
				Debug(`word is learned`)
		} else {
			h.
				log.
				With(
					zap.Int64(`peerId`, peerId),
					zap.Int64(`taskId`, task.Id),
					zap.Error(err),
				).
				Error(`could not calculate new date`)
		}

		return
	}

	if err = h.progressRepo.Schedule(progress); err != nil {
		h.log.With(
			zap.Int64(`peerId`, peerId),
			zap.Int64(`taskId`, task.Id),
			zap.Error(err),
		).
			Error(`could not scheduled task`)
	}
}

func (h *Answer) Comparator(vocabulary *domain.Vocabulary) domain.Comparator {
	switch vocabulary.Topic.Mode.Marker {
	case string(domain.IrregularMode):
		return comparator.NewIrregularComparator()
	default:
		return comparator.NewSimpleComparator()
	}
}
