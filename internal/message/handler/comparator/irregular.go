package comparator

import (
	"github.com/sepuka/focalism/internal/domain"
	"regexp"
	"strings"
)

type (
	irregularComparator struct {
	}
)

func NewIrregularComparator() domain.Comparator {
	return &irregularComparator{}
}

func (c *irregularComparator) Compare(vocabulary *domain.Vocabulary, msg string) bool {
	if vocabulary == nil {
		return false
	}

	const (
		maxWords      = 4
		regexpPattern = `[,\ -]`
	)

	var (
		answer           = strings.ToLower(vocabulary.Answer)
		expected, actual []string
		index            = 0
	)

	expected = regexp.MustCompile(regexpPattern).Split(answer, maxWords)
	actual = regexp.MustCompile(regexpPattern).Split(strings.ToLower(msg), maxWords)

	if len(expected) > len(actual) {
		return false
	}

	for _, word := range expected {
		if actual[index] == `` {
			continue
		}

		if strings.Trim(actual[index], ` `) != word {
			return false
		}

		index++
	}

	return true
}
