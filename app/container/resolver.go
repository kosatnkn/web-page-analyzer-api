package container

import (
	"github.com/kosatnkn/web-page-analyzer-api/app/config"
)

// Resolve resolves the entire container.
//
// The order of resolution is very important. Low level dependencies need to be resolved before high level dependencies.
// It generally happens in this order.
//   - Adapters
//   - Repositories
//   - Services
func Resolve(cfg *config.Config) *Container {
	ctr := Container{}
	ctr.Adapters = resolveAdapters(cfg)
	ctr.Services = resolveServices(cfg.Services)

	return &ctr
}
