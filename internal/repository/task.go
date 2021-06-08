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

func (v *TaskRepository) Create(vocabulary domain.Vocabulary, peerId int64) (*domain.Task, error) {
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

func (v *TaskRepository) GetById(taskId int64) (domain.Task, error) {
	var (
		err  error
		task = domain.Task{
			Id: taskId,
		}
	)

	err = v.
		db.
		Model(&task).
		Column(`task.*`).
		Relation(`Vocabulary`).
		Relation(`Vocabulary.Topic`).
		Relation(`Vocabulary.Topic.Mode`).
		WherePK().
		Select()

	if err != nil {
		return task, err
	}

	return task, err
}

func (v *TaskRepository) GetLastUnanswered(peerId int64) (*domain.Task, error) {
	var (
		err  error
		task = &domain.Task{}
	)

	err = v.
		db.
		Model(&task).
		Column(`task.*`).
		Relation(`Vocabulary`).
		Relation(`Vocabulary.Topic`).
		Relation(`Vocabulary.Topic.Mode`).
		Where(`time is NULL AND peer_id = ?`, peerId).
		Order(`task_id DESC`).
		Limit(1).
		Select()

	if err == pg.ErrNoRows {
		return task, errors.NewDatabaseNoRowsError(`There's not any waiting tasks`, err)
	}

	return task, err
}

func (v *TaskRepository) Answer(task *domain.Task) error {
	var (
		err   error
		query = v.db.Model(&task)
	)

	task.Time = int64(time.Now().Sub(task.Datetime).Seconds())

	if task.IsCorrect {
		query.Column(`time`, `is_correct`)
	} else {
		query.Column(`time`)
	}

	_, err = query.
		WherePK().
		Update()

	return err
}

func (v *TaskRepository) GetTodayTasks(topicId int64, peerId int) (int, error) {
	var (
		err error
		cnt int
	)

	cnt, err = v.
		db.
		Model(&domain.Task{}).
		Relation(`Vocabulary`).
		Where(`peer_id = ? AND DATE(datetime) = ? AND topic_id = ? AND is_correct = true`, peerId, time.Now().Format(`2006-01-02`), topicId).
		Count()

	return cnt, err
}

func (v *TaskRepository) GetProgress(topicId int64, peerId int64) (success int, total int, err error) {
	var (
		tasks []domain.Task
		task  domain.Task
	)

	err = v.
		db.
		Model(&tasks).
		Column(`task.is_correct`).
		Relation(`Vocabulary`).
		Where(`peer_id = ? AND vocabulary.topic_id = ? AND DATE(datetime) = ?`, peerId, topicId, time.Now().Format(`2006-01-02`)).
		Select()

	if err != nil {
		return 0, 0, err
	}

	total = len(tasks)

	for _, task = range tasks {
		if task.IsCorrect {
			success++
		}
	}

	return success, total, err
}

func (v *TaskRepository) GetAverage(peerId int64) (domain.TaskProgress, error) {
	var (
		err    error
		result domain.TaskProgress
	)

	err = v.
		db.
		Model((*domain.Task)(nil)).
		ColumnExpr(`avg(time) AS total_average`).
		ColumnExpr(`count(*) AS total_attempts`).
		Where(`peer_id = ? AND time IS NOT NULL`, peerId).
		Select(&result)

	return result, err
}

func (v *TaskRepository) DeleteLast(peerId int64) error {
	var err error

	_, err = v.
		db.
		Model((*domain.Task)(nil)).
		Where(`peer_id = ? AND time IS NULL`, peerId).
		Delete()

	return err
}
