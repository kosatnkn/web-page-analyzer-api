package webpage

import (
	"context"
	"githubcom/kosatnkn/web-page-analyzer-api/domain/entities"
)

// Report returns a summarized report about the page.
func (wp *WebPage) Report(ctx context.Context, url string, components []string) (entities.Report, error) {
	p, err := wp.webPageService.Page(url, true)
	if err != nil {
		return entities.Report{}, err
	}

	return wp.analyze(ctx, p, components)
}
