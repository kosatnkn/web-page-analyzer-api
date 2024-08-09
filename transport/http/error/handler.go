package error

import (
	"context"
	"net/http"

	"github.com/kosatnkn/web-page-analyzer-api/app/adapters"

	domainErrs "github.com/kosatnkn/web-page-analyzer-api/domain/errors"
	serviceErrs "github.com/kosatnkn/web-page-analyzer-api/externals/services/errors"
	middlewareErrs "github.com/kosatnkn/web-page-analyzer-api/transport/http/middleware/errors"
	unpackerErrs "github.com/kosatnkn/web-page-analyzer-api/transport/http/request/unpackers/errors"
	transformerErrs "github.com/kosatnkn/web-page-analyzer-api/transport/http/response/transformers/errors"
)

// Handle handles all errors globally.
func Handle(ctx context.Context, err error, log adapters.LogAdapterInterface) (interface{}, int) {
	switch err.(type) {
	case *transformerErrs.TransformerError:
		logError(ctx, log, err)
		return formatGenericError(err), http.StatusInternalServerError
	case *middlewareErrs.MiddlewareError,
		*domainErrs.DomainError,
		*serviceErrs.ServiceError:
		logError(ctx, log, err)
		return formatGenericError(err), http.StatusBadRequest
	case *unpackerErrs.UnpackerError:
		logError(ctx, log, err)
		return formatUnpackerError(err), http.StatusUnprocessableEntity
	default:
		logError(ctx, log, err)
		return formatUnknownError(err), http.StatusInternalServerError
	}
}

// HandleValidatorErrors specifically handles validation errors thrown by the validator.
func HandleValidatorErrors(ctx context.Context, errs map[string]string, log adapters.LogAdapterInterface) (interface{}, int) {
	log.Error(ctx, "Validation Errors", errs)
	return formatValidatorErrors(errs), http.StatusUnprocessableEntity
}
