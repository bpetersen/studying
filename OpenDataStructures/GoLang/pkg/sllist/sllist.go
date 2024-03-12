package sllist

import (
	"fmt"
	"open_data_structures/pkg/collections"
)

type SLList[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

func (s SLList[T]) String() string {
	return fmt.Sprintf("{size: %v, head: %v, tail: %v}", s.size, s.head, s.tail)
}

type node[T any] struct {
	data T
	next *node[T]
}

func (s node[T]) String() string {
	return fmt.Sprintf("{data: %v, next: %v}", s.data, s.next)
}

var _ collections.List[int] = &SLList[int]{}

func NewSLList[T any]() *SLList[T] {
	return &SLList[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (sll *SLList[T]) Size() int {
	return sll.size
}

func (sll *SLList[T]) Get(i int) (T, error) {
	var nullValue T
	if i < 0 || i >= sll.size {
		return nullValue, fmt.Errorf("Index out of range: %d", i)
	}
	var node = sll.getNode(i)
	return node.data, nil
}

func (sll *SLList[T]) Set(i int, x T) (T, error) {
	if i < 0 || i >= sll.size {
		return *new(T), fmt.Errorf("Index out of range: %d", i)
	}
	var node = sll.getNode(i)
	var result = node.data
	node.data = x
	return result, nil
}

func (sll *SLList[T]) Add(i int, x T) error {
	if i < 0 || i > sll.size {
		return fmt.Errorf("Index out of range: %d", i)
	}
	var newNode = node[T]{
		data: x,
	}

	if sll.size == 0 {
		//first node
		sll.head = &newNode
		sll.tail = &newNode
	} else if i == 0 {
		//head
		newNode.next = sll.head
		sll.head = &newNode
	} else if i == sll.size {
		sll.tail.next = &newNode
		sll.tail = &newNode
		//Tail
	} else {
		var prev = sll.getNode(i)
		newNode.next = prev.next
		prev.next = &newNode
	}

	sll.size++
	return nil
}

func (sll *SLList[T]) Remove(i int) (T, error) {
	var result T
	if i < 0 || i >= sll.size {
		return result, fmt.Errorf("Index out of range: %d", i)
	}
	if i == 0 {
		result = sll.head.data
		sll.head = sll.head.next
	} else {
		var prev = sll.getNode(i - 1)
		var next = prev.next.next
		result = prev.next.data
		prev.next = next
	}
	sll.size--
	return result, nil
}

func (sll *SLList[T]) getNode(i int) *node[T] {
	var thisNode = sll.head
	for j := 0; j < i; j++ {
		thisNode = thisNode.next
	}
	return thisNode
}
