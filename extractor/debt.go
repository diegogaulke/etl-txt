package extractor

import (
	"etl-txt/etl"
	"regexp"
)

// EDebt debit extractor
type EDebt struct {
	Extractor
	re *regexp.Regexp
}

// Debt structure
type Debt struct {
	Key      string
	Document string
	Value    float64
	Creditor string
}

// Extract debt function
func (dx *EDebt) Extract(line string) (*etl.Operation, error) {
	match := dx.re.FindStringSubmatch(line)

	if len(match) > 0 {
		d := new(Debt)
		d.Document = match[2]

		op := new(etl.Operation)
		op.Op = match[1]
		op.Data = d

		return op, nil
	}

	return nil, nil
}
