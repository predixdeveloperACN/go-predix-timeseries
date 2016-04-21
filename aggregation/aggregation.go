// Package aggregation provides primitives and a helper for specifying aggregation parameters in a query.
package aggregation

const (
	//Default aggregations
	Avg = "avg" // Returns the average of the values
	Max = "max" // Returns the most recent largest value
	Min = "min" // Returns the most recent smallest value
	Sum = "sum" // Returns the sum of all the values
)

// Represents the aggregation parameter
type Parameter interface {
	Name() string
	Value() interface{}
}

type param struct {
	name  string
	value interface{}
}

func (p *param) Name() string       { return p.name }
func (p *param) Value() interface{} { return p.value }

// Param creates the aggregation parameter.
func Param(name string, value interface{}) Parameter {
	return &param{name, value}
}
