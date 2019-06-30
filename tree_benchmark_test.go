package goaugmented

import (
	"math/rand"
	"testing"

	gt "github.com/Kerah/goaugmented"
	gta "github.com/Kerah/goaugmented/augmented"
)

type record struct {
	A, B int
}

func (*record) Foo() {}

var (
	 countData = 1500000
	 base = 1 << 16
	 testData []Interval
)

//
func init()  {
	testData = make([]Interval, 0, countData)
	for i := 0; i < countData; i++ {
		from := base +  rand.Intn(base)
		to := from + rand.Intn(base)
		record := &record{from, to}

		interval := &segment{left: int64(from), right: int64(to), data: record}
		testData = append(testData, interval)
	}
}



//
func BenchmarkTree(b *testing.B) {
	tree := New()
	tree.Add(testData...)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {

		for pb.Next() {
			q := int64(base + rand.Intn(base))
			query := &segment{left: q-1, right: q+1}
			list := tree.Query(query)
			for _, item := range list {
				if r, ok := item.Data().(*record); ok {
					r.Foo()
				}
			}
		}
	})
}

//
func BenchmarkIntervalTree(b *testing.B) {
	tree := BuildITree(testData)
	println( "get_max_height: ", get_max_height(tree) )
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {

		for pb.Next() {
			q := int64(base + rand.Intn(base))
			query := &segment{left: q-1, right: q+1}
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
	m := map[uint64]*record{}
	tree := gt.New(1)
	for _, item := range testData {
		interval := gt.SingleDimensionInterval(
			gta.NewInt64(item.Low()),
			gta.NewInt64(item.High()),
			item.ID(),
		)
		m[item.ID()] = item.Data().(*record)
		tree.Add(interval)
	}

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {

		for pb.Next() {
			q := int64(base + rand.Intn(base))
			query := gt.SingleDimensionInterval(
				gta.NewInt64(q-1),
				gta.NewInt64(q+1),
				0,
			)
			list := tree.Query(query)

			for _, item := range list {
				if r, ok := m[item.ID()]; ok {
					r.Foo()
				}
			}
		}
	})
}
