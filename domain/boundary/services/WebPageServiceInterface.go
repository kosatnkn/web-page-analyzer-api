package services

import "github.com/kosatnkn/web-page-analyzer-api/domain/entities"

// WebPageServiceInterface contract to manipulate web page retrieval.
type WebPageServiceInterface interface {
	// Analyze returns a report after analyzing the web page against provided criteria.
	Analyze(url string, components []string) (entities.Report, error)
}
