package repository

import (
	"github.com/go-pg/pg"
	"github.com/sepuka/focalism/internal/domain"
	"time"
)

type (
	VocabularyRepository struct {
		db *pg.DB
	}
)

func NewVocabularyRepository(db *pg.DB) domain.VocabularyRepository {
	return &VocabularyRepository{db: db}
}

func (v *VocabularyRepository) FindActual(topicId int64, peerId int64) (domain.Vocabulary, error) {
	var (
		err        error
		vocabulary = domain.Vocabulary{}
	)

	err = v.
		db.
		Model(&vocabulary).
		Join(`LEFT OUTER JOIN tasks AS task`).
		JoinOn(`task.vocabulary_id = vocabulary.vocabulary_id AND DATE(datetime) = ? AND peer_id = ?`, time.Now().Format(`2006-01-02`), peerId).
		Where(`topic_id = ? AND task_id IS NULL`, topicId).
		Limit(1).
		Select()

	return vocabulary, err
}

func (v *VocabularyRepository) IncrViews(vocabulary domain.Vocabulary) {
	_, _ = v.
		db.
		Model(&vocabulary).
		Set(`views = views + 1`).
		WherePK().
		Update()
}
