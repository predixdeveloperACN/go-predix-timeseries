package query

import (
	"github.com/Altoros/go-predix-timeseries/aggregation"
	"github.com/Altoros/go-predix-timeseries/filter"
)

type QueryTag interface {
	InvertOrder() QueryTag
	Limit(limit int) QueryTag
	Filter(f filter.Filter) QueryTag
	Aggregation(atype string, opts ...aggregation.Parameter) QueryTag
	GroupByAttributes(attributes ...string) QueryTag
	GroupByQuality() QueryTag
	GroupByMeasurement(rangeSize int) QueryTag
	GroupByTime(rangeSize string, groupCount int) QueryTag
}

type tag struct {
	Name         []string               `json:"name"`
	Order        string                 `json:"order,omitempty"`
	Limit_       int                    `json:"limit,omitempty"`
	Aggregations []interface{}          `json:"aggregations,omitempty"`
	Filters      map[string]interface{} `json:"filters,omitempty"`
	Groups       []interface{}          `json:"groups,omitempty"`
}

func Tag(name ...string) QueryTag {
	return &tag{Name: name}
}

func (t *tag) InvertOrder() QueryTag {
	if t.Order == "" {
		t.Order = "desc"
	} else if t.Order == "desc" {
		t.Order = ""
	}
	return t
}

func (t *tag) Limit(limit int) QueryTag {
	t.Limit_ = limit
	return t
}

func (t *tag) Filter(f filter.Filter) QueryTag {
	if t.Filters == nil {
		t.Filters = make(map[string]interface{})
	}
	t.Filters[f.Type()] = f.Attributes()
	return t
}

func (t *tag) Aggregation(atype string, params ...aggregation.Parameter) QueryTag {
	aggregation := make(map[string]interface{})
	aggregation["type"] = atype
	for _, param := range params {
		aggregation[param.Name()] = param.Value()
	}
	t.Aggregations = append(t.Aggregations, aggregation)
	return t
}

func (t *tag) GroupByAttributes(attributes ...string) QueryTag {
	t.Groups = append(t.Groups, map[string]interface{}{"name": "attribute",
		"attributes": attributes,
	})
	return t
}

func (t *tag) GroupByQuality() QueryTag {
	t.Groups = append(t.Groups, map[string]string{"name": "quality"})
	return t
}

func (t *tag) GroupByMeasurement(rangeSize int) QueryTag {
	t.Groups = append(t.Groups, map[string]interface{}{"name": "measurement", "rangeSize": rangeSize})
	return t
}

func (t *tag) GroupByTime(rangeSize string, groupCount int) QueryTag {
	t.Groups = append(t.Groups, map[string]interface{}{"name": "time",
		"rangeSize":  rangeSize,
		"groupCount": groupCount})
	return t
}
