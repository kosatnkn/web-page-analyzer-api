package unpackers

// ComponentsUnpacker contains the unpacking structure for component list in request url param 'cmp'.
type ComponentsUnpacker []string

// NewComponentsUnpacker creates a new instance of the unpacker.
func NewComponentsUnpacker() *ComponentsUnpacker {
	return &ComponentsUnpacker{}
}

// RequiredFormat returns the applicable JSON format for the address data structure.
func (u *ComponentsUnpacker) RequiredFormat() string {
	return `["h1","a","login"]`
}
