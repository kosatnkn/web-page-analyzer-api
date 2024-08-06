package sample

import (
	"githubcom/kosatnkn/web-page-analyzer-api/app/adapters"
	"githubcom/kosatnkn/web-page-analyzer-api/app/container"
	"githubcom/kosatnkn/web-page-analyzer-api/domain/boundary/repositories"
)

// Sample contains all usecases for samples
type Sample struct {
	db               adapters.DBAdapterInterface
	sampleRepository repositories.SampleRepositoryInterface
}

// NewSample creates a new instance of sample usecase.
func NewSample(ctr *container.Container) *Sample {
	return &Sample{
		db:               ctr.Adapters.DB,
		sampleRepository: ctr.Repositories.SampleRepository,
	}
}
