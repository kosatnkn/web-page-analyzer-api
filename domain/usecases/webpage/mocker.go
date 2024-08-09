package webpage

import "github.com/kosatnkn/web-page-analyzer-api/domain/entities"

func (wp *WebPage) _mock(url string) entities.Report {
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
