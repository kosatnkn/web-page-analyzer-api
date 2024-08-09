package webpage

import (
	"github.com/kosatnkn/web-page-analyzer-api/app/container"
	"github.com/kosatnkn/web-page-analyzer-api/domain/boundary/services"
)

// WebPage contains all usecases for web page analyzing
type WebPage struct {
	webPageService services.WebPageServiceInterface
}

// NewWebPage creates a new instance of web page usecase.
func NewWebPage(ctr *container.Container) *WebPage {
	return &WebPage{
		webPageService: ctr.Services.WebPageService,
	}
}
