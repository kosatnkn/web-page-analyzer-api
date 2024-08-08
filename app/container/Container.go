package container

import (
	"githubcom/kosatnkn/web-page-analyzer-api/app/adapters"
	"githubcom/kosatnkn/web-page-analyzer-api/domain/boundary/services"
)

// Container holds all resolved dependencies that needs to be injected at run time.
type Container struct {
	Adapters Adapters
	Services Services
}

// Adapters hold resolved adapter instances.
//
// These are wrappers around third party libraries. All adapters will be of a corresponding adapter interface type.
type Adapters struct {
	Log       adapters.LogAdapterInterface
	Validator adapters.ValidatorAdapterInterface
}

// Services hold resolved service instances.
//
// These are abstractions to third party APIs. All services will be of a corresponding service interface type.
type Services struct {
	WebPageService services.WebPageServiceInterface
}
