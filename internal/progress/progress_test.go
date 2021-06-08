package progress

import (
	"fmt"
	"github.com/sepuka/focalism/errors"
	"github.com/sepuka/focalism/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestScheduleProgress(t *testing.T) {
	var (
		actualErr      error
		actualProgress *domain.Progress
		failMsg        string
		testCases      = map[string]struct {
			srcProgress      *domain.Progress
			expectedProgress *domain.Progress
			err              error
		}{
			`limit is reached`: {
				srcProgress: &domain.Progress{
					Counter: maxAttempts + 1,
				},
				expectedProgress: nil,
				err:              errors.LearnerLimitError,
			},
			`1 shows`: {
				srcProgress: &domain.Progress{
					Counter: 0,
					Date:    time.Now(),
				},
				expectedProgress: &domain.Progress{
					Counter: 1,
					Date:    time.Now().Add(day),
				},
				err: nil,
			},
			`2 shows`: {
				srcProgress: &domain.Progress{
					Counter: 1,
					Date:    time.Now(),
				},
				expectedProgress: &domain.Progress{
					Counter: 2,
					Date:    time.Now().Add(2 * day),
				},
				err: nil,
			},
			`3 shows`: {
				srcProgress: &domain.Progress{
					Counter: 2,
					Date:    time.Now(),
				},
				expectedProgress: &domain.Progress{
					Counter: 3,
					Date:    time.Now().Add(4 * day),
				},
				err: nil,
			},
			`4 shows`: {
				srcProgress: &domain.Progress{
					Counter: 3,
					Date:    time.Now(),
				},
				expectedProgress: &domain.Progress{
					Counter: 4,
					Date:    time.Now().Add(6 * day),
				},
				err: nil,
			},
			`5 shows`: {
				srcProgress: &domain.Progress{
					Counter: 4,
					Date:    time.Now(),
				},
				expectedProgress: &domain.Progress{
					Counter: 5,
					Date:    time.Now().Add(8 * day),
				},
				err: nil,
			},
			`6 shows`: {
				srcProgress: &domain.Progress{
					Counter: 5,
					Date:    time.Now(),
				},
				expectedProgress: &domain.Progress{
					Counter: 6,
					Date:    time.Now().Add(10 * day),
				},
				err: nil,
			},
			`7 shows`: {
				srcProgress: &domain.Progress{
					Counter: 6,
					Date:    time.Now(),
				},
				expectedProgress: &domain.Progress{
					Counter: 7,
					Date:    time.Now().Add(12 * day),
				},
				err: nil,
			},
		}
	)

	for testName, testCase := range testCases {
		failMsg = fmt.Sprintf(`test case "%s" failure`, testName)
		actualProgress, actualErr = ScheduleProgress(testCase.srcProgress)
		assert.ErrorIs(t, actualErr, testCase.err, failMsg)
		if actualErr == nil {
			assert.Equal(t, testCase.expectedProgress.Counter, actualProgress.Counter, failMsg)
			assert.Equal(t, testCase.expectedProgress.PeerId, actualProgress.PeerId, failMsg)
			assert.InDelta(t, testCase.expectedProgress.Date.Unix(), actualProgress.Date.Unix(), 1, failMsg)
		}
	}
}
