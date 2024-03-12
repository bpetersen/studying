package slstack

import (
	"fmt"
	"open_data_structures/pkg/collections"
	"open_data_structures/pkg/sllist"
)

type SLStack[T any] struct {
	data *sllist.SLList[T]
}

var _ collections.Stack[int] = &SLStack[int]{}

func NewSLStack[T any]() *SLStack[T] {
	return &SLStack[T]{
		data: sllist.NewSLList[T](),
	}
}

func (sls *SLStack[T]) Push(x T) {
	sls.data.Add(0, x)
}

func (sls *SLStack[T]) Pop() (T, error) {
	value, err := sls.data.Remove(0)
	if err != nil {
		return *new(T), fmt.Errorf("SLStack is empty")
	}
	return value, err
}

func (slq *SLStack[T]) Size() int {
	return slq.data.Size()
}
