package domain

import (
	"strconv"
	"time"
)

type (
	TaskRepository interface {
		Create(vocabulary Vocabulary, peerId int64) (*Task, error)
		GetById(taskId int64) (Task, error)
		GetLast() (Task, error)
		Answer(task Task) error
		GetTodayTasks(topicId int64, peerId int) (int, error)
	}

	TaskProgressRepository interface {
		GetProgress(topicId int64, peerId int64) (success int, attempts int, err error)
	}

	Comparator interface {
		Compare(vocabulary *Vocabulary, msg string) bool
	}

	Task struct {
		Id           int64       `sql:"task_id,pk"`
		VocabularyId int64       `sql:"vocabulary_id"`
		PeerId       int64       `sql:"peer_id"`
		Datetime     time.Time   `sql:"datetime"`
		Time         int64       `sql:"time"`
		Vocabulary   *Vocabulary `pg:"rel:belongs-to"`
		IsCorrect    bool        `sql:"is_correct"`
	}
)

func (v *Task) GetId() string {
	return strconv.FormatInt(v.Id, 10)
}
