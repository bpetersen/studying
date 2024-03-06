package arrayqueue

import (
	"fmt"
	"open_data_structures/pkg/arraylist"
	"open_data_structures/pkg/collections"
)

var _ collections.Queue[int] = &ArrayQueue[int]{}

type ArrayQueue[T any] struct {
	data arraylist.ArrayList[T]
}

func NewArrayQueue[T any]() *ArrayQueue[T] {
	return &ArrayQueue[T]{
		data: *arraylist.NewArrayList[T](),
	}
}

func (aq *ArrayQueue[T]) Size() int {
	return aq.data.Size()
}

func (aq *ArrayQueue[T]) Enqueue(x T) {
	aq.data.Add(aq.data.Size(), x)
}

func (aq *ArrayQueue[T]) Dequeue() (T, error) {
	value, err := aq.data.Remove(aq.data.Size() - 1)
	if err != nil {
		return *new(T), fmt.Errorf("ArrayQueue is empty")
	}
	return value, err
}
