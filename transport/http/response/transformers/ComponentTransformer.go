package transformers

import (
	"githubcom/kosatnkn/web-page-analyzer-api/domain/entities"
	"githubcom/kosatnkn/web-page-analyzer-api/transport/http/response/transformers/errors"
)

// ComponentTransformer is used to transform report data
type ComponentTransformer struct {
	Name    string                   `json:"name"`
	Count   uint32                   `json:"count"`
	Summary []map[string]interface{} `json:"summary,omitempty"`
}

// NewComponentTransformer creates a new instance of the transformer.
func NewComponentTransformer() TransformerInterface {
	return &ComponentTransformer{}
}

// TransformAsObject map data to a transformer object.
func (t *ComponentTransformer) TransformAsObject(data interface{}) (interface{}, error) {
	c, ok := data.(entities.Component)
	if !ok {
		return nil, t.dataMismatchError()
	}

	tr := ComponentTransformer{
		Name:    c.Name,
		Count:   c.Count,
		Summary: c.Summary,
	}

	return tr, nil
}

// TransformAsCollection map data to a collection of transformer objects.
func (t *ComponentTransformer) TransformAsCollection(data interface{}) (interface{}, error) {
	// Make sure that you declare the transformer slice in this manner.
	// Otherwise the marshaller will return `null` instead of `[]` when
	// marshalling empty slices
	// https://apoorvam.github.io/blog/2017/golang-json-marshal-slice-as-empty-array-not-null/
	trComponents := make([]ComponentTransformer, 0)

	components, ok := data.([]entities.Component)
	if !ok {
		return nil, t.dataMismatchError()
	}

	for _, component := range components {
		tr, err := t.TransformAsObject(component)
		if err != nil {
			return nil, err
		}

		trComponent := tr.(ComponentTransformer)
		trComponents = append(trComponents, trComponent)
	}

	return trComponents, nil
}

// dataMismatchError returns a data mismatch error of TransformerError type.
func (t *ComponentTransformer) dataMismatchError() error {
	return errors.NewTransformerError("100", "Cannot map given data to ComponentTransformer", nil)
}
