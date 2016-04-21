package aggregation

const (
	//Default Aggregations
	Avg = "avg" //Returns the average of the values
	Max = "max" //Returns the most recent largest value
	Min = "min" //Returns the most recent smallest value
	Sum = "sum" // Returns the sum of all the values
)

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

func Param(name string, value interface{}) Parameter {
	return &param{name, value}
}
