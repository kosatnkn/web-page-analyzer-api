package services_test

import (
	"fmt"
	"testing"

	"github.com/kosatnkn/web-page-analyzer-api/externals/services"
)

func TestAnalyze(t *testing.T) {
	// input
	svc := services.NewWebPageService()
	url := "http://example.com"
	cmp := []string{"h1", "h2", "h3"}

	// run
	_, err := svc.Analyze(url, cmp)
	if err != nil {
		t.Errorf("got %v", err)
	}

	// check
	fmt.Println("Done")
}
