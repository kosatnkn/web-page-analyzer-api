package sample

import (
	"githubcom/kosatnkn/web-page-analyzer-api/app/container"
)

// Sample contains all usecases for samples
type Sample struct{}

// NewSample creates a new instance of sample usecase.
func NewSample(ctr *container.Container) *Sample {
	return &Sample{}
}
