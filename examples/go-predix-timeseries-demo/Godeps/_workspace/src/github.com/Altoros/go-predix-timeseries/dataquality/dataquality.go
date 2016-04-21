package dataquality

type Quality int

const (
	Bad           Quality = iota
	Uncertain             = iota
	NotApplicable         = iota
	Good                  = iota
)
