package controllers

import (
	"net/http"

	"githubcom/kosatnkn/web-page-analyzer-api/app/container"
	"githubcom/kosatnkn/web-page-analyzer-api/domain/usecases/webpage"
	"githubcom/kosatnkn/web-page-analyzer-api/transport/http/request"
	"githubcom/kosatnkn/web-page-analyzer-api/transport/http/request/unpackers"
	"githubcom/kosatnkn/web-page-analyzer-api/transport/http/response/transformers"
)

// ReportController contains controller logic for endpoints.
type ReportController struct {
	*Controller
	webPageUseCase *webpage.WebPage
}

// NewReportController creates a new instance of the controller.
func NewReportController(c *container.Container) *ReportController {
	return &ReportController{
		Controller:     NewController(c),
		webPageUseCase: webpage.NewWebPage(c),
	}
}

// Get handles retrieving a list of samples.
func (ctl *ReportController) Get(w http.ResponseWriter, r *http.Request) {
	// add additional metadata
	ctx := ctl.withTrace(r.Context(), "ReportController.Get")

	// read data from the request
	url := ctl.urlParam(r, "url")
	if errs := ctl.validator.ValidateField("url", url, "required,url"); errs != nil {
		ctl.sendError(ctx, w, errs)
		return
	}

	cmp := ctl.urlParam(r, "cmp")
	if errs := ctl.validator.ValidateField("cmp", cmp, "required,json"); errs != nil {
		ctl.sendError(ctx, w, errs)
		return
	}
	u := unpackers.NewComponentsUnpacker()
	if err := request.Unpack([]byte(cmp), u); err != nil {
		ctl.sendError(ctx, w, err)
		return
	}

	// process
	report, err := ctl.webPageUseCase.Report(ctx, url, *u)
	if err != nil {
		ctl.sendError(ctx, w, err)
		return
	}

	// transform data for the response
	trR, err := ctl.transform(report, transformers.NewReportTransformer(), false)
	if err != nil {
		ctl.sendError(ctx, w, err)
		return
	}

	// send response
	ctl.sendResponse(ctx, w, http.StatusOK, trR)
}
