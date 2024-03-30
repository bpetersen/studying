package skiplist

import (
	"fmt"
	"math/rand"
	"open_data_structures/pkg/collections"
	"open_data_structures/pkg/utils"
	"strings"
	"time"
)

// A generic type for comparison functions
// It should return an integer that is:
//
//	less than 0 if the first parameter is less than the second
//	greater than 0 if the first parameter is greater than the second
//	equal to 0 if the first parameter is equal to the second
type Comparator[T any] func(a T, b T) int

type SkipList[T any] struct {
	sentinel *slNode[T]
	size     int
	comparer Comparator[T]
}

type slNode[T any] struct {
	data T
	next []*slNode[T]
}

func newSlNode[T any](x T, height int) *slNode[T] {
	return &slNode[T]{
		data: x,
		next: make([]*slNode[T], height+1),
	}
}

func (n *slNode[T]) String() string {
	return fmt.Sprintf("%v->", n.data)
}

var _ collections.SSet[int] = &SkipList[int]{}

func (sl *SkipList[T]) Size() int {
	return sl.size
}

func NewSkipList[T any](comparer Comparator[T]) *SkipList[T] {
	return &SkipList[T]{
		comparer: comparer,
		sentinel: &slNode[T]{
			next: make([]*slNode[T], 1),
		},
	}
}

func (sl *SkipList[T]) resize(i int) {
	newNext := make([]*slNode[T], i+1)
	copy(newNext, sl.sentinel.next)
	sl.sentinel.next = newNext
}

func (sl *SkipList[T]) String() string {
	var builder strings.Builder
	height := 0
	for height <= sl.GetHeight() {
		thisNode := sl.sentinel
		for thisNode != nil {
			builder.WriteString(fmt.Sprintf("node:%+v, ", thisNode))
			thisNode = thisNode.next[height]
		}
		height++
		builder.WriteString("nil\n")
	}

	return builder.String()
}

// The sentinel will always have the same height as the height of the "tallest" node.
// height is defined as the value r such that x appears in L sub r.  So it is zero based index
func (sl *SkipList[T]) GetHeight() int {
	return len(sl.sentinel.next) - 1
}

func (n *slNode[T]) GetHeight() int {
	return len(n.next) - 1
}

// non-deterministic. Not tested because we can take for granted that rand.Int() does it's job. utils.GetHeight is tested as a pure function should be.
func (sl *SkipList[T]) getNewNodeHeight() int {
	rand.Seed(time.Now().UnixNano())
	num := rand.Int()
	return utils.GetRandomHeight(num)
}

func (sl *SkipList[T]) Add(x T) bool {
	height := sl.getNewNodeHeight()
	_, wasAdded := sl.addNode(x, height)
	return wasAdded
}

func (sl *SkipList[T]) addNode(x T, newNodeHeight int) (*slNode[T], bool) {
	u := sl.sentinel
	r := sl.GetHeight()
	stack := make([]*slNode[T], 0)

	for r >= 0 {
		for u.next[r] != nil && sl.comparer(u.next[r].data, x) < 0 {
			u = u.next[r]
		}
		if u.next[r] != nil && sl.comparer(u.next[r].data, x) == 0 {
			return new(slNode[T]), false
		}
		//prepend to the node
		stack = append([]*slNode[T]{u}, stack...)
		r--
	}
	w := newSlNode[T](x, newNodeHeight)
	h := sl.sentinel.GetHeight()
	for h < w.GetHeight() {
		stack = append(stack, sl.sentinel)
		h++
	}
	if sl.GetHeight() < w.GetHeight() {
		sl.resize(newNodeHeight)
	}
	for i := 0; i < len(w.next); i++ {
		w.next[i] = stack[i].next[i]
		stack[i].next[i] = w
	}
	sl.size++
	return w, true
}

func (sl *SkipList[T]) getPredNode(x T) *slNode[T] {
	thisNode := sl.sentinel
	height := sl.GetHeight()
	for height >= 0 {
		//fmt.Printf("thisNode (outside): %+v, next[height]: %+v\n", thisNode, thisNode.next[height])
		for thisNode.next[height] != nil && sl.comparer(thisNode.next[height].data, x) < 0 {
			//	fmt.Printf("thisNode (inside): %+v\n", thisNode)
			thisNode = thisNode.next[height] //move right
		}
		height-- //move down
	}
	return thisNode
}

// return the smallest value y such that y >= x
func (sl *SkipList[T]) Find(x T) (T, error) {
	var predNode = sl.getPredNode(x)
	//fmt.Printf("predNode: %+v\n", predNode)
	if predNode.next[0] == nil {
		return *new(T), fmt.Errorf("no such value y such that y >= %v", x)
	}

	return predNode.next[0].data, nil
}

func (sl *SkipList[T]) Remove(x T) bool {
	removed := false
	u := sl.sentinel
	h := sl.GetHeight()
	newHeight := h
	for h >= 0 {
		for u.next[h] != nil && sl.comparer(u.next[h].data, x) < 0 {
			u = u.next[h]
		}
		if u.next[h] != nil && sl.comparer(u.next[h].data, x) == 0 {
			removed = true
			//we found the node we need to remove
			u.next[h] = u.next[h].next[h]
			if u == sl.sentinel && u.next[h] == nil {
				newHeight--
			}
		}
		h--
	}
	if newHeight < sl.GetHeight() {
		sl.resize(newHeight)
	}
	return removed
}
