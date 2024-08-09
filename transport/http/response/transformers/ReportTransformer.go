package transformers

import (
	"github.com/kosatnkn/web-page-analyzer-api/domain/entities"
	"github.com/kosatnkn/web-page-analyzer-api/transport/http/response/transformers/errors"
)

// ReportTransformer is used to transform report data
type ReportTransformer struct {
	URL        string                 `json:"url"`
	Version    string                 `json:"version"`
	Title      string                 `json:"title"`
	Components []ComponentTransformer `json:"components,omitempty"`
}

// NewReportTransformer creates a new instance of the transformer.
func NewReportTransformer() TransformerInterface {
	return &ReportTransformer{}
}

// TransformAsObject map data to a transformer object.
func (t *ReportTransformer) TransformAsObject(data interface{}) (interface{}, error) {
	report, ok := data.(entities.Report)
	if !ok {
		return nil, t.dataMismatchError()
	}

	tr := ReportTransformer{
		URL:     report.URL,
		Version: report.Version,
		Title:   report.Title,
	}

	cs, err := NewComponentTransformer().TransformAsCollection(report.Components)
	if err != nil {
		return nil, err
	}
	tr.Components = cs.([]ComponentTransformer)

	return tr, nil
}

// TransformAsCollection map data to a collection of transformer objects.
func (t *ReportTransformer) TransformAsCollection(data interface{}) (interface{}, error) {
	// Make sure that you declare the transformer slice in this manner.
	// Otherwise the marshaller will return `null` instead of `[]` when
	// marshalling empty slices
	// https://apoorvam.github.io/blog/2017/golang-json-marshal-slice-as-empty-array-not-null/
	trReports := make([]ReportTransformer, 0)

	reports, ok := data.([]entities.Report)
	if !ok {
		return nil, t.dataMismatchError()
	}

	for _, report := range reports {
		tr, err := t.TransformAsObject(report)
		if err != nil {
			return nil, err
		}

		trReport := tr.(ReportTransformer)
		trReports = append(trReports, trReport)
	}

	return trReports, nil
}

// dataMismatchError returns a data mismatch error of TransformerError type.
func (t *ReportTransformer) dataMismatchError() error {
	return errors.NewTransformerError("100", "Cannot map given data to ReportTransformer", nil)
}
