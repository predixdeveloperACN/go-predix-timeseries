package aggregation

import (
	"fmt"
)

type IntervalParam interface {
	Millisecond() Parameter
	Second() Parameter
	Minute() Parameter
	Hour() Parameter
	Day() Parameter
	Week() Parameter
	Month() Parameter
	Year() Parameter
}

type interval struct {
	value int
}

func Interval(value int) IntervalParam {
	return &interval{value: value}
}

func (i *interval) Millisecond() Parameter { return Param("interval", fmt.Sprintf("%dms", i.value)) }
func (i *interval) Second() Parameter      { return Param("interval", fmt.Sprintf("%ds", i.value)) }
func (i *interval) Minute() Parameter      { return Param("interval", fmt.Sprintf("%dmi", i.value)) }
func (i *interval) Hour() Parameter        { return Param("interval", fmt.Sprintf("%dh", i.value)) }
func (i *interval) Day() Parameter         { return Param("interval", fmt.Sprintf("%dd", i.value)) }
func (i *interval) Week() Parameter        { return Param("interval", fmt.Sprintf("%dw", i.value)) }
func (i *interval) Month() Parameter       { return Param("interval", fmt.Sprintf("%dmm", i.value)) }
func (i *interval) Year() Parameter        { return Param("interval", fmt.Sprintf("%dy", i.value)) }
