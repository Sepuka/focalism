package masker

import (
	"fmt"
	"github.com/sepuka/focalism/internal/domain"
)

type (
	irregularMasker struct {
	}
)

func NewIrregularMasker() domain.Masker {
	return &irregularMasker{}
}

func (m *irregularMasker) Mask(vocabulary *domain.Vocabulary) string {
	return fmt.Sprintf(`"%c***"`, vocabulary.Answer[0])
}
