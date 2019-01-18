package loader

import "etl-txt/etl"

// Loader interface contract
type Loader interface {
	Load(op *etl.Operation) error
}

// New loader by file type
func New(fileType string) Loader {
	switch fileType {
	case etl.TypeDebt:
		return &LDebt{}
	case etl.TypeClient:
		return &LClient{}
	}

	return nil
}
