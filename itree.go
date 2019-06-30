package goaugmented

import (
	"sort"
	"math/rand"
)

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

	leftChild := make([]Interval, 0)
	rightChild := make([]Interval, 0)
	leftSegments := make([]Interval, 0)
	rightSegments := make([]Interval, 0)
	for _, s := range intervals {
		if s.High() < median {
			leftChild = append(leftChild, s)
		} else if s.Low() > median {
			rightChild = append(rightChild, s)
		} else {
			leftSegments = append(leftSegments, s)
			rightSegments = append(rightSegments, s)
		}
	}

	// by left
	sort.Slice(leftSegments, func(i, j int) bool {
		return leftSegments[i].Low() < leftSegments[j].Low()
	})
	// by right desc
	sort.Slice(rightSegments, func(i, j int) bool {
		return rightSegments[i].High() > rightSegments[j].High()
	})
	result := &inode{}
	result.left = BuildITree(leftChild)
	result.right = BuildITree(rightChild)
	result.ileft = leftSegments
	result.iright = rightSegments
	result.median = median
	return result
}

//
func getDeep(tree *inode) uint64{
	var (
		right, left uint64 = 1, 1
	)
	if tree.right != nil {
		right += getDeep(tree.right)
	}
	if tree.left != nil {
		left += getDeep(tree.left)
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

func medianQ(s []Interval) int64 {
	n := len(s)
	if n == 0{
		return 0
	}
	if n & 0x01 == 1 {
		return quickselect(s, n/2)
	}
	return (quickselect(s, n / 2 - 1) + quickselect(s, n / 2)) / 2
}

func quickselect(s []Interval, k int) int64 {
	if len(s) == 1{
		return intervalMean(s[0])
	}

	in := 0
	if k > 0 {
		in = rand.Intn(k)
	}

	pivot := intervalMean(s[in])
	lows,highs,pivots := []Interval{},[]Interval{},[]Interval{}
	for _, el := range  s {
		m := intervalMean(el)
		if m < pivot {
			lows = append(lows, el)
		}
		if m > pivot {
			highs = append(highs, el)
		}
		if m == pivot {
			pivots = append(pivots, el)
		}
	}

	if k < len(lows) {
		return quickselect(lows, k)
	} else if k < len(lows) + len(pivots) {
		return intervalMean(pivots[0])
	}

	return quickselect(highs, k - len(lows) - len(pivots))
}

//
func intervalMean(i Interval) int64 {
	return (i.High() + i.Low()) / 2
}

//
func get_ans(tree *inode, q Interval) (result Intervals) {
	if tree == nil {
		return result
	}
	m := intervalMean(q)

	if q.Low() == 78 {
		println("median: ", tree.median)
	}

	if m < tree.median {
		result = append(result, get_ans(tree.left, q)...)
	} else if m > tree.median {
		result = append(result, get_ans(tree.right, q)...)
	}

	if m < tree.median {
		for _, item := range tree.ileft {
			if item.Low() < q.Low() {
				result = append(result, item)
			} else {
				break
			}
		}
	} else if m >= tree.median {
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
