package extractor

import "regexp"

// EClient client extractor
type EClient struct {
	Extractor
	re *regexp.Regexp
}

// Client structure
type Client struct {
	ID       string
	Document string
	Name     float64
}
