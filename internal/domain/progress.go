package domain

import "time"

type (
	ProgressRepository interface {
		Fetch(task *Task) (*Progress, error)
		Schedule(learner *Progress) error
	}

	Progress struct {
		TableName    struct{}  `sql:"progress"`
		VocabularyId int64     `sql:"vocabulary_id"`
		PeerId       int64     `sql:"peer_id"`
		Date         time.Time `sql:"date"`
		Counter      int8      `sql:"counter"`
		Vocabulary   *Vocabulary
	}
)
