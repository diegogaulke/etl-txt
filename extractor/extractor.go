package extractor

import (
	"etl-txt/etl"
	"regexp"
	"strconv"
)

// Extractor interface contract
type Extractor interface {
	Extract(line string) (*etl.Operation, error)
}

// NewExtractor factory for extractors
func NewExtractor(fileType string) Extractor {
	switch fileType {
	case etl.TypeDebt:
		return &EDebt{re: regexp.MustCompile(etl.RegexDebt)}
	case etl.TypeClient:
		return &EClient{re: regexp.MustCompile(etl.RegexDebt)}
	}

	return nil
}

func formatValue(raw string) float64 {
	f, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return 0
	}
	return f / 100
}
