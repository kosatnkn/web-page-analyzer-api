package errors

import e "githubcom/kosatnkn/web-page-analyzer-api/errors"

// TransformerError is the type of errors thrown by response transformers.
type TransformerError struct {
	*e.BaseError
}

// NewTransformerError creates a new TransformerError instance.
func NewTransformerError(code, msg string, cause error) error {
	return &TransformerError{
		BaseError: e.NewBaseError("TransformerError", code, msg, cause),
	}
}
