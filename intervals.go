/*
Copyright 2014 Workiva, LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package goaugmented

import (
	"sync"
)

var intervalsPool = sync.Pool{
	New: func() interface{} {
		return make(Intervals, 0, 10)
	},
}

// Intervals represents a list of Intervals.
type Intervals []Interval

// Dispose will free any consumed resources and allow this list to be
// re-allocated.
func (ivs *Intervals) Dispose() {
	for i := 0; i < len(*ivs); i++ {
		(*ivs)[i] = nil
	}

	*ivs = (*ivs)[:0]
	intervalsPool.Put(*ivs)
}

type dimension struct {
	low, high int64
}

type interval struct {
	dimension *dimension
	id       uint64
	data  interface{}
}

func (mi *interval) LowAtDimension() int64 {
	return mi.dimension.low
}

func (mi *interval) HighAtDimension() int64 {
	return mi.dimension.high
}

func (mi *interval) OverlapsAtDimension(iv Interval) bool {
	return mi.HighAtDimension() > iv.LowAtDimension() &&
		mi.LowAtDimension() < iv.HighAtDimension()
}

func (mi interval) Data() interface{} {
	return mi.data
}

func (mi interval) ID() uint64 {
	return mi.id
}

func SingleDimensionInterval(low, high int64, id uint64, data interface{}) *interval {
	return &interval{&dimension{low: low, high: high}, id, data}
}


func ValueInterval(val int64) *interval {
	return SingleDimensionInterval(val, val, 0, nil)
}
