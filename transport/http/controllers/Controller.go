package controllers

import (
	"context"
	"io"
	"net/http"

	"githubcom/kosatnkn/web-page-analyzer-api/app/adapters"
	"githubcom/kosatnkn/web-page-analyzer-api/app/container"
	"githubcom/kosatnkn/web-page-analyzer-api/transport/http/request"
	"githubcom/kosatnkn/web-page-analyzer-api/transport/http/request/unpackers"
	"githubcom/kosatnkn/web-page-analyzer-api/transport/http/response"
	"githubcom/kosatnkn/web-page-analyzer-api/transport/http/response/transformers"

	"github.com/gorilla/mux"
)

// Controller is the base struct that holds fields and functionality common to all controllers.
type Controller struct {
	logger    adapters.LogAdapterInterface
	validator adapters.ValidatorAdapterInterface
}

// NewController creates a new instance of the controller.
func NewController(c *container.Container) *Controller {
	return &Controller{
		logger:    c.Adapters.Log,
		validator: c.Adapters.Validator,
	}
}

// withTrace adds an optional tracing string that will be displayed in error messages.
func (ctl *Controller) withTrace(ctx context.Context, point string) context.Context {
	return ctl.logger.AppendTracePoint(ctx, point)
}

// routeVar returns the value of the route variable denoted by the name.
func (ctl *Controller) routeVar(r *http.Request, name string) string {
	return mux.Vars(r)[name]
}

// urlParam returns the value of the url parameter denoted by the name.
func (ctl *Controller) urlParam(r *http.Request, name string) string {
	p := r.URL.Query()[name]
	if len(p) > 0 {
		return p[0]
	}

	return ""
}

// unpackBody unpacks and validates the request body.
func (ctl *Controller) unpackBody(r *http.Request, u unpackers.UnpackerInterface) interface{} {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = request.Unpack(body, u)
	if err != nil {
		return err
	}

	// validate unpacked data
	errs := ctl.validator.Validate(u)
	if errs != nil {
		return errs
	}

	return nil
}

// transform is a convenience function wrapping the actual `response.Transform` function
// to provide a cleaner usage interface.
func (ctl *Controller) transform(data interface{}, t transformers.TransformerInterface, isCollection bool) (interface{}, error) {
	return response.Transform(data, t, isCollection)
}

// sendResponse is a convenience function wrapping the actual `response.Send` function
// to provide a cleaner usage interface.
func (ctl *Controller) sendResponse(ctx context.Context, w http.ResponseWriter, code int, payload ...interface{}) {
	if len(payload) == 0 {
		response.Send(w, code, nil)
		return
	}

	response.Send(w, code, payload)
}

// sendError is a convenience function wrapping the actual `response.Error` function
// to provide a cleaner usage interface.
func (ctl *Controller) sendError(ctx context.Context, w http.ResponseWriter, err interface{}) {
	response.Error(ctx, w, ctl.logger, err)
}
