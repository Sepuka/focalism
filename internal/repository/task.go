package repository

import (
	"github.com/go-pg/pg"
	"github.com/sepuka/focalism/errors"
	"github.com/sepuka/focalism/internal/domain"
	"time"
)

type (
	TaskRepository struct {
		db *pg.DB
	}
)

func NewTaskRepository(db *pg.DB) domain.TaskRepository {
	return &TaskRepository{db: db}
}

func (v *TaskRepository) Create(vocabulary *domain.Vocabulary, peerId int64) (*domain.Task, error) {
	var (
		err  error
		task = &domain.Task{
			VocabularyId: vocabulary.Id,
			PeerId:       peerId,
			Datetime:     time.Now(),
		}
	)

	_, err = v.
		db.
		Model(task).
		Insert()

	return task, err
}

func (v *TaskRepository) GetById(taskId int64) (*domain.Task, error) {
	var (
		err  error
		task = &domain.Task{
			Id: taskId,
		}
	)

	err = v.
		db.
		Model(task).
		Column(`task.*`).
		Relation(`Vocabulary`).
		WherePK().
		Select()

	if err != nil {
		return nil, err
	}

	return task, err
}

func (v *TaskRepository) GetLast() (domain.Task, error) {
	var (
		err  error
		task = domain.Task{}
	)

	err = v.
		db.
		Model(&task).
		Column(`task.*`).
		Relation(`Vocabulary`).
		Where(`time is NULL`).
		Order(`task_id DESC`).
		Limit(1).
		Select()

	if err == pg.ErrNoRows {
		return task, errors.NewDatabaseNoRowsError(`There's not any waiting tasks`, err)
	}

	return task, err
}

func (v *TaskRepository) Answer(task domain.Task) error {
	var (
		err error
	)

	task.Time = time.Now().Sub(task.Datetime)

	err = v.
		db.
		Update(task)

	return err
}
