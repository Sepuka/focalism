package repository

import (
	"github.com/go-pg/pg"
	"github.com/sepuka/focalism/errors"
	"github.com/sepuka/focalism/internal/domain"
)

type (
	ProgressRepository struct {
		db *pg.DB
	}
)

func NewProgressRepository(db *pg.DB) domain.ProgressRepository {
	return &ProgressRepository{db: db}
}

func (r *ProgressRepository) Fetch(task *domain.Task) (*domain.Progress, error) {
	var (
		progress domain.Progress
		err      error
	)

	err = r.
		db.
		Model(&progress).
		Relation(`Vocabulary`).
		Where(`progress.vocabulary_id = ? AND progress.peer_id = ?`, task.VocabularyId, task.PeerId).
		Select()

	if err == pg.ErrNoRows {
		return &progress, errors.NewDatabaseNoRowsError(`There's not any progress tasks`, err)
	}

	return &progress, err
}

func (r *ProgressRepository) Schedule(progress *domain.Progress) error {
	var (
		err error
	)

	_, err = r.
		db.
		Model(progress).
		OnConflict(`(vocabulary_id, peer_id) DO UPDATE`).
		Insert()

	return err
}
