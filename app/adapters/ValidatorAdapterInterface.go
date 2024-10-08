package adapters

// ValidatorAdapterInterface is implemented by all validator adapters.
type ValidatorAdapterInterface interface {
	// Validate validates fields of a struct.
	Validate(data interface{}) map[string]string

	// ValidateField validates a single variable.
	ValidateField(name string, value interface{}, rules string) map[string]string
}
