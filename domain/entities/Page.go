package entities

// Page entity
type Page struct {
	URL        string
	Version    string
	Title      string
	StatusCode int
	Body       []byte
}
