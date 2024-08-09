package services

import (
	"fmt"
	"githubcom/kosatnkn/web-page-analyzer-api/domain/boundary/services"
	"githubcom/kosatnkn/web-page-analyzer-api/domain/entities"
	"githubcom/kosatnkn/web-page-analyzer-api/externals/services/errors"
	"io"
	"net/http"
	"strings"

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
	counter := svc.initCounterMap(components)

	// create an HTML tokenizer
	z := html.NewTokenizer(res.Body)
	for {
		// this will move the pointer forward
		tokenType := z.Next()

		switch tokenType {
		case html.DoctypeToken:
			r.Version = svc.toHTMLVersion(z.Token().Data)
		case html.StartTagToken, html.SelfClosingTagToken:
			t := z.Token()
			svc.incrCounterMap(counter, t.Data)
			if t.Data == "title" {
				// move pointer forward one position because
				// the immediate token after the <title> token is the text token with the value of the title
				z.Next()
				r.Title = z.Token().Data
			}
		}

		// when in error break out of the loop
		if tokenType == html.ErrorToken {
			break
		}
	}

	// TODO: trace
	fmt.Println(r)
	fmt.Println(counter)

	return r, nil
}

// initCounterMap creates a map of components as keys initialized to 0
// while also removing duplicate components.
func (svc *WebPageService) initCounterMap(components []string) map[string]uint32 {
	m := make(map[string]uint32)
	for _, c := range components {
		m[c] = 0
	}

	return m
}

// incrCounterMap increments the count of the component in the map by one if it is there.
func (svc *WebPageService) incrCounterMap(counter map[string]uint32, component string) {
	if v, ok := counter[component]; ok {
		counter[component] = v + 1
	}
}

// toHTMLVersion infer the HTML version from the DOCTYPE data.
// ref: https://www.w3.org/QA/2002/04/valid-dtd-list.html
func (svc *WebPageService) toHTMLVersion(data string) string {
	data = strings.ToLower(data)

	switch data {
	case `html`:
		return `html5`
	case `html public "-//w3c//dtd html 4.01//en" "http://www.w3.org/tr/html4/strict.dtd`:
		return `html 4.01 strict`
	case `html public "-//w3c//dtd html 4.01 transitional//en" "http://www.w3.org/tr/html4/loose.dtd`:
		return `html 4.01 transitional`
	case `html public "-//w3c//dtd html 4.01 frameset//en" "http://www.w3.org/tr/html4/frameset.dtd`:
		return `html 4.01 frameset`
	case `html public "-//w3c//dtd xhtml 1.0 strict//en" "http://www.w3.org/tr/xhtml1/dtd/xhtml1-strict.dtd`:
		return `xhtml 1.0 strict`
	case `html public "-//w3c//dtd xhtml 1.0 transitional//en" "http://www.w3.org/tr/xhtml1/dtd/xhtml1-transitional.dtd`:
		return `xhtml 1.0 transitional`
	case `html public "-//w3c//dtd xhtml 1.0 frameset//en" "http://www.w3.org/tr/xhtml1/dtd/xhtml1-frameset.dtd`:
		return `xhtml 1.0 frameset`
	case `html public "-//w3c//dtd xhtml 1.1//en" "http://www.w3.org/tr/xhtml11/dtd/xhtml11.dtd`:
		return `xhtml 1.1`
	case `html public "-//w3c//dtd xhtml basic 1.1//en" "http://www.w3.org/tr/xhtml-basic/xhtml-basic11.dtd`:
		return `xhtml basic 1.1`
	case `math public "-//w3c//dtd mathml 2.0//en" "http://www.w3.org/math/dtd/mathml2/mathml2.dtd`:
		return `mathml 2.0`
	case `math system "http://www.w3.org/math/dtd/mathml1/mathml.dtd`:
		return `mathml 1.01`
	default:
		return `unknown`
	}
}

func (svc *WebPageService) errorPageNotFound(cause error, statusCode int) error {
	return errors.NewServiceError("101", fmt.Sprintf("webpage-service: %d, error retrieving page", statusCode), cause)
}

func (svc *WebPageService) errorReadingBody(cause error) error {
	return errors.NewServiceError("102", "webpage-service: error reading page body", cause)
}
