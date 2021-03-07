package domain

type (
	TopicRepository interface {
		GetList(name Name) ([]Topic, error)
	}

	Topic struct {
		TopicId int64  `sql:"topic_id,pk"`
		Title   string `sql:"title"`
		ModeId  int    `sql:"mode_id"`
		Mode    *Mode  `pg:"rel:belongs-to"`
	}
)
