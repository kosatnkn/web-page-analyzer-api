package controllers

import (
	"net/http"

	"githubcom/kosatnkn/web-page-analyzer-api/app/container"
	"githubcom/kosatnkn/web-page-analyzer-api/transport/http/response/transformers"
)

// APIController contains controller logic for endpoints.
type APIController struct {
	*Controller
}

// NewAPIController creates a new instance of the controller.
func NewAPIController(c *container.Container) *APIController {
	return &APIController{
		Controller: NewController(c),
	}
}

// GetInfo return basic details of the API.
func (ctl *APIController) GetInfo(w http.ResponseWriter, r *http.Request) {
	// transform
	tr := transformers.APITransformer{
		Name:    "Catalyst",
		Version: "v2.5.1",
		Purpose: "A REST API base written in Golang",
	}

	// send response
	ctl.sendResponse(r.Context(), w, http.StatusOK, tr)
}
