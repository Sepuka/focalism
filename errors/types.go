package errors

import "errors"

var (
	NoError = errors.New(`not an error`)
	InvalidJson = errors.New(`invalid JSON`)
	DatabaseError = errors.New(`database error`)
)

func NewInvalidJsonError(msg string, originalErr error) FocalismError {
	return FocalismError{
		err:           InvalidJson,
		message:       msg,
		originalError: originalErr,
	}
}

func NewDatabaseError(msg string, originalErr error) FocalismError {
	return FocalismError{
		err:           DatabaseError,
		message:       msg,
		originalError: originalErr,
	}
}

func NewDatabaseNoRowsError(msg string, originalErr error) FocalismError {
	return FocalismError{
		err:           NoError,
		message:       msg,
		originalError: originalErr,
	}
}