package comparator

import (
	"fmt"
	"github.com/sepuka/focalism/internal/domain"
	"testing"
)

func TestSimpleCompare(t *testing.T) {
	const (
		varMsgTmpl = `unexpected result for test "%s": expected %t but got %t`
	)
	var (
		comparator   = NewSimpleComparator()
		actualResult bool
		tests        = map[string]struct {
			vocabulary *domain.Vocabulary
			msg        string
			isMatched  bool
		}{
			`empty message and no vocabulary`: {vocabulary: &domain.Vocabulary{Answer: `blahblah`}, msg: `BlahBlah`, isMatched: true},
		}
	)

	for testKey, testValue := range tests {
		actualResult = comparator.Compare(testValue.vocabulary, testValue.msg)
		if actualResult != testValue.isMatched {
			t.Error(fmt.Sprintf(varMsgTmpl, testKey, testValue.isMatched, actualResult))
		}
	}
}
