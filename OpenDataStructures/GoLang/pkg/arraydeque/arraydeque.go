package arraydeque

import (
	"fmt"
	"open_data_structures/pkg/arraylist"
	"open_data_structures/pkg/collections"
)

var _ collections.Deque[int] = &ArrayDeque[int]{}

const DEFAULT_INITIAL_CAPACITY = 1

type ArrayDeque[T any] struct {
	data *arraylist.ArrayList[T]
}

func NewArrayDeque[T any](initialCapacity ...int) *ArrayDeque[T] {
	return &ArrayDeque[T]{
		data: arraylist.NewArrayList[T](),
	}
}

func (s ArrayDeque[T]) String() string {
	return fmt.Sprintf("{data: %v}", s.data)
}

func (ad *ArrayDeque[T]) Size() int {
	return ad.data.Size()
}

func (ad *ArrayDeque[T]) AddFirst(x T) {
	ad.data.Add(0, x)
}

func (ad *ArrayDeque[T]) AddLast(x T) {
	ad.data.Add(ad.data.Size(), x)
}

func (ad *ArrayDeque[T]) RemoveFirst() (T, error) {
	value, err := ad.data.Remove(0)
	if err != nil {
		return *new(T), fmt.Errorf("ArrayDeque is empty")
	}
	return value, err
}

func (ad *ArrayDeque[T]) RemoveLast() (T, error) {
	value, err := ad.data.Remove(ad.data.Size() - 1)
	if err != nil {
		return *new(T), fmt.Errorf("ArrayDeque is empty")
	}
	return value, err
}
