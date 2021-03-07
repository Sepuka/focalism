package domain

import "strconv"

type (
	VocabularyRepository interface {
		FindActual(topicId int64) (*Vocabulary, error)
		IncrViews(*Vocabulary)
	}

	Vocabulary struct {
		Id         int64  `sql:"vocabulary_id,pk"`
		Attachment string `sql:"attachment"`
		Views      int64  `sql:"views"`
		Answer     string `sql:"answer"`
		Example    string `sql:"example"`
		Question   string `sql:"question"`
	}
)

func (v *Vocabulary) GetId() string {
	return strconv.FormatInt(v.Id, 10)
}
