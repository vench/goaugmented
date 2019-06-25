package goaugmented

import (
	"testing"
)

type testPointer struct {
	Value string
}

//
func TestTree(t *testing.T) {
	tree := New(1)
	if tree.Len() != 0 {
		t.Fatalf(`Error compare size tree`)
	}

	iv1 := SingleDimensionInterval(
		NewInt64(50),
		NewInt64(100),
		10,
		&testPointer{"Some text"},
	)
	tree.Add(iv1)

	iv2 := SingleDimensionInterval(
		NewInt64(50),
		NewInt64(100),
		12,
		&testPointer{"Some text 2"},
	)
	tree.Add(iv2)

	iv3 := SingleDimensionInterval(
		NewInt64(100),
		NewInt64(200),
		15, &testPointer{"Some text 3"},
	)
	tree.Add(iv3)

	iv4 := SingleDimensionInterval(
		NewInt64(300),
		NewInt64(400),
		99, &testPointer{"Some text 4"},
	)
	tree.Add(iv4)

	iv5 := SingleDimensionInterval(
		NewInt64(300),
		NewInt64(400),
		99, &testPointer{"Some text 5"},
	)
	tree.Add(iv5) // not unique

	if tree.Len() != 4 {
		t.Fatalf(`Error compare size tree`)
	}

	query := ValueInterval(
		NewInt64(301),
	)

	intervals := tree.Query(query)
	if len(intervals) != 1 {
		t.Fatalf(`Error compare size query intervals`)
	}

	d, ok := intervals[0].Data().Data().(*testPointer)
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
}
