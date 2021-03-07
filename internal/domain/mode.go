package domain

const (
	IrregularMode Name = `irregularMode`
	TopicMode     Name = `topicMode`
)

type (
	Name string
	Mode struct {
		Id     int    `sql:"mode_id,pk"`
		Marker string `sql:"name"`
	}
)
