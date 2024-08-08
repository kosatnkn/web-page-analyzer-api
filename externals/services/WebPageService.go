package services

import (
	"fmt"
	"githubcom/kosatnkn/web-page-analyzer-api/domain/boundary/services"
	"githubcom/kosatnkn/web-page-analyzer-api/domain/entities"
	"githubcom/kosatnkn/web-page-analyzer-api/externals/services/errors"
	"io"
	"net/http"
)

// WebPageService retrieves and manipulate web pages.
type WebPageService struct{}

// NewWebPageService creates a new instance of the service.
func NewWebPageService() services.WebPageServiceInterface {
	return &WebPageService{}
}

// Page retrieves the web page
func (svc *WebPageService) Page(url string, withBody bool) (entities.Page, error) {
	p := entities.Page{URL: url}

	res, err := http.Get(url)
	p.StatusCode = res.StatusCode
	if err != nil {
		return p, errors.NewServiceError("101", fmt.Sprintf("webpage-service: %d, error retrieving page", res.StatusCode), err)
	}

	if withBody {
		defer res.Body.Close()
		b, err := io.ReadAll(res.Body)
		if err != nil {
			return p, errors.NewServiceError("102", "webpage-service: error reading page body", err)
		}
		p.Body = b
	}

	return p, nil
}
