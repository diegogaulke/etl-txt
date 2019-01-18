package transformer

import (
	"etl-txt/etl"
)

// Transformer interface to manipulate data
type Transformer interface {
	Transform(op *etl.Operation) error
}

// New creates new transform
func New(fileType string) Transformer {
	switch fileType {
	case etl.TypeClient:
		return &TClient{}
	case etl.TypeDebt:
		return &TDebt{}
	}
	return nil
}
