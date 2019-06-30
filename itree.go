package goaugmented

import (
	"sort"
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
	equal      bool
}

//
func (t *inode) Query(interval Interval) Intervals {
	return getAns(t, interval, t.equal)
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

//
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


//
func intervalMean(i Interval) int64 {
	return (i.High() + i.Low()) / 2
}

//
func getAns(tree *inode, q Interval, equal bool) (result Intervals) {
	if tree == nil {
		return result
	}
	m := intervalMean(q)
	if m < tree.median {
		result = append(result, getAns(tree.left, q, equal)...)
	} else if m > tree.median {
		result = append(result, getAns(tree.right, q, equal)...)
	}

	if m < tree.median {
		for _, item := range tree.ileft {
			if aLessB(item.Low(), q.Low(), equal) {
				result = append(result, item)
			} else {
				break
			}
		}
	} else if m >= tree.median {
		for _, item := range tree.iright {
			if aGreaterB(item.High(), q.High(), equal) {
				result = append(result, item)
			} else {
				break
			}
		}
	}

	return result
}

//
func aLessB(a,b int64, equal bool) bool {
	if equal {
		return a <= b
	}
	return a < b
}

//
func aGreaterB(a,b int64, equal bool) bool {
	if equal {
		return a >= b
	}
	return a > b
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
