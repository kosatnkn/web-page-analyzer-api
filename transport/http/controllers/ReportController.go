package controllers

import (
	"fmt"
	"net/http"

	"githubcom/kosatnkn/web-page-analyzer-api/app/container"
	"githubcom/kosatnkn/web-page-analyzer-api/domain/usecases/sample"
)

// ReportController contains controller logic for endpoints.
type ReportController struct {
	*Controller
	sampleUseCase *sample.Sample
}

// NewReportController creates a new instance of the controller.
func NewReportController(c *container.Container) *ReportController {
	return &ReportController{
		Controller:    NewController(c),
		sampleUseCase: sample.NewSample(c),
	}
}

// Get handles retrieving a list of samples.
func (ctl *ReportController) Get(w http.ResponseWriter, r *http.Request) {
	// add a trace string to the request context
	ctx := ctl.withTrace(r.Context(), "ReportController.Get")

	// FIXME: the validator field
	url := ctl.Controller.urlParam(r, "url")
	if errs := ctl.validator.ValidateField("url", url, "required,url"); errs != nil {
		ctl.sendError(ctx, w, errs)
		return
	}

	cmp := ctl.Controller.urlParam(r, "cmp")
	if errs := ctl.validator.ValidateField("cmp", cmp, "required,json"); errs != nil {
		ctl.sendError(ctx, w, errs)
		return
	}

	fmt.Println(url)
	fmt.Println(cmp)

	// // transform report data
	// trS, err := ctl.transform(samples, transformers.NewSampleTransformer(), true)
	// if err != nil {
	// 	ctl.sendError(ctx, w, err)
	// 	return
	// }

	// send response
	// ctl.sendResponse(ctx, w, http.StatusOK, trS)
	ctl.sendResponse(ctx, w, http.StatusOK)
}
