package goaugmented

import "sort"

//
type Segment struct {
	left,right int64
	data interface{}
}

//
func ( s *Segment) mean() int64 {
	return (s.right + s.left) / 2
}

//
type inode struct {
	median int64
	left, right *inode
	ileft, iright []*Segment
}

//
func (t*inode) Query(left, right int64) []*Segment {
	return get_ans(t, &Segment{left, right, nil})
}

//
func BuildITree( segments []*Segment) *inode {
	if len(segments) == 0 {
		return nil
	}
	median := median(segments)

	left_child := []*Segment{}
	right_child := []*Segment{}
	left_segments := []*Segment{}
	right_segments := []*Segment{}
	for _,s := range segments {
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
	result.left = BuildITree(left_child);
	result.right = BuildITree(right_child);
	result.ileft = left_segments
	result.iright = right_segments
	result.median = median;
	return result
}

// TODO optimize O(N)
func median(s []*Segment) int64 {
	sort.Slice(s, func(i, j int) bool {
		return s[i].mean() >  s[j].mean()
	})
	n := len(s)
	if n & 0x01 == 1 {
		return  s[n / 2].mean()
	}
	return (s[n / 2-1].mean() + s[n / 2 ].mean()) / 2
}

//
func get_ans( tree *inode,  q *Segment) []*Segment {
	if (tree == nil) {
		return []*Segment{}
	}

	result :=  []*Segment{}
	if q.left < tree.median {
		result = append(result, get_ans(tree.left, q)...)
	}

	if q.right > tree.median {
		result = append(result, get_ans(tree.right, q)...)
	}

	if q.right < tree.median {
		for _,item := range tree.ileft {
			if item.left < q.left {
				result = append(result, item)
			} else {
				break
			}
		}
	} else if q.left >= tree.median {
		for _,item := range tree.iright {
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
	inorder(root.left);

	println("median = ", root.median)

	inorder(root.right);
}

/*
func inorder(root *inode) {
    if root == nil {
    	return
	}
	inorder(root.left);

	println("[", root.i.low , ", " , root.i.high, "] max = ", root.max)

	inorder(root.right);
}


func newINode( i *segment) *inode {
 	temp := &inode{}
	temp.i = i
	temp.max = i.high
	temp.left,temp.right = nil, nil
	return temp
}

func insert(root *inode, i *segment) *inode {
	if (root == nil) {
		return newINode(i)
	}

	l := root.i.low;

	if i.low < l {
		root.left = insert(root.left, i)
	} else {
		root.right = insert(root.right, i);
	}

	if (root.max < i.high) { // update max value root
		root.max = i.high
	}

	return root;
}

// A utility function to check if given two intervals overlap
func doOVerlap(i1 *segment, i2 *segment) bool {
	if i1.low <= i2.high && i2.low <= i1.high {
		return true
	}
	return false
}

// The main function that searches a given interval i in a given
// Interval Tree.
func overlapSearch(root *inode, i *segment) *segment {
	if root == nil {
		return nil
	}

	if doOVerlap(root.i, i) {
		return root.i
	}

	if root.left != nil && root.left.max >= i.low {
		return overlapSearch(root.left, i);
	}

	return overlapSearch(root.right, i);
}
*/