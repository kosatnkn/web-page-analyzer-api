package webpage

import (
	"context"

	"githubcom/kosatnkn/web-page-analyzer-api/domain/entities"
)

// Get returns
func (s *WebPage) Report(ctx context.Context, url string, components []string) (entities.Report, error) {
	// mock report
	r := entities.Report{
		Version: "1.1",
		Title:   "Test Page",
	}

	var extra []map[string]interface{}
	ex := make(map[string]interface{})
	ex["f1"] = "v1"
	extra = append(extra, ex)

	c := entities.Component{Name: "h1", Count: 1, Extra: extra}
	r.Components = append(r.Components, c)

	return r, nil
}
