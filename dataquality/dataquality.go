// Package dataquality contains data quality constants and helper function to create
// a filter by data quality in queries.
package dataquality

type Quality int

const (
	Bad           Quality = iota
	Uncertain             = iota
	NotApplicable         = iota
	Good                  = iota
)
