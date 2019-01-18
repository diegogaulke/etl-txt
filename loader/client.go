package loader

import "etl-txt/etl"

// LClient client loader
type LClient struct {
	Loader
}

// Load the loader
func (p *LClient) Load(op *etl.Operation) error {
	return nil
}
