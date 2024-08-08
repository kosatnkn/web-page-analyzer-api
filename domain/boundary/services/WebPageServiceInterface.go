package services

import "githubcom/kosatnkn/web-page-analyzer-api/domain/entities"

// WebPageServiceInterface contract to manipulate web page retrieval.
type WebPageServiceInterface interface {
	// Page retrieves the web page that the url points to.
	Page(url string, withBody bool) (entities.Page, error)
}
