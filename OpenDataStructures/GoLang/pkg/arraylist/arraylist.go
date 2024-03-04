package arraylist

import (
	"open_data_structures/pkg/collections"
)

var _ collections.List[int] = &ArrayList[int]{}

const DEFAULT_INITIAL_CAPACITY = 10

type ArrayList[T any] struct {
	data []T
	size int
	head int
}

func NewArrayList[T any](initialCapacity ...int) *ArrayList[T] {
	capacity := DEFAULT_INITIAL_CAPACITY
	if len(initialCapacity) > 0 {
		capacity = initialCapacity[0]
	}

	return &ArrayList[T]{
		data: make([]T, capacity),
		size: 0,
		head: 0,
	}
}

func (al *ArrayList[T]) Size() int {
	return al.size
}

func (al *ArrayList[T]) Add(i int, x T) error {
	return nil
}

func (al *ArrayList[T]) Get(i int) (T, error) {
	var nullValue T
	return nullValue, nil
}

func (al *ArrayList[T]) Remove(i int) (T, error) {
	var nullValue T
	return nullValue, nil
}

func (al *ArrayList[T]) Set(i int, x T) (T, error) {
	var nullValue T
	return nullValue, nil
}
