package transformer

import (
	"etl-txt/etl"
)

// TClient the client transform
type TClient struct {
	Transformer
}

// Transform do the data manipulation
func (t *TClient) Transform(op *etl.Operation) error {
	// transform Data here as you wish
	return nil
}
