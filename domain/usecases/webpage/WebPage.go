package webpage

import (
	"githubcom/kosatnkn/web-page-analyzer-api/app/container"
)

// WebPage contains all usecases for web page analyzing
type WebPage struct {
	URL string
}

// NewWebPage creates a new instance of web page usecase.
func NewWebPage(ctr *container.Container) *WebPage {
	return &WebPage{}
}
