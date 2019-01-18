package transform

import (
	"etl-txt/etl"
)

// TDebt the debt transform
type TDebt struct {
	Transformer
}

// Transform do the data manipulation
func (t *TDebt) Transform(op *etl.Operation) error {
	return nil
}
