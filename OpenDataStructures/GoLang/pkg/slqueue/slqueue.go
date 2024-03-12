package slqueue

import (
	"fmt"
	"open_data_structures/pkg/collections"
	"open_data_structures/pkg/sllist"
)

var slQueue collections.Queue[int] = &SLQueue[int]{}

type SLQueue[T any] struct {
	data *sllist.SLList[T]
}

func NewSLQueue[T any]() *SLQueue[T] {
	return &SLQueue[T]{
		data: sllist.NewSLList[T](),
	}
}

func (slq *SLQueue[T]) Size() int {
	return slq.data.Size()
}

func (slq *SLQueue[T]) Dequeue() (T, error) {
	value, err := slq.data.Remove(0)
	if err != nil {
		return *new(T), fmt.Errorf("SLQueue is empty")
	}
	return value, err
}

func (slq *SLQueue[T]) Enqueue(x T) {
	slq.data.Add(slq.data.Size(), x)
}
