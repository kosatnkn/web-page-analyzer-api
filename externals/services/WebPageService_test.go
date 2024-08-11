package services_test

import (
	"reflect"
	"testing"

	i "github.com/kosatnkn/web-page-analyzer-api/domain/boundary/services"
	"github.com/kosatnkn/web-page-analyzer-api/externals/services"
)

func TestNewWebPageServiceInterfaceImpl(t *testing.T) {
	// check
	var _ i.WebPageServiceInterface = &services.WebPageService{}
}

func TestNewWebPageServiceType(t *testing.T) {
	// run
	s := services.NewWebPageService()

	// check
	need := reflect.TypeOf(&services.WebPageService{})
	got := reflect.TypeOf(s)
	if got != need {
		t.Errorf("need '%v', got '%v'", need, got)
	}
}

func TestAnalyzeWithValidURL(t *testing.T) {
	// input
	svc := services.NewWebPageService()
	validURL := "http://example.com"
	cmp := []string{"title"}

	// run
	_, err := svc.Analyze(validURL, cmp)
	if err != nil {
		t.Errorf("need '%v', got '%v'", nil, err)
	}
}

func TestAnalyzeWithInvalidURL(t *testing.T) {
	// check panic
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected to panic but did not")
		}
	}()

	// input
	invalidURL := "http://kcvjkvhjkvbhxmcvjhxcvjhzxcbvjxcvjkhxcjbv.xkjvjkvjknvkjnadjknv"

	// run
	svc := services.NewWebPageService()
	cmp := []string{"title"}

	svc.Analyze(invalidURL, cmp)
}
