package container

import (
	"fmt"

	"githubcom/kosatnkn/web-page-analyzer-api/app/config"
	"githubcom/kosatnkn/web-page-analyzer-api/externals/services"
)

// resolveServices resolves all services.
func resolveServices(cfgs []config.ServiceConfig) Services {
	svs := Services{}
	svs.WebPageService = services.NewWebPageService()

	return svs
}

// getServiceConfig returns the service config by the name of the service.
//
// Will panic if there is no service config found for a given service name.
func getServiceConfig(cfgs []config.ServiceConfig, name string) config.ServiceConfig {
	for i := range cfgs {
		if cfgs[i].Name == name {
			return cfgs[i]
		}
	}

	panic(fmt.Sprintf("Cannot find service configurations for `%s` service", name))
}
