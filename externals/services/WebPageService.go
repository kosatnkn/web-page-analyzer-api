package services

import (
	"fmt"
	"githubcom/kosatnkn/web-page-analyzer-api/domain/boundary/services"
	"githubcom/kosatnkn/web-page-analyzer-api/domain/entities"
	"githubcom/kosatnkn/web-page-analyzer-api/externals/services/errors"
	"io"
	"net/http"

	"golang.org/x/net/html"
)

// WebPageService retrieves and manipulate web pages.
type WebPageService struct{}

// NewWebPageService creates a new instance of the service.
func NewWebPageService() services.WebPageServiceInterface {
	return &WebPageService{}
}

// Page retrieves the web page.
func (svc *WebPageService) Page(url string, withBody bool) (entities.Page, error) {
	// retrieve page
	res, err := http.Get(url)
	if err != nil {
		return entities.Page{}, svc.errorPageNotFound(err, res.StatusCode)
	}
	defer res.Body.Close()

	p := entities.Page{
		URL:        url,
		StatusCode: res.StatusCode,
	}
	if withBody {
		b, err := io.ReadAll(res.Body)
		if err != nil {
			return p, svc.errorReadingBody(err)
		}
		p.Body = b
	}

	return p, nil
}

// Analyze returns a report after analyzing the web page against provided criteria.
func (svc *WebPageService) Analyze(url string, components []string) (entities.Report, error) {
	// retrieve page
	res, err := http.Get(url)
	if err != nil {
		return entities.Report{}, svc.errorPageNotFound(err, res.StatusCode)
	}
	defer res.Body.Close()

	return svc.analyze(res, components)
}

func (svc *WebPageService) analyze(res *http.Response, components []string) (entities.Report, error) {
	r := entities.Report{
		URL:        res.Request.URL.String(),
		StatusCode: res.StatusCode,
	}

	// Create an HTML tokenizer
	z := html.NewTokenizer(res.Body)

	// Loop through HTML tokens
	for {
		tokenType := z.Next()

		switch tokenType {
		case html.StartTagToken, html.SelfClosingTagToken:
			token := z.Token()

			if token.Data == "li" {
				svc.processProductDetails(z)
				// Exit the loop after processing the details
				//return
			}
		}

		if tokenType == html.ErrorToken {
			break
		}
	}

	return r, nil
}

func (svc *WebPageService) processProductDetails(z *html.Tokenizer) {
	fmt.Println(z)
}

// --- errors ---

func (svc *WebPageService) errorPageNotFound(cause error, statusCode int) error {
	return errors.NewServiceError("101", fmt.Sprintf("webpage-service: %d, error retrieving page", statusCode), cause)
}

func (svc *WebPageService) errorReadingBody(cause error) error {
	return errors.NewServiceError("102", "webpage-service: error reading page body", cause)
}
