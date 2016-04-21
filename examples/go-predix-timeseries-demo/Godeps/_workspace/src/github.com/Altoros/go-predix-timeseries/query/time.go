package query

import (
	"fmt"
	"time"
)

type Time interface {
	Value() interface{}
}

type absTime struct {
	t time.Time
}

func (a *absTime) Value() interface{} {
	return a.t.UnixNano() / int64(time.Millisecond)
}

type relTime struct {
	t string
}

func (r *relTime) Value() interface{} {
	return r.t
}

type RelativeTime interface {
	MillisecondsAgo() Time
	SecondsAgo() Time
	MinutesAgo() Time
	HoursAgo() Time
	DaysAgo() Time
	WeeksAgo() Time
	MonthsAgo() Time
	YearsAgo() Time
}

type rel struct {
	value int
}

func (r *rel) MillisecondsAgo() Time { return &relTime{fmt.Sprintf("%dms-ago", r.value)} }
func (r *rel) SecondsAgo() Time      { return &relTime{fmt.Sprintf("%ds-ago", r.value)} }
func (r *rel) MinutesAgo() Time      { return &relTime{fmt.Sprintf("%dmi-ago", r.value)} }
func (r *rel) HoursAgo() Time        { return &relTime{fmt.Sprintf("%dh-ago", r.value)} }
func (r *rel) DaysAgo() Time         { return &relTime{fmt.Sprintf("%dd-ago", r.value)} }
func (r *rel) WeeksAgo() Time        { return &relTime{fmt.Sprintf("%dw-ago", r.value)} }
func (r *rel) MonthsAgo() Time       { return &relTime{fmt.Sprintf("%dmm-ago", r.value)} }
func (r *rel) YearsAgo() Time        { return &relTime{fmt.Sprintf("%dy-ago", r.value)} }

func R(value int) RelativeTime {
	return &rel{value}
}

func AbsTime(t time.Time) Time {
	return &absTime{t}
}
