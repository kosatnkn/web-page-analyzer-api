package entities

// Report entity
type Report struct {
	URL        string
	Version    string
	Title      string
	StatusCode int
	Components []Component
}
