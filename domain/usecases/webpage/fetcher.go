package webpage

import (
	"context"
	"githubcom/kosatnkn/web-page-analyzer-api/domain/entities"
)

// Report returns a summarized report about the page.
func (wp *WebPage) Report(ctx context.Context, url string, components []string) (entities.Report, error) {
	return wp.webPageService.Analyze(url, components)
}
