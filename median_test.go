package goaugmented

import (
	"testing"
)


var (
	data = []Interval{ }
)

func init() {
	for i := 0; i < 100000; i ++ {
		data = append(data,&segment{int64(10-i),int64(10+i), 0, nil})
	}
}

//
func BenchmarkMedian(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			median(data)
		}
	})
}
