package errors

import (
	"fmt"
)

// Common error
type FocalismError struct {
	err           error
	message       string
	originalError error
	context       map[string]string
}

func (e FocalismError) Error() string {
	if e.originalError != nil {
		return fmt.Sprintf(`%s (%s)`, e.message, e.originalError)
	}

	return fmt.Sprintf(`%s`, e.message)
}

func (e FocalismError) Is(target error) bool {
	return e.err == target
}
