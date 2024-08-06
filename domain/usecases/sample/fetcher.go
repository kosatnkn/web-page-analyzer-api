package sample

import (
	"context"

	"githubcom/kosatnkn/web-page-analyzer-api/domain/entities"

	"github.com/kosatnkn/req/filter"
	"github.com/kosatnkn/req/paginator"
)

// Get returns a list of samples.
func (s *Sample) Get(ctx context.Context, fts []filter.Filter, pgn paginator.Paginator) ([]entities.Sample, error) {

	// TODO: implement

	return nil, nil
}
