package arraystack

import (
	"fmt"
	"open_data_structures/pkg/arraylist"
	"open_data_structures/pkg/collections"
)

var _ collections.Stack[int] = &ArrayStack[int]{}

const DEFAULT_INITIAL_CAPACITY = 1

type ArrayStack[T any] struct {
	data *arraylist.ArrayList[T] //The backing slice that contains the data
}

func (s ArrayStack[T]) String() string {
	return fmt.Sprintf("{data: %v}", s.data)
}

func NewArrayStack[T any](initialCapacity ...int) *ArrayStack[T] {
	return &ArrayStack[T]{
		data: arraylist.NewArrayList[T](),
	}
}

func (as *ArrayStack[T]) Size() int {
	return as.data.Size()
}

func (as *ArrayStack[T]) Push(x T) {
	as.data.Add(as.data.Size(), x)
}

func (as *ArrayStack[T]) Pop() (T, error) {
	value, err := as.data.Remove(as.data.Size() - 1)
	if err != nil {
		return *new(T), fmt.Errorf("Stack is empty")
	}
	return value, err
}
