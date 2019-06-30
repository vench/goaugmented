package goaugmented

import "sort"

//
type segment struct {
	left, right int64
	id          uint64
	data        interface{}
}

//
func (s *segment) Low() int64 {
	return s.left
}

func (s *segment) High() int64 {
	return s.right
}

func (s *segment) Overlaps(Interval) bool {
	return false
}

func (s *segment) ID() uint64 {
	return s.id
}

func (s *segment) Data() interface{} {
	return s.data
}



//
func NewSegment(left, right int64, data interface{}) Interval {
	return &segment{left: left, right: right, data: data}
}

//
type inode struct {
	median        int64
	left, right   *inode
	ileft, iright []Interval
	height      uint64
}

//
func (t *inode) Query(interval Interval) Intervals {
	return get_ans(t, interval)
}

//
func BuildITree(intervals []Interval) *inode {
	if len(intervals) == 0 {
		return nil
	}
	median := median(intervals)

	left_child := []Interval{}
	right_child := []Interval{}
	left_segments := []Interval{}
	right_segments := []Interval{}
	for _, s := range intervals {
		if s.High() < median {
			left_child = append(left_child, s)
		} else if s.Low() > median {
			right_child = append(right_child, s)
		} else {
			left_segments = append(left_segments, s)
			right_segments = append(right_segments, s)
		}
	}

	// by left
	sort.Slice(left_segments, func(i, j int) bool {
		return left_segments[i].Low() < left_segments[j].Low()
	})
	// by right desc
	sort.Slice(right_segments, func(i, j int) bool {
		return right_segments[i].High() > right_segments[j].High()
	})
	result := &inode{}
	result.left = BuildITree(left_child)
	result.right = BuildITree(right_child)
	result.ileft = left_segments
	result.iright = right_segments
	result.median = median
	return result
}

func get_max_height(tree *inode) int{
	right := 1
	left := 1
	if tree.right != nil {
		right += get_max_height(tree.right)
	}
	if tree.left != nil {
		left += get_max_height(tree.left)
	}

	if right > left {
		return right
	}
	return left
}

// TODO optimize O(N)
func median(s []Interval) int64 {
	if len (s) == 0 {
		return  0
	}
	sort.Slice(s, func(i, j int) bool {
		return intervalMean(s[i]) > intervalMean(s[j])
	})
	n := len(s)
	if n & 0x01 == 1 {
		return intervalMean(s[n/2])
	}
	return (intervalMean(s[n/2-1]) + intervalMean(s[n/2])) / 2
}

func intervalMean(i Interval) int64 {
	return (i.High() + i.Low()) / 2
}

//
func get_ans(tree *inode, q Interval) (result Intervals) {
	if tree == nil {
		return result
	}

	if q.Low() < tree.median {
		result = append(result, get_ans(tree.left, q)...)
	} else if q.High() > tree.median {
		result = append(result, get_ans(tree.right, q)...)
	}

	if q.High() < tree.median {
		for _, item := range tree.ileft {
			if item.Low() < q.Low() {
				result = append(result, item)
			} else {
				break
			}
		}
	} else if q.Low() >= tree.median {
		for _, item := range tree.iright {
			if item.High() > q.High() {
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
		println("\t", "id: ", i.ID(), ", L: ", i.Low(), ", R: ", i.High())
	}

	inorder(root.right)
}
