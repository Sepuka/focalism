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
		comparator   = NewIrregularComparator()
		actualResult bool
		tests        = map[string]struct {
			vocabulary *domain.Vocabulary
			msg        string
			isMatched  bool
		}{
			`empty message and no vocabulary`: {vocabulary: nil, msg: ``, isMatched: false},
			`a lot of words via space`:        {vocabulary: &domain.Vocabulary{Answer: `first`}, msg: `first second`, isMatched: true}, // we don't care about a tail
			`a few words via space`:           {vocabulary: &domain.Vocabulary{Answer: `first,second`}, msg: `first`, isMatched: false},
			`simple message`:                  {vocabulary: &domain.Vocabulary{Answer: `simple`}, msg: `simple`, isMatched: true},
			`capital letter`:                  {vocabulary: &domain.Vocabulary{Answer: `capital`}, msg: `Capital`, isMatched: true},
			`with commas`:                     {vocabulary: &domain.Vocabulary{Answer: `be,was/were,been`}, msg: `be,was/were,been`, isMatched: true},
			`with spaces`:                     {vocabulary: &domain.Vocabulary{Answer: `be,was/were,been`}, msg: `be was/were been`, isMatched: true},
			`with dash`:                       {vocabulary: &domain.Vocabulary{Answer: `be,was/were,been`}, msg: `be-was/were-been`, isMatched: true},
			`commas and spaces`:               {vocabulary: &domain.Vocabulary{Answer: `deal,dealt,dealt`}, msg: `deal, dealt,dealt`, isMatched: true},
			`commas and many spaces`:          {vocabulary: &domain.Vocabulary{Answer: `deal,dealt,dealt`}, msg: `deal,   dealt ,dealt`, isMatched: true},
			`comma with space`:                {vocabulary: &domain.Vocabulary{Answer: `drink,drank,drunk`}, msg: `drink, drank, drunk`, isMatched: true},
		}
	)

	for testKey, testValue := range tests {
		actualResult = comparator.Compare(testValue.vocabulary, testValue.msg)
		if actualResult != testValue.isMatched {
			t.Error(fmt.Sprintf(varMsgTmpl, testKey, testValue.isMatched, actualResult))
		}
	}
}
