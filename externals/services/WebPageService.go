package services

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/kosatnkn/web-page-analyzer-api/domain/boundary/services"
	"github.com/kosatnkn/web-page-analyzer-api/domain/entities"
	"github.com/kosatnkn/web-page-analyzer-api/externals/services/errors"

	"golang.org/x/net/html"
)

// WebPageService retrieves and manipulate web pages.
type WebPageService struct{}

// NewWebPageService creates a new instance of the service.
func NewWebPageService() services.WebPageServiceInterface {
	return &WebPageService{}
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

// analyze does all the heavy lifting of the analysis.
func (svc *WebPageService) analyze(res *http.Response, components []string) (entities.Report, error) {
	r := entities.Report{
		URL:        res.Request.URL.String(),
		StatusCode: res.StatusCode,
	}
	counter := svc.initCounterMap(components)
	var aSummary []map[string]interface{}

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
				break
			}
			if t.Data == "a" {
				aSummary = append(aSummary, svc.aExtras(t))
			}
		}

		// when in error break out of the loop
		if tokenType == html.ErrorToken {
			break
		}
	}

	for k, v := range counter {
		if v != 0 {
			c := entities.Component{Name: k, Count: v}
			if k == "a" {
				c.Summary = aSummary
			}
			r.Components = append(r.Components, c)
		}
	}

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

// aExtras returns additional info on an <a> tag.
func (svc *WebPageService) aExtras(t html.Token) map[string]interface{} {
	m := make(map[string]interface{})
	for _, a := range t.Attr {
		if a.Key == "href" {
			m["href"] = a.Val
			m["external"] = svc.isExternalLink(a.Val)
		}
	}

	return m
}

// isExternalLink checks whether the link is an external url.
func (svc *WebPageService) isExternalLink(link string) bool {
	return strings.HasPrefix(link, "http")
}

// toHTMLVersion infer the HTML version from the DOCTYPE data.
func (svc *WebPageService) toHTMLVersion(data string) string {
	data = strings.ToLower(data)
	if data == "html" {
		return "html 5"
	}

	return data
}

func (svc *WebPageService) errorPageNotFound(cause error, statusCode int) error {
	return errors.NewServiceError("101", fmt.Sprintf("webpage-service: %d, error retrieving page", statusCode), cause)
}
