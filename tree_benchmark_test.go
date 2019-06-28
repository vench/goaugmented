package goaugmented

import (
	"testing"
	"math/rand"

	gt "github.com/Kerah/goaugmented"
	gta "github.com/Kerah/goaugmented/augmented"
)

type record struct {
	A,B int
}


func (*record) Foo(){}


func BenchmarkTree(b *testing.B) {
	base := 100
	tree := New(1)
	for i := 0; i < 500000; i ++ {
		from := base + rand.Intn(base)
		to := base + rand.Intn(base)
		record := &record{from, to}

		interval := SingleDimensionInterval(
			NewInt64(int64(from)),
			NewInt64(int64(to)),
			uint64(i + 1),
			record,
		)
		tree.Add(interval)
	}


	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {

		for pb.Next() {
			query := ValueInterval(
				NewInt64(int64(base + rand.Intn(base))),
			)
			list := tree.Query(query)
			for _, item := range list {
				if r, ok := item.Data().(*record); ok {
					r.Foo()
				}
			}
		}
	})
}

func BenchmarkTreeOrigin(b *testing.B) {
	base := 100
	m := map[uint64]*record{}
	tree := gt.New(1)
	for i := 0; i < 500000; i ++ {
		from := base + rand.Intn(base)
		to := base + rand.Intn(base)
		id := uint64(i + 1)
		record := &record{from, to}
		m[id] = record

		interval := gt.SingleDimensionInterval(
			gta.NewInt64(int64(from)),
			gta.NewInt64(int64(to)),
			id,
		)
		tree.Add(interval)
	}


	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {

		for pb.Next() {
			query := gt.ValueInterval(
				gta.NewInt64(int64(base + rand.Intn(base))),
			)
			list := tree.Query(query)

			for _,item := range list {
				if r, ok := m[item.ID()]; ok {
					r.Foo()
				}
			}
		}
	})
}