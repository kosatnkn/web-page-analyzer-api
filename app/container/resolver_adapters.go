package container

import (
	"fmt"

	"github.com/kosatnkn/web-page-analyzer-api/app/adapters"
	"github.com/kosatnkn/web-page-analyzer-api/app/config"

	"github.com/kosatnkn/log"
	"github.com/kosatnkn/validator"
)

// resolveAdapters resolves all adapters.
func resolveAdapters(cfg *config.Config) Adapters {
	ats := Adapters{}
	ats.Log = resolveLogAdapter(cfg.Log)
	ats.Validator = resolveValidatorAdapter()

	return ats
}

// resolveLogAdapter resolves the logging adapter.
func resolveLogAdapter(cfg log.Config) adapters.LogAdapterInterface {
	la, err := log.NewAdapter(cfg)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	return la
}

// resolveValidatorAdapter resolves the validation adapter.
func resolveValidatorAdapter() adapters.ValidatorAdapterInterface {
	v, err := validator.NewAdapter()
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	return v
}
