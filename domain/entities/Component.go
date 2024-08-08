package entities

// Component entity
type Component struct {
	Name  string
	Count uint32
	Extra []map[string]interface{}
}
