package progress

import (
	"github.com/sepuka/focalism/errors"
	"github.com/sepuka/focalism/internal/domain"
	"time"
)

const (
	maxAttempts = 7
	day         = 24 * time.Hour
)

var (
	periods = map[int8]time.Duration{
		0: day,      // next day
		1: 2 * day,  // 3 days
		2: 4 * day,  // 7 days
		3: 6 * day,  // 13 days
		4: 8 * day,  // 21 days
		5: 10 * day, // 31 days
		6: 12 * day, // 43 days
	}
)

func NewProgress(vocabulary *domain.Vocabulary, peerId int64) *domain.Progress {
	return &domain.Progress{
		VocabularyId: vocabulary.Id,
		PeerId:       peerId,
		Counter:      0,
		Vocabulary:   vocabulary,
		Date:         time.Now(),
	}
}

func ScheduleProgress(progress *domain.Progress) (*domain.Progress, error) {
	if progress.Counter >= maxAttempts {
		return nil, errors.NewLearnerBuilderError(`attempts limit is reached`)
	}

	var (
		result = &domain.Progress{
			VocabularyId: progress.VocabularyId,
			PeerId:       progress.PeerId,
			Date:         progress.Date,
			Counter:      progress.Counter,
		}
	)

	if duration, ok := periods[progress.Counter]; ok {
		result.Counter++
		result.Date = result.Date.Add(duration)
	}

	return result, nil
}
