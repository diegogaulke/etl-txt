// Package etl contains the common structures and interfaces of the etl, to exchange data between extractor and processor
package etl

// types constants
const (
	TypeDebt   = "debt"
	TypeClient = "client"
)

// regex constants
const (
	RegexDebt   = "(.{1})(.{11})(.{10})(.{14})"
	RegexClient = "(.{1})(.{11})(.{10})(.{14})"
)

// Operation structure
type Operation struct {
	// I insert, R remove
	Op   string
	Data interface{}
}
