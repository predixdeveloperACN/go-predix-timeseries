package dataquality

import (
	"strconv"

	"github.com/Altoros/go-predix-timeseries/filter"
)

func Qualities(values ...Quality) filter.Filter {
	ret := make(map[string]interface{})
	vals := make([]string, len(values))
	for i, v := range values {
		vals[i] = strconv.Itoa(int(v))
	}
	ret["values"] = vals
	return filter.New("qualities", ret)
}
