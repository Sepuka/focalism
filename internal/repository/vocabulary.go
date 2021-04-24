package repository

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/sepuka/focalism/errors"
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

func (v *VocabularyRepository) FindActual(topicId int64, peerId int64, date time.Time) (domain.Vocabulary, error) {
	var (
		err        error
		vocabulary = domain.Vocabulary{}
		subQuery   *orm.Query
	)

	subQuery = v.
		db.
		Model((*domain.Task)(nil)).
		ColumnExpr(`vocabulary_id`).
		Where(`DATE(task.datetime) >= ? AND peer_id = ? AND is_correct = true`, date.Format(`2006-01-02`), peerId)

	err = v.
		db.
		Model(&vocabulary).
		Relation(`Topic`).
		Where(`vocabulary.topic_id = ? AND vocabulary_id NOT IN (?)`, topicId, subQuery).
		Limit(1).
		Select()

	if err == pg.ErrNoRows {
		return vocabulary, errors.NewDatabaseNoRowsError(`There were not any vocabularies found`, err)
	}

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

func (v *VocabularyRepository) GetTotal(topicId int64) (int, error) {
	var (
		err   error
		total int
	)

	total, err = v.
		db.
		Model(&domain.Vocabulary{}).
		Where(`topic_id = ?`, topicId).
		Count()

	return total, err
}
