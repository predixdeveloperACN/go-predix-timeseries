package measurement

import (
	"reflect"
	"strconv"
)

func M(measure interface{}) (Measurement, error) {
	switch measure.(type) {
	case int:
		return Int(measure.(int)), nil
	case int64:
		return Longint(measure.(int64)), nil
	case float32:
		return Float(measure.(float32)), nil
	case float64:
		return Double(measure.(float64)), nil
	}
	return nil, &UnsupportedTypeError{reflect.TypeOf(measure)}
}

func FromString(s string) (Measurement, error) {
	var intMeasure int64
	var floatMeasure float64

	intMeasure, err := strconv.ParseInt(s, 10, 32)
	if err == nil {
		return Int(intMeasure), nil
	}
	switch err.(*strconv.NumError).Err {
	case strconv.ErrRange:
		intMeasure, err = strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, err
		}
		return Longint(intMeasure), nil
	case strconv.ErrSyntax:
		if floatMeasure, err = strconv.ParseFloat(s, 32); err == nil {
			return Float(floatMeasure), nil
		}
		switch err.(*strconv.NumError).Err {
		case strconv.ErrRange:
			floatMeasure, err = strconv.ParseFloat(s, 64)
			if err != nil {
				return nil, err
			}
			return Double(floatMeasure), nil
		default:
			return nil, err
		}
	default:
		return nil, err
	}
}
