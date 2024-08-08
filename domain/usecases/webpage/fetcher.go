package webpage

import (
	"context"

	"githubcom/kosatnkn/web-page-analyzer-api/domain/entities"
)

// Report returns a summarized report about the page.
func (wp *WebPage) Report(ctx context.Context, url string, components []string) (entities.Report, error) {
	return wp.mock(url), nil

	// p, err := wp.webPageService.Page(url, true)
	// if err != nil {
	// 	return entities.Report{}, err
	// }

	// return wp.analyze(ctx, p, components)
}

// analyze the page against the given set of components and produce a report.
func (wp *WebPage) analyze(ctx context.Context, page entities.Page, components []string) (entities.Report, error) {
	report := entities.Report{
		URL:        page.URL,
		StatusCode: page.StatusCode,
	}

	report.Version = "html 5"   // wp.httpVersion(&p)
	report.Title = "Page Title" // wp.title(&p)
	// report.Components = []entities.Component{} // wp.components(&p, components)

	return report, nil
}

func (wp *WebPage) mock(url string) entities.Report {
	r := entities.Report{
		URL:     url,
		Version: "1.1",
		Title:   "Test Page",
	}

	var extra []map[string]interface{}
	ex := make(map[string]interface{})
	ex["f1"] = "v1"
	extra = append(extra, ex)

	c := entities.Component{Name: "h1", Count: 1, Summary: extra}
	r.Components = append(r.Components, c)

	return r
}
