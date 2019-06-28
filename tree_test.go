package goaugmented

import (
	"testing"
)

type testPointer struct {
	Value string
}

//
func TestTree(t *testing.T) {
	tree := New()
	if tree.Len() != 0 {
		t.Fatalf(`Error compare size tree`)
	}

	iv1 := SingleDimensionInterval(
		50,
		100,
		10,
		&testPointer{"Some text"},
	)
	tree.Add(iv1)

	iv2 := SingleDimensionInterval(
		50,
		100,
		12,
		&testPointer{"Some text 2"},
	)
	tree.Add(iv2)

	iv3 := SingleDimensionInterval(
		100,
		200,
		15, &testPointer{"Some text 3"},
	)
	tree.Add(iv3)

	iv4 := SingleDimensionInterval(
		300,
		400,
		99, &testPointer{"Some text 4"},
	)
	tree.Add(iv4)

	iv5 := SingleDimensionInterval(
		300,
		400,
		99, &testPointer{"Some text 5"},
	)
	tree.Add(iv5) // not unique

	if tree.Len() != 4 {
		t.Fatalf(`Error compare size tree`)
	}

	query := ValueInterval(301	)

	intervals := tree.Query(query)
	if len(intervals) != 1 {
		t.Fatalf(`Error compare size query intervals`)
	}

	d, ok := intervals[0].Data().(*testPointer)
	if !ok {
		t.Fatalf(`Error cast data to pointer`)
	}

	if d.Value != `Some text 4` {
		t.Fatalf(`Error compare data value`)
	}

	tree.Delete(iv5)
	if tree.Len() != 3 {
		t.Fatalf(`Error compare size tree`)
	}

	tree.Delete(iv5)
	if tree.Len() != 3 {
		t.Fatalf(`Error compare size tree`)
	}

	intervals = tree.Query(query)
	if len(intervals) != 0 {
		t.Fatalf(`Error compare size query intervals`)
	}

	intervals = tree.Query(ValueInterval(69	))
	if len(intervals) != 2 {
		t.Fatalf(`Error compare size query intervals`)
	}

	if intervals[0].ID() != iv1.ID() {
		t.Fatalf(`Error compare ids`)
	}

	if intervals[1].ID() != iv2.ID() {
		t.Fatalf(`Error compare ids`)
	}

	if intervals[1].ID() == intervals[0].ID() {
		t.Fatalf(`Error not compare ids`)
	}

	tree.Delete(iv2)

	intervals = tree.Query(ValueInterval(76	))
	if len(intervals) != 1 {
		t.Fatalf(`Error compare size query intervals`)
	}
	if intervals[0].ID() != iv1.ID() {
		t.Fatalf(`Error compare ids`)
	}
}
