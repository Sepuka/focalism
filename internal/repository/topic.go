package repository

import (
	"github.com/go-pg/pg"
	"github.com/sepuka/focalism/internal/domain"
)

type (
	TopicRepository struct {
		db *pg.DB
	}
)

func NewTopicRepository(db *pg.DB) *TopicRepository {
	return &TopicRepository{db: db}
}

func (t *TopicRepository) GetList(name domain.Name) ([]domain.Topic, error) {
	var (
		err  error
		list []domain.Topic
	)

	err = t.
		db.
		Model(&list).
		Relation(`Mode`).
		Where(`mode.name = ?`, name).
		Select()

	return list, err
}
