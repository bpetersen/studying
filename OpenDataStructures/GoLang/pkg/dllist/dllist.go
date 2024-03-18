package dllist

import (
	"fmt"
	"open_data_structures/pkg/collections"
	"strings"
)

type DLList[T any] struct {
	size  int
	dummy *node[T]
}

func (dll *DLList[T]) String() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("dllist: { size: %d }\n", dll.size))
	//return fmt.Sprintf("dllist: { size: %d }", dll.size)
	var current = dll.dummy
	for i := 0; i < dll.size; i++ {
		builder.WriteString(current.String() + "\n")
		current = current.next
	}
	return builder.String()
}

type node[T any] struct {
	next *node[T]
	prev *node[T]
	data T
}

func (n *node[T]) String() string {
	return fmt.Sprintf("Node: { data: %v, next: %v, prev: %v }", n.data, n.next.data, n.prev.data)
}

var _ collections.List[int] = &DLList[int]{}

func NewDLList[T any]() *DLList[T] {
	var dummyNode = node[T]{}
	dummyNode.next = &dummyNode
	dummyNode.prev = &dummyNode
	return &DLList[T]{
		size:  0,
		dummy: &dummyNode,
	}
}

func (dll *DLList[T]) Size() int {
	return dll.size
}

func (dll *DLList[T]) Add(i int, x T) error {
	if i < 0 || i > dll.size {
		return fmt.Errorf("Index out of range: %d", i)
	}
	var next = dll.getNode(i)
	dll.addBefore(next, x)
	return nil
}

func (dll *DLList[T]) addBefore(n *node[T], x T) {
	var newNode = node[T]{
		data: x,
	}
	newNode.prev = n.prev
	newNode.next = n
	newNode.next.prev = &newNode
	newNode.prev.next = &newNode
	dll.size++
}

func (dll *DLList[T]) getNode(i int) *node[T] {
	var thisNode = dll.dummy
	if i < dll.size/2 {
		thisNode = dll.dummy.next
		for j := 0; j < i; j++ {
			thisNode = thisNode.next
		}
	} else {
		for j := dll.size; j > i; j-- {
			thisNode = thisNode.prev
		}
	}
	return thisNode
}

func (dll *DLList[T]) Get(i int) (T, error) {
	var nullValue T
	if i < 0 || i >= dll.size {
		return nullValue, fmt.Errorf("Index out of range: %d", i)
	}
	var result = dll.getNode(i)
	return result.data, nil
}

func (dll *DLList[T]) Set(i int, x T) (T, error) {
	var nullValue T
	if i < 0 || i >= dll.size {
		return nullValue, fmt.Errorf("Index out of range: %d", i)
	}
	var node = dll.getNode(i)
	var previousValue = node.data
	node.data = x
	return previousValue, nil
}

func (dll *DLList[T]) Remove(i int) (T, error) {
	var nullValue T
	if i < 0 || i >= dll.size {
		return nullValue, fmt.Errorf("Index out of range: %d", i)
	}
	var nodeToRemove = dll.getNode(i)
	nodeToRemove.next.prev = nodeToRemove.prev
	nodeToRemove.prev.next = nodeToRemove.next
	dll.size--
	return nodeToRemove.data, nil
}
