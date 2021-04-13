package errors

import "errors"

var (
	NoError             = errors.New(`not an error`)
	InvalidJson         = errors.New(`invalid JSON`)
	DatabaseError       = errors.New(`database error`)
	NoRowsDatabaseError = errors.New(`empty query result`)
	InternalError       = errors.New(`internal error`)
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
		err:           NoRowsDatabaseError,
		message:       msg,
		originalError: originalErr,
	}
}

func NewInternalError(msg string, originalErr error) FocalismError {
	return FocalismError{
		err:           InternalError,
		message:       msg,
		originalError: originalErr,
	}
}
