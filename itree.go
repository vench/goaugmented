package goaugmented

import "sort"

//
type Segment struct {
	left, right int64
	id          uint64
	data        interface{}
}

//
func (s *Segment) Low() int64 {
	return s.left
}

func (s *Segment) High() int64 {
	return s.right
}

func (s *Segment) Overlaps(Interval) bool {
	return false
}

func (s *Segment) ID() uint64 {
	return s.id
}

func (s *Segment) Data() interface{} {
	return s.data
}

//
func (s *Segment) mean() int64 {
	return (s.right + s.left) / 2
}

//
func NewSegment(left, right int64, data interface{}) *Segment {
	return &Segment{left: left, right: right, data: data}
}

//
type inode struct {
	median        int64
	left, right   *inode
	ileft, iright []*Segment
}

//
func (t *inode) Query(interval Interval) Intervals {
	r := get_ans(t, &Segment{left: interval.Low(), right: interval.High(), data: nil})
	i := make(Intervals, 0, len(r))
	for _, s := range r {
		i = append(i, s)
	}
	return i
}

//
func BuildITree(segments []*Segment) *inode {
	if len(segments) == 0 {
		return nil
	}
	median := median(segments)

	left_child := []*Segment{}
	right_child := []*Segment{}
	left_segments := []*Segment{}
	right_segments := []*Segment{}
	for _, s := range segments {
		if s.right < median {
			left_child = append(left_child, s)
		} else if s.left > median {
			right_child = append(right_child, s)
		} else {
			left_segments = append(left_segments, s)
			right_segments = append(right_segments, s)
		}
	}

	// by left
	sort.Slice(left_segments, func(i, j int) bool {
		return left_segments[i].left < left_segments[j].left
	})
	// by right desc
	sort.Slice(right_segments, func(i, j int) bool {
		return right_segments[i].right > right_segments[j].right
	})
	result := &inode{}
	result.left = BuildITree(left_child)
	result.right = BuildITree(right_child)
	result.ileft = left_segments
	result.iright = right_segments
	result.median = median
	return result
}

// TODO optimize O(N)
func median(s []*Segment) int64 {
	sort.Slice(s, func(i, j int) bool {
		return s[i].mean() > s[j].mean()
	})
	n := len(s)
	if n&0x01 == 1 {
		return s[n/2].mean()
	}
	return (s[n/2-1].mean() + s[n/2].mean()) / 2
}

//
func get_ans(tree *inode, q *Segment) (result []*Segment) {
	if tree == nil {
		return result
	}

	if q.left < tree.median {
		result = append(result, get_ans(tree.left, q)...)
	}

	if q.right > tree.median {
		result = append(result, get_ans(tree.right, q)...)
	}

	if q.right < tree.median {
		for _, item := range tree.ileft {
			if item.left < q.left {
				result = append(result, item)
			} else {
				break
			}
		}
	} else if q.left >= tree.median {
		for _, item := range tree.iright {
			if item.right > q.right {
				result = append(result, item)
			} else {
				break
			}
		}
	}

	return result
}

//
func inorder(root *inode) {
	if root == nil {
		return
	}
	inorder(root.left)

	var l, r int64
	if root.left != nil {
		l = root.left.median
	}
	if root.right != nil {
		r = root.right.median
	}
	println("L: ", l, ", R: ", r, ", M:", root.median)
	for _, i := range root.iright {
		println("\t", "id: ", i.id, ", L: ", i.left, ", R: ", i.right)
	}

	inorder(root.right)
}
