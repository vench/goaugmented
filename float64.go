package goaugmented

import (
	"math"
)

type augmentedFloat64 float64

func NewFloat64(in float64) Value {
	var v augmentedFloat64 = augmentedFloat64(in)
	return &v
}

func NewFloat64Interval(start, end float64, id uint64, data interface{}) Interval {
	return SingleDimensionInterval(
		NewFloat64(start),
		NewFloat64(end),
		id,
		data,
	)
}

func NewFloat64VI(value float64) Interval {
	return ValueInterval(NewFloat64(value))
}

func (af *augmentedFloat64) typeCast(in Value) float64 {
	switch in.(type) {
	case *augmentedFloat64:
		data := in.(*augmentedFloat64)
		return float64(*data)
	}
	return 0
}

func (af *augmentedFloat64) Greater(in Value) (r bool) {
	if float64(*af) > af.typeCast(in) {
		r = true
	}
	return
}

func (af *augmentedFloat64) GreaterOrEq(in Value) (r bool) {
	if float64(*af) >= af.typeCast(in) {
		r = true
	}
	return
}
func (af *augmentedFloat64) Lesser(in Value) (r bool) {
	if float64(*af) < af.typeCast(in) {
		r = true
	}
	return
}
func (af *augmentedFloat64) LesserOrEq(in Value) (r bool) {
	if float64(*af) <= af.typeCast(in) {
		r = true
	}
	return
}
func (af *augmentedFloat64) Substract(in Value) (r int64) {

	return int64(float64(*af) - af.typeCast(in))
}

func (af *augmentedFloat64) Add(in int64) Value {
	*af = augmentedFloat64(float64(*af) + float64(in))
	return af
}

func (af *augmentedFloat64) MinimalValue() Value {
	df := augmentedFloat64(-1 * math.MaxFloat64)
	return &df
}
func (af *augmentedFloat64) MaximalValue() Value {
	df := augmentedFloat64(math.MaxFloat64)
	return &df
}
