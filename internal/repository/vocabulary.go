package repository

import (
	"github.com/go-pg/pg"
	"github.com/sepuka/focalism/internal/domain"
)

type (
	VocabularyRepository struct {
		db *pg.DB
	}
)

func NewVocabularyRepository(db *pg.DB) domain.VocabularyRepository {
	return &VocabularyRepository{db: db}
}

func (v *VocabularyRepository) FindActual(topicId int64) (domain.Vocabulary, error) {
	var (
		err        error
		vocabulary = domain.Vocabulary{}
	)

	err = v.
		db.
		Model(&vocabulary).
		Where(`topic_id = ?`, topicId).
		Order(`views ASC`).
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
