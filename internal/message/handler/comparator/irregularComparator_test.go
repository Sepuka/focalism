package comparator

import (
	"fmt"
	"github.com/sepuka/focalism/internal/domain"
	"testing"
)

func TestCompare(t *testing.T) {
	const (
		varMsgTmpl = `unexpected result for test "%s": expected %t but got %t`
	)
	var (
		comparator = NewIrregularComparator()
		actualResult bool
		tests = map[string]struct{
			vocabulary *domain.Vocabulary
			msg string
			isMatched bool
		}{
			`empty message and no vocabulary`: {vocabulary: nil, msg: ``, isMatched: false},
			`simple message`: {vocabulary: &domain.Vocabulary{Answer: `simple`}, msg: `simple`, isMatched: true},
			`capital letter`: {vocabulary: &domain.Vocabulary{Answer: `capital`}, msg: `Capital`, isMatched: true},
			`with commas`: {vocabulary: &domain.Vocabulary{Answer: `be,was/were,been`}, msg: `be,was/were,been`, isMatched: true},
			`with spaces`: {vocabulary: &domain.Vocabulary{Answer: `be,was/were,been`}, msg: `be was/were been`, isMatched: true},
			`with dash`: {vocabulary: &domain.Vocabulary{Answer: `be,was/were,been`}, msg: `be-was/were-been`, isMatched: true},
		}
	)

	for testKey, testValue := range tests {
		actualResult = comparator.Compare(testValue.vocabulary, testValue.msg)
		if actualResult != testValue.isMatched {
			t.Error(fmt.Sprintf(varMsgTmpl, testKey, testValue.isMatched, actualResult))
		}
	}
}
