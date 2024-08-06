package controllers

import (
	"net/http"

	"githubcom/kosatnkn/web-page-analyzer-api/app/container"
	"githubcom/kosatnkn/web-page-analyzer-api/domain/usecases/sample"
	"githubcom/kosatnkn/web-page-analyzer-api/transport/http/request/unpackers"
	"githubcom/kosatnkn/web-page-analyzer-api/transport/http/response/transformers"
)

// SampleController contains controller logic for endpoints.
type SampleController struct {
	*Controller
	sampleUseCase *sample.Sample
}

// NewSampleController creates a new instance of the controller.
func NewSampleController(c *container.Container) *SampleController {
	return &SampleController{
		Controller:    NewController(c),
		sampleUseCase: sample.NewSample(c),
	}
}

// Get handles retrieving a list of samples.
func (ctl *SampleController) Get(w http.ResponseWriter, r *http.Request) {
	// add a trace string to the request context
	ctx := ctl.withTrace(r.Context(), "SampleController.Get")

	// get filters from query params
	filters, err := ctl.filters(r, unpackers.NewSampleFiltersUnpacker())
	if err != nil {
		ctl.sendError(ctx, w, err)
		return
	}

	// get paginator from query parameters
	paginator, err := ctl.paginator(r)
	if err != nil {
		ctl.sendError(ctx, w, err)
		return
	}

	// get data
	samples, err := ctl.sampleUseCase.Get(ctx, filters, paginator)
	if err != nil {
		ctl.sendError(ctx, w, err)
		return
	}

	// transform school data
	trS, err := ctl.transform(samples, transformers.NewSampleTransformer(), true)
	if err != nil {
		ctl.sendError(ctx, w, err)
		return
	}

	// transform paginator
	trP, err := ctl.transform(paginator, transformers.NewPaginatorTransformer(), false)
	if err != nil {
		ctl.sendError(ctx, w, err)
		return
	}

	// send response
	ctl.sendResponse(ctx, w, http.StatusOK, trS, trP)
}
