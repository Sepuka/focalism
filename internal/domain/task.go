package domain

import (
	"strconv"
	"time"
)

type (
	TaskRepository interface {
		Create(vocabulary *Vocabulary, peerId int64) (*Task, error)
		GetById(taskId int64) (*Task, error)
		GetLast() (Task, error)
		Answer(task Task) error
	}

	Task struct {
		Id           int64         `sql:"task_id,pk"`
		VocabularyId int64         `sql:"vocabulary_id"`
		PeerId       int64         `sql:"peer_id"`
		Datetime     time.Time     `sql:"datetime"`
		Time         time.Duration `sql:"time"`
		Vocabulary   *Vocabulary   `pg:"rel:belongs-to"`
	}
)

func (v *Task) GetId() string {
	return strconv.FormatInt(v.Id, 10)
}
