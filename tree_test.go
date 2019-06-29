package goaugmented

import (
	"testing"
)

type testPointer struct {
	Value string
}


func TestTreeAlter(t *testing.T) {
	ss := []*Segment{
		{78,98, nil},{6,8, nil},{5,9, nil},{11,20, nil},{3, 10, nil},
		{20, 21, nil}, {1,12, nil}, {5,8, nil},  {5,14, nil},
	}
	root := build_tree(ss)
	inorder(root)
	x := &Segment{12,13, nil}
	res := get_ans(root, x)
	//for _, rr := range res {
	//	println("[",rr.left, "=", rr.right,"]")
	//}
	if len(res) != 2 {
		t.Fatalf(`Error compare size query intervals`)
	}
	if res[0].left != 11 || res[0].right != 20 {
		t.Fatalf(`Error find wrong element`)
	}
	if res[1].left != 5 || res[1].right != 14 {
		t.Fatalf(`Error find wrong element`)
	}

	x.right = 100
	res = get_ans(root, x)
	if len(res) != 0 {
		t.Fatalf(`Error compare size query intervals`)
	}

	x.left = -10
	x.right = 13
	if len(res) != 0 {
		t.Fatalf(`Error compare size query intervals`)
	}

	res = get_ans(root, x)
	//for _, rr := range res {
//		println("[",rr.left, "=", rr.right,"]")/
//	}
}

//
func TestTree(t *testing.T) {
	tree := New()
	if tree.Len() != 0 {
		t.Fatalf(`Error compare size tree`)
	}

	iv1 := SingleInterval(
		50,
		100,
		10,
		&testPointer{"Some text"},
	)
	tree.Add(iv1)

	iv2 := SingleInterval(
		50,
		100,
		12,
		&testPointer{"Some text 2"},
	)
	tree.Add(iv2)

	iv3 := SingleInterval(
		100,
		200,
		15, &testPointer{"Some text 3"},
	)
	tree.Add(iv3)

	iv4 := SingleInterval(
		300,
		400,
		99, &testPointer{"Some text 4"},
	)
	tree.Add(iv4)

	iv5 := SingleInterval(
		300,
		400,
		99, &testPointer{"Some text 5"},
	)
	tree.Add(iv5) // not unique

	if tree.Len() != 4 {
		t.Fatalf(`Error compare size tree`)
	}

	query := ValueInterval(301)

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

	intervals = tree.Query(ValueInterval(69))
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

	intervals = tree.Query(ValueInterval(76))
	if len(intervals) != 1 {
		t.Fatalf(`Error compare size query intervals`)
	}
	if intervals[0].ID() != iv1.ID() {
		t.Fatalf(`Error compare ids`)
	}

}
