package comparator

import (
	"github.com/sepuka/focalism/internal/domain"
	"strings"
)

type (
	simpleComparator struct {
	}
)

func NewSimpleComparator() domain.Comparator {
	return &simpleComparator{}
}

func (c *simpleComparator) Compare(vocabulary *domain.Vocabulary, msg string) bool {
	return strings.ToLower(vocabulary.Answer) == strings.ToLower(msg)
}
