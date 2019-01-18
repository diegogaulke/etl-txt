package loader

import (
	"encoding/json"
	"etl-txt/etl"
	"fmt"
)

// LDebt the debt loader
type LDebt struct {
	Loader
}

// Load debt
func (p *LDebt) Load(op *etl.Operation) error {
	o, _ := json.Marshal(op)
	fmt.Println(string(o))

	return nil
}
