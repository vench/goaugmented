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

type interval struct {
	low, high int64
	id        uint64
	data      interface{}
}

func (mi *interval) Low() int64 {
	return mi.low
}

func (mi *interval) High() int64 {
	return mi.high
}

func (mi *interval) Overlaps(iv Interval) bool {
	return mi.High() > iv.Low() &&
		mi.Low() < iv.High()
}

func (mi interval) Data() interface{} {
	return mi.data
}

func (mi interval) ID() uint64 {
	return mi.id
}

func SingleInterval(low, high int64, id uint64, data interface{}) *interval {
	return &interval{low, high, id, data}
}

func ValueInterval(val int64) *interval {
	return SingleInterval(val, val, 0, nil)
}
