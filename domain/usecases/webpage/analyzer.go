package webpage

import (
	"context"
	"githubcom/kosatnkn/web-page-analyzer-api/domain/entities"
)

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
