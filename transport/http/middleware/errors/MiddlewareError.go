package errors

import e "githubcom/kosatnkn/web-page-analyzer-api/errors"

// MiddlewareError is the type of errors thrown by middleware.
type MiddlewareError struct {
	*e.BaseError
}

// NewMiddlewareError creates a new MiddlewareError instance.
func NewMiddlewareError(code, msg string, errs error) error {
	return &MiddlewareError{
		BaseError: e.NewBaseError("MiddlewareError", code, msg, errs),
	}
}
