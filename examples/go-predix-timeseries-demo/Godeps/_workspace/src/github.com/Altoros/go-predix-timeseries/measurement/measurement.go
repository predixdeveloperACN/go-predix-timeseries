package measurement

import (
	"fmt"
	"reflect"
)

type Int int
type Longint int64
type Float float32
type Double float64

type Measurement interface {
	Value() string
}

func (m Int) Value() string {
	return fmt.Sprintf("%d", m)
}

func (m Longint) Value() string {
	return fmt.Sprintf("%d", m)
}

func (m Float) Value() string {
	return fmt.Sprintf("%g", m)
}

func (m Double) Value() string {
	return fmt.Sprintf("%g", m)
}

type UnsupportedTypeError struct {
	Type reflect.Type
}

func (e *UnsupportedTypeError) Error() string {
	return fmt.Sprintf("unsupported measurement type: %s (only int, long int, float and double are supported", e.Type.String())
}
