package measurement

import (
	"github.com/Altoros/go-predix-timeseries/filter"
)

func condition(condition string, values ...Measurement) filter.Filter {
	ret := make(map[string]interface{})
	ret["condition"] = condition
	vals := make([]string, len(values))
	for i, v := range values {
		vals[i] = v.Value()
	}
	ret["values"] = vals

	return filter.New("measurements", ret)
}

func Gt(values ...Measurement) filter.Filter { return condition("gt", values...) }
func Ge(values ...Measurement) filter.Filter { return condition("ge", values...) }
func Lt(values ...Measurement) filter.Filter { return condition("lt", values...) }
func Le(values ...Measurement) filter.Filter { return condition("le", values...) }
func Eq(values ...Measurement) filter.Filter { return condition("eq", values...) }
func Ne(values ...Measurement) filter.Filter { return condition("ne", values...) }
