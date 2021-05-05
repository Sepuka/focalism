package masker

import (
	"fmt"
	"github.com/sepuka/focalism/internal/domain"
	"strings"
)

type (
	simpleMasker struct {
	}
)

func NewSimpleMasker() *simpleMasker {
	return &simpleMasker{}
}

func (m *simpleMasker) Mask(vocabulary *domain.Vocabulary) string {
	var (
		length = len(vocabulary.Answer)
		answer string
	)

	if length > 2 {
		stars := strings.Repeat(`*`, length-2)
		answer = fmt.Sprintf(`"%c%s%c"`, vocabulary.Answer[0], stars, vocabulary.Answer[length-1])
	} else {
		answer = fmt.Sprintf(`"%c**"`, vocabulary.Answer[0])
	}

	if len(vocabulary.Example) > 0 {
		answer = fmt.Sprintf("%s\n\n%s", answer, vocabulary.Example)
	}

	return fmt.Sprintf(`%s`, answer)
}
