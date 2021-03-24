package comparator

import "github.com/sepuka/focalism/internal/domain"

type (
	irregularComparator struct {
	}
)

func NewIrregularComparator() domain.Comparator {
	return &irregularComparator{}
}

func (c *irregularComparator) Compare(vocabulary *domain.Vocabulary, msg string) bool {
	return vocabulary.Answer == msg
}
