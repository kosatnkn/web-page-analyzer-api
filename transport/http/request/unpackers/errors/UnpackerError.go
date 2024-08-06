package errors

import e "githubcom/kosatnkn/web-page-analyzer-api/errors"

// UnpackerError is the type of errors thrown by response transformers.
type UnpackerError struct {
	*e.BaseError
}

// NewUnpackerError creates a new UnpackerError instance.
func NewUnpackerError(code, msg string, cause error) error {
	return &UnpackerError{
		BaseError: e.NewBaseError("UnpackerError", code, msg, cause),
	}
}
